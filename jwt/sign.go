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

func (s *Signer) Sign(claims map[string]any) (string, error) {
	header := map[string]any{"alg": "RS256", "typ": "JWT", "kid": s.kid}

	claimBytes, claimMarshalErr := json.Marshal(claims)
	if claimMarshalErr != nil {
		return nil, claimMarshalErr
	}
	headerBytes, headerMarshalErr := json.Marshal(header)
	if headerMarshalErr != nil {
		return nil, headerMarshalErr
	}

	signingInput := base64.RawURLEncoding.EncodeToString(headerBytes) + "." + base64.RawURLEncoding.EncodeToString(claimBytes)
	hash := sha256.Sum256([]byte(signingInput))
	sig, signErr := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hash[:])
	if signErr != nil {
		return "", signErr
	}
	return signingInput + "." + base64.RawURLEncoding.EncodeToString(sig), nil
}
