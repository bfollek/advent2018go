package day02

import (
	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

// Part1 does a checksum on the box IDs.
func Part1(fileName string) (int, error) {
	cnt2 := 0
	cnt3 := 0
	ids, err := util.LoadData(fileName)
	if err != nil {
		return 0, err
	}
	for _, id := range ids {
		has2, has3 := checkID(id)
		if has2 {
			cnt2++
		}
		if has3 {
			cnt3++
		}
	}
	return cnt2 * cnt3, nil
}

// checkID checks IDs for letters that appear exactly 2 or 3 times.
func checkID(id string) (bool, bool) {
	runeMap := map[rune]int{}
	for _, r := range []rune(id) {
		runeMap[r]++
	}
	var has2, has3 bool
	for _, value := range runeMap {
		switch value {
		case 2:
			has2 = true
		case 3:
			has3 = true
		}
		if has2 && has3 {
			break // No need to keep checking
		}
	}
	return has2, has3
}
