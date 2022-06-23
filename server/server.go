package server

import (
	"github.com/go-zoox/crypto/aes"
	"github.com/go-zoox/crypto/rsa"
)

// Server is a server for TLS.
type Server struct {
	// private key
	privateKey string
	// secret
	secret []byte
	// algorithms
	asymmetric *rsa.RSA
	symmetric  *aes.CFB
}

// New creates a new TLS server.
func New(privateKey string) *Server {
	asymmetric, _ := rsa.New(privateKey)
	return &Server{
		privateKey: privateKey,
		asymmetric: asymmetric,
	}
}

// NegotiateVerify verifies the hash and decrypts the secret.
func (t *Server) NegotiateVerify(hash string) (bool, error) {
	secret, err := t.asymmetric.Decrypt([]byte(hash))
	if err != nil {
		return false, err
	}

	t.secret = secret
	t.symmetric, _ = aes.NewCFB(len(secret)*8, &aes.Base64Encoding{}, nil)
	// fmt.Println("server secret:", t.secret)
	return true, nil
}

// Encrypt encrypts the plaintext with the secret.
func (t *Server) Encrypt(plainbytes []byte) (cipherbytes []byte, err error) {
	cipherbytes, err = t.symmetric.Encrypt(plainbytes, []byte(t.secret))
	return
}

// Decrypt decrypts the ciphertext with the secret.
func (t *Server) Decrypt(cipherbytes []byte) (plainbytes []byte, err error) {
	plainbytes, err = t.symmetric.Decrypt(cipherbytes, []byte(t.secret))
	return
}

// GetSecret returns the secret.
func (t *Server) GetSecret() []byte {
	return t.secret
}

// GetPrivateKey returns the private key.
func (t *Server) GetPrivateKey() string {
	return t.privateKey
}

// GetPublicKey returns the public key.
func (t *Server) GetPublicKey() string {
	return t.asymmetric.GetPublickKey()
}
