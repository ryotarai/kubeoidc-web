.PHONY: assets

bin/kubeoidc-web: assets
	go build -o bin/kubeoidc-web

assets:
	go-assets-builder assets -o assets.go

