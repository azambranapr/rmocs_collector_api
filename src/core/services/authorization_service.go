package services

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"sync"
)

var (
	singnkey  *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadFiles rsa file anf pub file (publica)
func LoadFiles(privateFiles, publicFiles string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFiles, publicFiles)
	})
	return err
}

// loadFiles rsa file anf pub file (privada)
func loadFiles(privateFile, publicFile string) error {
	privateByte, err := os.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicByte, err := os.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateByte, publicByte)
}

// parseRSA interpreta la data del certificado a leer
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	singnkey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)

	if err != nil {
		return err
	}
	return nil
}
