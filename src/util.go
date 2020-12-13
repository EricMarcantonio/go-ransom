package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	sig      = ".merk"
	secret   []byte
	encBytes []byte
	keyName  = "hello"
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
		newFileName = string(temp[0 : len(name)-len(sig)]) //slice off the name of the encrypted file
	} else {
		newFileName = name + sig //add the extensions
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
