package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var version = "1.0.0"

func main() {
	opts, err := parseOptions()
	if err != nil {
		log.Fatal(err)
	}

	if opts.versionMode {
		fmt.Printf("kubeoidc-web v%s\n", version)
		return
	}

	err = startServer(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func generateState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	w := &bytes.Buffer{}
	e := base64.NewEncoder(base64.StdEncoding, w)
	e.Write(b)

	str, err := ioutil.ReadAll(w)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

type Options struct {
	issuerURL    string
	clientID     string
	clientSecret string
	listenAddr   string
	callbackAddr string
	versionMode  bool
}

func parseOptions() (*Options, error) {
	listenAddrDefault := ":8080"
	if s := os.Getenv("KUBEOIDC_LISTEN_ADDR"); s != "" {
		listenAddrDefault = s
	}

	issuerURL := flag.String("issuer", os.Getenv("KUBEOIDC_ISSUER"), "Issuer URL")
	clientID := flag.String("client-id", os.Getenv("KUBEOIDC_CLIENT_ID"), "Client ID")
	clientSecret := flag.String("client-secret", os.Getenv("KUBEOIDC_CLIENT_SECRET"), "Client secret")
	callbackAddr := flag.String("callback-addr", os.Getenv("KUBEOIDC_CALLBACK_ADDR"), "Callback address (e.g. https://example.com/callback)")
	listenAddr := flag.String("listen-addr", listenAddrDefault, "Listen address")
	versionMode := flag.Bool("version", false, "Show version")
	flag.Parse()

	opts := &Options{
		issuerURL:    *issuerURL,
		clientID:     *clientID,
		clientSecret: *clientSecret,
		listenAddr:   *listenAddr,
		callbackAddr: *callbackAddr,
		versionMode:  *versionMode,
	}

	if opts.issuerURL == "" {
		return nil, errors.New("--issuer is required")
	}
	if opts.clientID == "" {
		return nil, errors.New("--client-id is required")
	}
	if opts.clientSecret == "" {
		return nil, errors.New("--client-secret is required")
	}
	if opts.callbackAddr == "" {
		return nil, errors.New("--callback-addr is required")
	}

	return opts, nil
}

func startServer(opts *Options) error {
	ctx := context.Background() // FIXME
	provider, err := oidc.NewProvider(ctx, opts.issuerURL)
	if err != nil {
		return err
	}

	oauth := oauth2.Config{
		ClientID:     opts.clientID,
		ClientSecret: opts.clientSecret,
		RedirectURL:  opts.callbackAddr,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "groups", "offline_access"},
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: opts.clientID})

	s := Server{
		oauth:      oauth,
		verifier:   verifier,
		issuerURL:  opts.issuerURL,
		listenAddr: opts.listenAddr,
		states:     sync.Map{},
	}
	s.Start()

	return nil
}

type Server struct {
	oauth      oauth2.Config
	verifier   *oidc.IDTokenVerifier
	issuerURL  string
	listenAddr string
	states     sync.Map
}

func (s *Server) Start() {
	log.Println("Starting")

	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/assets/templates/index.html", gin.H{})
	})
	r.GET("/initiate", func(c *gin.Context) {
		state, err := generateState()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err)
			return
		}
		s.states.Store(state, nil)

		c.Redirect(http.StatusTemporaryRedirect, s.oauth.AuthCodeURL(state)) // FIXME
	})
	r.GET("/callback", func(c *gin.Context) {
		state := c.Query("state")
		if _, ok := s.states.Load(state); ok {
			s.states.Delete(state)
		} else {
			c.String(http.StatusBadRequest, "Invalid state parameter")
			return
		}

		oauth2Token, err := s.oauth.Exchange(context.Background(), c.Query("code"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error %s", err)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			c.String(http.StatusInternalServerError, "Invalid id_token")
			return
		}

		idToken, err := s.verifier.Verify(context.Background(), rawIDToken)
		if err != nil {
			c.String(http.StatusInternalServerError, "Token verification failed")
			return
		}

		var claims struct {
			Email string `json:"email"`
		}
		if err := idToken.Claims(&claims); err != nil {
			c.String(http.StatusInternalServerError, "Invalid claim")
			return
		}

		c.HTML(http.StatusOK, "/assets/templates/callback.html", gin.H{
			"configYAML":     s.generateKubeconfig(rawIDToken, oauth2Token.RefreshToken),
			"kubectlCommand": s.generateKubectlCommand(rawIDToken, oauth2Token.RefreshToken),
		})
	})
	r.Run(s.listenAddr)
}

func (s *Server) generateKubectlCommand(idToken string, refreshToken string) string {
	f := `kubectl config set-credentials '%s' \
    --auth-provider=oidc \
    '--auth-provider-arg=client-id=%s' \
    '--auth-provider-arg=client-secret=%s' \
    '--auth-provider-arg=id-token=%s' \
    '--auth-provider-arg=idp-issuer-url=%s' \
    '--auth-provider-arg=refresh-token=%s'`

	return fmt.Sprintf(f, s.issuerURL, s.oauth.ClientID, s.oauth.ClientSecret, idToken, s.issuerURL, refreshToken)
}

func (s *Server) generateKubeconfig(idToken string, refreshToken string) string {
	c := map[string]interface{}{
		"users": []map[string]interface{}{{
			"name": s.issuerURL,
			"user": map[string]interface{}{
				"auth-provider": map[string]interface{}{
					"name": "oidc",
					"config": map[string]interface{}{
						"client-id":      s.oauth.ClientID,
						"client-secret":  s.oauth.ClientSecret,
						"idp-issuer-url": s.issuerURL,
						"id-token":       idToken,
						"refresh-token":  refreshToken,
					},
				},
			},
		}},
	}

	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err) // never reach
	}

	return string(out)
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasPrefix(name, "/assets/templates/") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
