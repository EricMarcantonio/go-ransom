package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"log"
	"os"
)

func generateKeys() {
	secret = make([]byte, 32)
	if _, err := rand.Reader.Read(secret); err != nil {
		log.Fatalln(err)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalln(err)
	}
	publicKey := privateKey.PublicKey
	encryptedBytes, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		&publicKey,
		secret,
		nil)
	file, err := os.Create("oops.pem")
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = file.Write(encryptedBytes)
	_ = file.Close()
}
