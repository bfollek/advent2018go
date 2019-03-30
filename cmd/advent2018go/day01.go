package day01

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// Part1 sums the frequency changes.
func Part1(fileName string) (int, error) {
	freqs, err := loadFreqs(fileName)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, n := range freqs {
		sum += n
	}
	return sum, nil
}

// Part2 finds the first repeated frequency.
func Part2(fileName string) (int, error) {
	freqs, err := loadFreqs(fileName)
	if err != nil {
		return 0, err
	}
	sum := 0
	seen := map[int]bool{sum: true}
	// We may have to cycle through the frequencies more than once
	for {
		for _, n := range freqs {
			sum += n
			if seen[sum] {
				return sum, nil
			}
			seen[sum] = true
		}
	}
}

func loadFreqs(fileName string) ([]int, error) {
	ss, err := loadData(fileName)
	if err != nil {
		return nil, err
	}
	freqs := []int{}
	for _, s := range ss {
		var i int
		if i, err = strconv.Atoi(s); err != nil {
			return nil, err
		}
		freqs = append(freqs, i)
	}
	return freqs, nil
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
