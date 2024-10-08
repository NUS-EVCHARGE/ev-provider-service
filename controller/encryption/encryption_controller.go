package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	third_party "github.com/NUS-EVCHARGE/ev-provider-service/third_party/aws"
	"github.com/sirupsen/logrus"
)

type EncryptionController interface {
	DecryptPassword(value string) (string, error)
}

type EncryptionControllerImpl struct {
}

var (
	EncryptionControllerObj EncryptionController
)

func NewEncryptionController() {
	EncryptionControllerObj = &EncryptionControllerImpl{}
}

func (e *EncryptionControllerImpl) DecryptPassword(value string) (string, error) {
	decodedPassword, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	sm := third_party.SecretManager{
		Region:     "ap-southeast-1",
		SecretName: "Sigin-PrivateKey",
	}
	privateKey := ""
	secretErr := sm.GetSecrets(&privateKey)

	if secretErr != nil {
		logrus.WithField("error", err).Error("failed to get secret from secret manager")
	}

	rsaPrivateKey, err := parseRSAPrivateKeyFromString(privateKey)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodedPassword)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

// Function to convert PEM encoded string to *rsa.PrivateKey
func parseRSAPrivateKeyFromString(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}