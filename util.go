package main

import (
	"io/ioutil"
	"log"
)

func ReadFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0644)
}
