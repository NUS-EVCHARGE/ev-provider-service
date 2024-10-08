package handler

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func GetPreLoginHandler(c *gin.Context) {
	publicKey, err := loadPublicKey("public.pem")

	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateResponse("Failed to load public key"))
		return
	}

	pemkey, err := publicKeyToPEM(publicKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateResponse("Failed to convert public key to pem"))
		return
	}
	
	c.JSON(http.StatusOK, CreateResponse("success", pemkey))

	return

}

func publicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	pemKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pemKey), nil
}

func loadPublicKey(fileName string) (*rsa.PublicKey, error) {
	pubKeyFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer pubKeyFile.Close()

	pubKeyPEM, err := ioutil.ReadAll(pubKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pubKeyPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return rsaPubKey, nil
}
