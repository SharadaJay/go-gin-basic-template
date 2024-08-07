package utils

import (
	"basic-gin-app/internal/model"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

var Cfg map[string]model.Config

func init() {
	readPropertyFile(&Cfg)
}

func readPropertyFile(cfg *map[string]model.Config) {
	path := os.ExpandEnv("properties/${ENV}.yml")
	byteArray, err := os.ReadFile(path)
	if err != nil {
		handleError(err)
	}

	err = yaml.Unmarshal(byteArray, cfg)
	if err != nil {
		handleError(err)
	}
}

func DecryptAes(stringToDecrypt string, keyString string) (decodedString string) {
	key := []byte(keyString)
	encryptedText, _ := hex.DecodeString(stringToDecrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		handleError(err)
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		handleError(err)
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	if len(encryptedText) < nonceSize {
		fmt.Println(err)
	}

	//Extract the nonce from the encrypted data
	nonce, ciphertext := encryptedText[:nonceSize], encryptedText[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		handleError(err)
	}

	return string(plaintext)
}

func handleError(err error) {
	log.Fatalf("FATAL_ERROR : %v", err)
}
