package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Signer struct {
	privateKey *rsa.PrivateKey
	kid        string
}

func NewSigner(key *rsa.PrivateKey, kid string) *Signer {
	return &Signer{
		privateKey: key,
		kid:        kid,
	}
}

func (s *Signer) Sign(claims map[string]any) ([]byte, error) {
	header := map[string]any{"alg": "RS256", "typ": "JWT", "kid": s.kid}

	claimBytes, claimMarshalErr := json.Marshal(claims)
	if claimMarshalErr != nil {
		return nil, claimMarshalErr
	}
	headerBytes, HeaderMarshalErr := json.Marshal(header)
	if HeaderMarshalErr != nil {
		return nil, HeaderMarshalErr
	}

	signingInput := base64.RawURLEncoding.EncodeToString(headerBytes) + "." + base64.RawURLEncoding.EncodeToString(claimBytes)
	hash := sha256.Sum256([]byte(signingInput))
	sig, _ := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hash[:])

	return []byte(signingInput + "." + base64.RawURLEncoding.EncodeToString(sig)), nil
}
