package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
)

func LoadOrGenerateKey(path string) (*rsa.PrivateKey, error) {
	if fileData, err := os.ReadFile(path); errors.Is(err, os.ErrNotExist) {
		return generateKey(path)
	} else if err != nil {
		return nil, errors.New("Error on Open New key : " + err.Error())
	} else {
		return loadKey(fileData)
	}
}

func loadKey(fileData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(fileData)
	if block == nil {
		return nil, errors.New("error on pem decode")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, errors.New("error on pem type decode , type is : " + block.Type)
	}
	rsaKeyTmp, parsErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	if parsErr != nil {
		return nil, errors.New("Error on Parsing the key : " + parsErr.Error())
	}
	rsaKey, ok := rsaKeyTmp.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("error on type assertion the key")
	}
	return rsaKey, nil
}

func generateKey(path string) (*rsa.PrivateKey, error) {
	rsaKey, pvKeyErr := rsa.GenerateKey(rand.Reader, 2048)
	if pvKeyErr != nil {
		return nil, errors.New("Error on Generate New key : " + pvKeyErr.Error())
	}
	x509Encoded, marshalErr := x509.MarshalPKCS8PrivateKey(rsaKey)
	if marshalErr != nil {
		return nil, errors.New("Error on marshaling the generated key : " + marshalErr.Error())
	}
	data := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	mkDirErr := os.MkdirAll(filepath.Dir(path), 0700)
	if mkDirErr != nil {
		return nil, errors.New("Error on creating the directory : " + mkDirErr.Error())
	}
	writeFileErr := os.WriteFile(path, data, 0600)
	if writeFileErr != nil {
		return nil, errors.New("Error on Write New key : " + writeFileErr.Error())
	}
	return rsaKey, nil
}
