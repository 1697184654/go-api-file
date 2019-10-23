package services

import (
	"log"
	"os"
)

type Files struct {
	FileName string
	File     *os.File
}

func (f *Files) openFile() *os.File {
	file, err := os.Open(f.FileName)
	if err != nil {
		log.Fatal(err)
	}
	f.File = file
	return f.File
}

func (f *Files) createFile() *os.File {
	file, err := os.Create(f.FileName)
	if err != nil {
		log.Fatal(err)
	}
	f.File = file
	return f.File
}

func (f *Files) IsDir() bool {
	fileInfo, err := f.File.Stat()
	if err != nil {
		log.Fatal(err)
	}
	isDir := fileInfo.IsDir()
	return isDir
}

func (f *Files) isExists() bool {
	_, err := os.Stat(f.FileName)
	if err == nil {
		return true
	}
	log.Fatal(err)
	return false
}

func (f *Files) GetFiles() *os.File {
	if f.isExists() {
		return f.openFile()
	}
	return f.createFile()
}

func (f *Files) Write(str string) {

}
