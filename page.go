package main

import (
	"log"
	"os"
)

var (
	magicNumber       = []byte{'c', 'u', 'b', 'e'}
	majorVersion byte = 0
	minorVersion byte = 0

	meta = append(append(magicNumber, majorVersion), minorVersion)
	file *os.File
)

func Open(name string) {
	var err error

	if _, err = os.Stat(name); err != nil && os.IsNotExist(err) { // 文件不存在
		err := WriteFile(name, meta)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err = os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

type Page struct {
}

func (page *Page) getData() ([]byte, error) {
	return nil, nil
}

func (page *Page) saveData(data []byte) error {
	return nil
}

func getPageOffset(index uint32) uint32 {
	return 4 + 1 + 1 + 0*index
}
