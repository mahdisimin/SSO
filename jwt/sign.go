package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Sign struct {
	sercretKey *rsa.PrivateKey
	kid        string
}

func NewSigner(key *rsa.PrivateKey, kid string) *Sign {
	return &Sign{
		sercretKey: key,
		kid:        kid,
	}
}

func (s *Sign) Sign(claims map[string]any) ([]byte, error) {
	header := map[string]any{"alg": "RS256", "typ": "JWT", "kid": s.kid}

	claimBytes, _ := json.Marshal(claims)
	headerBytes, _ := json.Marshal(header)

	signingInput := base64.RawURLEncoding.EncodeToString(headerBytes) + "." + base64.RawURLEncoding.EncodeToString(claimBytes)
	hash := sha256.Sum256([]byte(signingInput))
	sig, _ := rsa.SignPKCS1v15(rand.Reader, s.sercretKey, crypto.SHA256, hash[:])

	return []byte(signingInput + "." + base64.RawURLEncoding.EncodeToString(sig)), nil
}
