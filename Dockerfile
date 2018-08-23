FROM golang:1.10 AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /usr/bin/kubeoidc-web .

###############################################

FROM ubuntu:16.04
RUN apt update && apt -y install ca-certificates
COPY --from=builder /usr/bin/kubeoidc-web /usr/bin/kubeoidc-web
CMD ["/usr/bin/kubeoidc-web"]
