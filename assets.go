package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets420fafe7e9b3329f2d135e3e0ec9509c53e74a57 = "<!doctype html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <meta name=\"description\" content=\"\">\n    <meta name=\"author\" content=\"\">\n\n    <title>kubeoidc-web</title>\n\n    <!-- Bootstrap core CSS -->\n    <link href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n    <style>\n/* Sticky footer styles\n-------------------------------------------------- */\nhtml {\n  position: relative;\n  min-height: 100%;\n}\nbody {\n  margin-bottom: 60px; /* Margin bottom by footer height */\n}\n.footer {\n  position: absolute;\n  bottom: 0;\n  width: 100%;\n  height: 60px; /* Set the fixed height of the footer here */\n  line-height: 60px; /* Vertically center the text there */\n  background-color: #f5f5f5;\n}\n\n\n/* Custom page CSS\n-------------------------------------------------- */\n/* Not required for template or sticky footer method. */\n\n.container {\n  width: auto;\n  max-width: 680px;\n  padding: 0 15px;\n}\n    </style>\n  </head>\n\n  <body>\n\n    <!-- Begin page content -->\n    <main role=\"main\" class=\"container\">\n      <h1 class=\"mt-5\">kubeoidc</h1>\n      <a href=\"/initiate\" class=\"btn btn-primary mt-2\">Get credential</a>\n    </main>\n\n    <footer class=\"footer\">\n      <div class=\"container\">\n        <span class=\"text-muted\">kubeoidc</span>\n      </div>\n    </footer>\n  </body>\n</html>\n"
var _Assets463cd279ea50e380157c49e27202495a99d46f8e = "<!doctype html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <meta name=\"description\" content=\"\">\n    <meta name=\"author\" content=\"\">\n\n    <title>kubeoidc-web</title>\n\n    <!-- Bootstrap core CSS -->\n    <link href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n    <style>\n/* Sticky footer styles\n-------------------------------------------------- */\nhtml {\n  position: relative;\n  min-height: 100%;\n}\nbody {\n  margin-bottom: 60px; /* Margin bottom by footer height */\n}\n.footer {\n  position: absolute;\n  bottom: 0;\n  width: 100%;\n  height: 60px; /* Set the fixed height of the footer here */\n  line-height: 60px; /* Vertically center the text there */\n  background-color: #f5f5f5;\n}\n\n\n/* Custom page CSS\n-------------------------------------------------- */\n/* Not required for template or sticky footer method. */\n\n.container {\n  width: auto;\n  max-width: 680px;\n  padding: 0 15px;\n}\n\n.code {\n  font-family: SFMono-Regular,Menlo,Monaco,Consolas,\"Liberation Mono\",\"Courier New\",monospace;\n}\n    </style>\n  </head>\n\n  <body>\n\n    <!-- Begin page content -->\n    <main role=\"main\" class=\"container\">\n      <h1 class=\"mt-5\">Your credential is successfully prepared</h1>\n      <p>Please execute the kubectl command in your terminal or place the kubeconfig in <code>~/.kube/config</code></p>\n\n      <h2 class=\"mt-2\">kubectl commands</h2>\n      <textarea id=\"kubectlCommand\" class=\"code form-control mt-2\" readonly>{{.kubectlCommand}}</textarea>\n      <button class=\"btn btn-primary clipboard mt-2\" data-clipboard-target=\"#kubectlCommand\">\n        Copy to clipboard\n      </button>\n\n      <h2 class=\"mt-2\">kubeconfig snippet</h2>\n      <textarea id=\"configYAML\" class=\"code form-control mt-2\" readonly>{{.configYAML}}</textarea>\n      <button class=\"btn btn-primary clipboard mt-2\" data-clipboard-target=\"#configYAML\">\n        Copy to clipboard\n      </button>\n    </main>\n\n    <footer class=\"footer\">\n      <div class=\"container\">\n        <span class=\"text-muted\">kubeoidc</span>\n      </div>\n    </footer>\n\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.0/clipboard.min.js\"></script>\n    <script>\n        var clipboard = new ClipboardJS('.clipboard');\n        clipboard.on('success', function(e) {\n            console.log(e);\n        });\n        clipboard.on('error', function(e) {\n            console.log(e);\n        });\n    </script>\n  </body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{}, "/assets/templates": []string{"index.html", "callback.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1534998551, 1534998551349818278),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1532082053, 1532082053608875894),
		Data:     nil,
	}, "/assets/templates": &assets.File{
		Path:     "/assets/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1532082053, 1532082053608875828),
		Data:     nil,
	}, "/assets/templates/index.html": &assets.File{
		Path:     "/assets/templates/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1534996836, 1534996836646441259),
		Data:     []byte(_Assets420fafe7e9b3329f2d135e3e0ec9509c53e74a57),
	}, "/assets/templates/callback.html": &assets.File{
		Path:     "/assets/templates/callback.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1534996909, 1534996909936008934),
		Data:     []byte(_Assets463cd279ea50e380157c49e27202495a99d46f8e),
	}}, "")
