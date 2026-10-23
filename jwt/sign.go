package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type sign struct {
	sercretKey *rsa.PrivateKey
	kid        string
}

/*
۱. header   = {"alg":"RS256", "typ":"JWT", "kid": s.kid}
۲. هر دو رو json.Marshal کن (header و claims)
۳. هر دو رو base64.RawURLEncoding کن
۴. signingInput = headerB64 + "." + claimsB64
۵. hash = sha256.Sum256([]byte(signingInput))
۶. sig  = rsa.SignPKCS1v15(rand.Reader, s.key, crypto.SHA256, hash[:])
۷. return signingInput + "." + base64url(sig)
*/

func NewSign(key *rsa.PrivateKey, kid string) *sign {
	return &sign{
		sercretKey: key,
		kid:        kid,
	}
}

func (s *sign) Sign(claims map[string]any) ([]byte, error) {
	header := map[string]any{"alg": "RS256", "typ": "JWT", "kid": s.kid}

	claimBytes, _ := json.Marshal(claims)
	headerBytes, _ := json.Marshal(header)

	base64.RawURLEncoding.Encode(headerBytes, headerBytes)
	base64.RawURLEncoding.Encode(claimBytes, claimBytes)

	signingInput := base64.RawURLEncoding.EncodeToString(headerBytes) + "." + base64.RawURLEncoding.EncodeToString(claimBytes)
	hash := sha256.Sum256([]byte(signingInput))
	sig, _ := rsa.SignPKCS1v15(rand.Reader, s.sercretKey, crypto.SHA256, hash[:])

	return []byte(signingInput + "." + base64.RawURLEncoding.EncodeToString(sig)), nil
}
