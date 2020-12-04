package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	log.Println("Started working...")

	filenamePTR := flag.String("f", "", "The name of the file/dir")
	flag.Parse()

	fi, err := os.Stat(*filenamePTR)
	CheckErr(err)
	if fi.IsDir() {
		_ = os.Chdir(*filenamePTR)
		files, err := ioutil.ReadDir(*filenamePTR)
		CheckErr(err)
		for _, file := range files {

			if !file.IsDir() {
				if !strings.HasSuffix(file.Name(), sig) {
					generateKeysAndSetEnc()
				}
				wg.Add(1)
				go AES(initAnFileStruct(file.Name()))
			}
		}
	} else {
		if !strings.HasSuffix(*filenamePTR, sig) {
			generateKeysAndSetEnc()
		}
		wg.Add(1)
		go AES(initAnFileStruct(*filenamePTR))
	}
	wg.Wait()
	log.Println("Finished working...")
}
