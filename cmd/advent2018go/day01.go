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

// Part2 finds the first repeated frequency.
func Part2(fileName string) (int, error) {
	sum := 0
	seen := map[int]bool{sum: true}
	numsAsStrings, err := loadData(fileName)
	if err != nil {
		return 0, err
	}
	nums := []int{}
	for _, s := range numsAsStrings {
		var i int
		if i, err = strconv.Atoi(s); err != nil {
			return 0, err
		}
		nums = append(nums, i)
	}
	// We may have to cycle through the frequencies more than once
	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				return sum, nil
			}
			seen[sum] = true
		}
	}
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
