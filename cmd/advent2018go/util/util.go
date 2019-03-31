package util

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// LoadData loads a text file into a slice of strings.
func LoadData(fileName string) ([]string, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "\n"), nil
}

// MustLoadData stops program execution if LoadData() returns an error.
func MustLoadData(fileName string) []string {
	data, err := LoadData(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
