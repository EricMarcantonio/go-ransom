package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"io/ioutil"
	"log"
	"os"
)

func generateKeysAndSetEnc() {
	secret = make([]byte, 32)
	if _, err := rand.Reader.Read(secret); err != nil {
		log.Fatalln(err)
	}
	gPriv, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalln(err)
	}
	gPub := gPriv.PublicKey
	enc, err = rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		&gPub,
		secret,
		nil)
	priv, _ := os.Create("private.pem")
	_, _ = priv.Write(ExportRSAPrivateKeyAsPEM(gPriv))
}

func findDecryptAndSetSecret() {
	privateKey, _ := ioutil.ReadFile("private.pem")
	var err error
	gPriv, err = ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		panic(err)
	}
	secret, _ = gPriv.Decrypt(nil, enc, crypto.SHA512)
}
