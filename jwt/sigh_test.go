package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func generate_signer(t *testing.T) (signer Signer, key *rsa.PrivateKey) {
	t.Helper()

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return Signer{
		privateKey: key,
		kid:        "test_kid",
	}, key
}

func verify(token string, key *rsa.PublicKey) error {
	jwtPart := strings.Split(token, ".")
	if len(jwtPart) != 3 {
		return errors.New("invalid token")
	}
	hash := sha256.Sum256([]byte(jwtPart[0] + "." + jwtPart[1]))
	sign, decodeErr := base64.RawURLEncoding.DecodeString(jwtPart[2])
	if decodeErr != nil {
		return errors.New("Error decoding sign : " + decodeErr.Error())
	}
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hash[:], sign)
}
