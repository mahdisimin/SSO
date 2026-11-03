package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

func TestSign_ValidSignature(t *testing.T) {
	signer, key := generate_signer(t)
	token, signErr := signer.Sign(map[string]any{"foo": "bar"})
	if signErr != nil {
		t.Fatal(signErr)
	}
	if verificationErr := verify(token, &key.PublicKey); verificationErr != nil {
		t.Fatal(verificationErr)
	}

}

func TestSign_InvalidSignature(t *testing.T) {
	signer, key := generate_signer(t)
	token, signErr := signer.Sign(map[string]any{"foo": "bar"})
	tokenPart := strings.Split(token, ".")
	if signErr != nil {
		t.Fatal(signErr)
	}
	fakePayload := map[string]any{"foooo": "barrr"}
	fakePayloadJson, _ := json.Marshal(fakePayload)
	fakePayloadBase64 := base64.RawURLEncoding.EncodeToString(fakePayloadJson)
	fakeToken := tokenPart[0] + "." + fakePayloadBase64 + "." + tokenPart[2]
	verifyErr := verify(fakeToken, &key.PublicKey)
	if verifyErr == nil {
		t.Fatal(verifyErr)
	}
}

func TestSign_TokenConent(t *testing.T) {
	t.Helper()

	signer, _ := generate_signer(t)
	token, signErr := signer.Sign(map[string]any{"foo": "bar"})
	if signErr != nil {
		t.Fatal(signErr)
	}
	tokenPart := strings.Split(token, ".")
	var header map[string]any
	if len(tokenPart) != 3 {
		t.Fatal("token should have 3 parts")
	}
	headerJson, DecodeErr := base64.RawURLEncoding.DecodeString(tokenPart[0])
	if DecodeErr != nil {
		t.Fatal(DecodeErr)
	}
	if UnmarshalErr := json.Unmarshal(headerJson, &header); UnmarshalErr != nil {
		t.Fatal(UnmarshalErr)
	}
	if header["kid"] != "test_kid" || header["alg"] != "RS256" || header["typ"] != "JWT" {
		t.Fatal("invalid token")
	}
}

func TestSign_WrongKey(t *testing.T) {
	signer, _ := generate_signer(t)
	token, err := signer.Sign(map[string]any{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	fakeKey, generateKeyErr := rsa.GenerateKey(rand.Reader, 2048)
	if generateKeyErr != nil {
		t.Fatal(generateKeyErr)
	}
	if verifyErr := verify(token, &fakeKey.PublicKey); verifyErr == nil {
		t.Fatal(verifyErr)
	}
}

func TestSign_BadToken(t *testing.T) {
	signer, _ := generate_signer(t)
	token, err := signer.Sign(map[string]any{"foo": make(chan int)})
	if err == nil || token != "" {
		t.Fatal(err)
	}
}
