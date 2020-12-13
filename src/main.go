package main

import (
	"crypto/rand"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	log.Println("Started working...")
	start := time.Now()
	numOfFiles := 0
	filenamePTR := flag.String("f", "", "The name of the file/dir")
	flag.Parse()

	//Generate random bytes and set the secret to it
	secret = make([]byte, 32)
	_, _ = rand.Read(secret)

	fi, err := os.Stat(*filenamePTR)
	CheckErr(err)
	if fi.IsDir() {
		_ = os.Chdir(*filenamePTR)
		generateKeys()
		encryptSecret(ReadPublicKey())
		decryptEncBytes(ReadPrivateKey())
		files, err := ioutil.ReadDir(*filenamePTR)
		CheckErr(err)
		for _, file := range files {
			if !file.IsDir() {
				numOfFiles += 1
				wg.Add(1)
				go AES(initAnFileStruct(file.Name()))
			}
		}
	} else {
		generateKeys()
		encryptSecret(ReadPublicKey())
		decryptEncBytes(ReadPrivateKey())
		numOfFiles += 1
		wg.Add(1)
		go AES(initAnFileStruct(*filenamePTR))
	}
	wg.Wait()
	log.Println("Finished", numOfFiles, "files in", time.Since(start).Milliseconds())
}
