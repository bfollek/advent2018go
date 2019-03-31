package util

import (
	"io/ioutil"
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
