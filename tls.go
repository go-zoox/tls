package tls

import (
	"github.com/go-zoox/tls/client"
	"github.com/go-zoox/tls/server"
)

// NewClient creates a new TLS client.
func NewClient(publicKey string) *client.Client {
	return client.New(publicKey)
}

// NewServer creates a new TLS server.
func NewServer(privateKey string) *server.Server {
	return server.New(privateKey)
}
