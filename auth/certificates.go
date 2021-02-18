package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

//LoadFiles carga los certificados
func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

func loadFiles(privateFile, publicFile string) error {
	privateBites, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBites, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBites, publicBites)
}

func parseRSA(privateBites, publicBites []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBites)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBites)
	if err != nil {
		return err
	}
	return nil
}
