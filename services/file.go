package services

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Files struct {
	FileName string
	file     *os.File
	Info     os.FileInfo
}

func (f *Files) openFile(mode string, perm os.FileMode) *os.File {
	var flag int
	if mode == "a+" {
		flag = os.O_RDWR | os.O_APPEND | os.O_CREATE
	} else if mode == "r" {
		flag = os.O_RDONLY
	} else if mode == "w" {
		flag = os.O_WRONLY | os.O_CREATE
	} else {
		flag = os.O_RDWR | os.O_CREATE
	}
	file, err := os.OpenFile(f.FileName, flag, 0777)
	if err != nil {
		log.Fatal(err)
	}
	f.file = file
	return file
}

func (f *Files) getFileInfo() os.FileInfo {
	info, err := f.file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	f.Info = info
	return info
}

func (f *Files) createFile() *os.File {
	file, err := os.Create(f.FileName)
	if err != nil {
		log.Fatal(err)
	}
	f.file = file
	return file
}

func (f *Files) IsDir() bool {
	fileInfo, err := f.file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	isDir := fileInfo.IsDir()
	return isDir
}

func (f *Files) isExists() bool {
	_, err := os.Stat(f.FileName)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (f *Files) GetFiles(mode string, perm os.FileMode) *os.File {
	if !f.isExists() {
		f.createFile()
	}
	file := f.openFile(mode, perm)
	f.getFileInfo()
	return file
}

func (f *Files) Write(str string, mode string) {
	file := f.GetFiles(mode, 777)
	n, err := io.WriteString(file, str)
	if err != nil {
		fmt.Println(n, err)
	}
	f.Close()
}

func (f *Files) Close() {
	_ = f.file.Close()
}

func (f *Files) ReadAll() string {
	_ = f.GetFiles("r", 400)
	fileSize := f.Info.Size()
	buffer := make([]byte, fileSize)
	_, err := f.file.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
	return string(buffer)
}

func (f *Files) ReadLine() []string {
	file := f.GetFiles("r", 400)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result = make([]string, 10)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
