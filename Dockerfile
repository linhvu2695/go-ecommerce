FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o go.ecommerce.com ./cmd/server

FROM alpine:latest

# Install CA certificates for TLS validation
RUN apk add --no-cache ca-certificates

COPY ./config /config
COPY ./templates /templates

COPY --from=builder /build/go.ecommerce.com /

ENTRYPOINT [ "/go.ecommerce.com", "config/local.yaml" ]