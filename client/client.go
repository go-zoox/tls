package client

import (
	"fmt"
	"time"

	"github.com/go-zoox/crypto/aes"
	"github.com/go-zoox/crypto/rsa"
	"github.com/go-zoox/random"
)

// Client is a client for TLS.
type Client struct {
	// public key
	publicKey string
	// secret
	secret []byte
	// algorithms
	asymmetric *rsa.RSAEncryptor
	symmetric  *aes.CFB
}

// New creates a new TLS client.
func New(publicKey string) *Client {
	asymmetric, _ := rsa.NewEncryptor(publicKey)
	return &Client{
		publicKey:  publicKey,
		asymmetric: asymmetric,
	}
}

// NegotiateGenerate generates a secret and encrypts it with the public key.
//	The secret is used to encrypt the data.
//	The hash is used to negotiate the secret.
func (t *Client) NegotiateGenerate() string {
	// @TODO milliseconds, length: 13
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))

	// AES-256-CFB
	// length = 32 - 13 - 1 = 18
	random := random.String(18)
	t.secret = []byte(fmt.Sprintf("%s:%s", timestamp, random))

	// fmt.Println("client secret:", t.secret)
	t.symmetric, _ = aes.NewCFB(len(t.secret)*8, &aes.Base64Encoding{}, nil)

	hash, _ := t.asymmetric.Encrypt([]byte(t.secret))
	return string(hash)
}

// Encrypt encrypts the plaintext with the secret.
func (t *Client) Encrypt(plainbytes []byte) (cipherbytes []byte, err error) {
	cipherbytes, err = t.symmetric.Encrypt(plainbytes, []byte(t.secret))
	return
}

// Decrypt decrypts the ciphertext with the secret.
func (t *Client) Decrypt(cipherbytes []byte) (plainbytes []byte, err error) {
	plainbytes, err = t.symmetric.Decrypt(cipherbytes, t.secret)
	return
}

// GetSecret returns the secret.
func (t *Client) GetSecret() []byte {
	return t.secret
}

// GetPublicKey returns the public key.
func (t *Client) GetPublicKey() string {
	return t.publicKey
}
