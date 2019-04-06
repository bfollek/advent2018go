package day01

import (
	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
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
	ss := util.MustLoadStringSlice(fileName)
	freqs := []int{}
	for _, s := range ss {
		i := util.MustAtoi(s)
		freqs = append(freqs, i)
	}
	return freqs, nil
}
