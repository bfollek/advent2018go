package day01

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// Part1 does Chronal Calibration.
func Part1(fileName string) (int, error) {
	numsAsStrings, err := loadData(fileName)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, s := range numsAsStrings {
		var i int
		if i, err = strconv.Atoi(s); err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}

func loadData(fileName string) ([]string, error) {
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
