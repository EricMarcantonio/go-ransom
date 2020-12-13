package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"sync"
)

func generateKeys() {
	if fileExists(keyName + "_priv.keyme") {
		return
	}
	var wg1 sync.WaitGroup
	wg1.Add(2)
	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	publicKey := privateKey.PublicKey
	go func() {
		WritePrivateKey(privateKey)
		wg1.Done()
	}()
	go func() {
		WritePublicKey(&publicKey)
		wg1.Done()
	}()
	wg1.Wait()
}

func encryptSecret(pub *rsa.PublicKey) {
	var err error
	encBytes, err = rsa.EncryptOAEP(sha512.New(), rand.Reader, pub, secret, nil)
	CheckErr(err)
}

func decryptEncBytes(priv *rsa.PrivateKey) {
	var err error
	secret, err = rsa.DecryptOAEP(sha512.New(), rand.Reader, priv, encBytes, nil)
	CheckErr(err)
}

func WritePrivateKey(priv *rsa.PrivateKey) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(priv)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create(keyName + "_priv.keyme")
	CheckErr(err)
	err = pem.Encode(privatePem, privateKeyBlock)
	CheckErr(err)
}
func WritePublicKey(pub *rsa.PublicKey) {
	publicKeyBytes := x509.MarshalPKCS1PublicKey(pub)
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create(keyName + "_pub.keyme")
	err = pem.Encode(publicPem, publicKeyBlock)
	CheckErr(err)
}

func ReadPublicKey() *rsa.PublicKey {
	file, err := ioutil.ReadFile(keyName + "_pub.keyme")
	CheckErr(err)
	pemBlock, _ := pem.Decode(file)
	pub, _ := x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	return pub

}
func ReadPrivateKey() *rsa.PrivateKey {
	file, err := ioutil.ReadFile(keyName + "_priv.keyme")
	CheckErr(err)
	pemBlock, _ := pem.Decode(file)
	priv, _ := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	return priv
}
