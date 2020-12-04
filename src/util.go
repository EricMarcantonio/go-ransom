package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	sig    = ".merk"
	secret []byte
)

type File struct {
	filename    string
	newFileName string
	fileBytes   []byte
	newFile     *os.File
}

func (f File) deleteOldFile() {
	err := os.Remove(f.filename)
	CheckErr(err)
}

func initAnFileStruct(name string) File {
	var newFileName string
	if strings.HasSuffix(name, sig) {
		temp := []rune(name)
		newFileName = string(temp[0 : len(name)-len(sig)])
	} else {
		newFileName = name + sig
	}
	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Fatalln(err)
	}
	fileBytes, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}
	return File{
		filename:    name,
		newFileName: newFileName,
		fileBytes:   fileBytes,
		newFile:     newFile,
	}
}
