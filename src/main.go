package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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
		for _, file := range files{
			if !file.IsDir(){
				wg.Add(1)
				go kickOff(initAnFileStruct(file.Name()))
			}
		}
	} else {
		wg.Add(1)
		go kickOff(initAnFileStruct(*filenamePTR))
	}
	wg.Wait()
	log.Println("Finished working...")
}

func kickOff(f File){
	AES(f)
}