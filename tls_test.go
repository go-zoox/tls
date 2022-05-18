package tls

import (
	"fmt"
	"testing"
)

func TestClientServer(t *testing.T) {
	privateKey := "MIICXAIBAAKBgQCmuEbvwpNB//67t/2g5cGMFkNbkmRbRmtDBK+hjboF66ml7hdbYPF09GNHQbl7b8Ru1hWTNhtu37GF4X0zg8nIU0HSMRSfvzUJ2SWKoAKXPy0jYQk2pxYpkYr3RfMMYVaEa55sT+0MGipSmpibrwkUL6W7k7CTYZpd/9J3JjAWRwIDAQABAoGANo0tiN4d2QaujzXQ44jKH9BZEemAtO0Bw9gQr8f0CmPmCskxE1FRMHeW1IYI7v7PQ4UBYj3eFBVVvPzfPq/sofxQnwVYVSVjWgz6NdZsaTAFh1YxnJx/IzAQFeWFyigZbmSBrMtLAer2G6inEOornzqT0+n8GEkeOpG+h7s54iECQQDaxwtpa11fiZov7dKdFJeOsYoGizBHafkA3/PrnRZxjhj+orWy87ev8Ltp+rz/5JnrHn7Pq31hgDn8LIWRiCyLAkEAwxXWCYTwE3N6KZ7UgxHpxBOaKyiQBLUfhu9rMeSyM4xdgbT6ByEwPjJxuBrqFQoaBUSLoX6vFGohJkFoUdTItQJACgctortlIEfyZVgFW2XiPIwuw3YF1IArBbs+NwKQUMwuoR1cLsO1G79xF76Cg0g7NefD8EjwClQSVFjGFpGjWQJAXcE4xApndnGg3C/A4dzSA7GH/gXYcOq65BZb5faKzcs/hP58ysBgdwO3M0t8A/B+4Nk4YbyIV79JfyEgCXPBoQJBALMJZOROVZZND4dUQAxk2+aRR+JIC7R1VDvNiUsQTiv9BRIs7l4qAiwCuDpdIL7y9t2AO8kc+5wINkjJUs4dq1Q="
	publickKey := "MIGJAoGBAKa4Ru/Ck0H//ru3/aDlwYwWQ1uSZFtGa0MEr6GNugXrqaXuF1tg8XT0Y0dBuXtvxG7WFZM2G27fsYXhfTODychTQdIxFJ+/NQnZJYqgApc/LSNhCTanFimRivdF8wxhVoRrnmxP7QwaKlKamJuvCRQvpbuTsJNhml3/0ncmMBZHAgMBAAE="

	client := NewClient(publickKey)
	server := NewServer(privateKey)

	fmt.Println("client public key:", client.GetPublicKey())
	fmt.Println("server private key:", server.GetPrivateKey())
	if client.GetPublicKey() != publickKey {
		t.Error("client public key is not equal to server private key")
	}
	if server.GetPrivateKey() != privateKey {
		t.Error("server private key is not equal to client public key")
	}
	if server.GetPublicKey() != client.GetPublicKey() {
		t.Error("server public key is not equal to client public key")
	}

	hash := client.NegotiateGenerate()
	ok, err := server.NegotiateVerify(hash)
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("client and server negotiate failed")
	}

	fmt.Println("client secret:", client.GetSecret())
	fmt.Println("server secret:", server.GetSecret())
	if client.GetSecret() != server.GetSecret() {
		t.Error("client and server secret not match")
	}

	plainbytesC2S := []byte("hello world")
	cipherbytes, err := client.Encrypt(plainbytesC2S)
	if err != nil {
		t.Error(err)
	}

	plainbytesC2S2, err := server.Decrypt(cipherbytes)
	if err != nil {
		t.Error(err)
	}

	if string(plainbytesC2S) != string(plainbytesC2S2) {
		t.Error("client and server encrypt and decrypt failed")
	}

	plainbytesS2C := []byte("hello world 2")
	cipherbytes, err = server.Encrypt(plainbytesS2C)
	if err != nil {
		t.Error(err)
	}

	plainbytesS2C2, err := client.Decrypt(cipherbytes)
	if err != nil {
		t.Error(err)
	}

	if string(plainbytesS2C) != string(plainbytesS2C2) {
		t.Error("server and client encrypt and decrypt failed")
	}
}
