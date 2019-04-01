package day02

import (
	"bytes"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

// Part1 does a checksum on the box IDs.
func Part1(fileName string) int {
	cnt2 := 0
	cnt3 := 0
	ids := util.MustLoadData(fileName)
	for _, id := range ids {
		has2, has3 := checkID(id)
		if has2 {
			cnt2++
		}
		if has3 {
			cnt3++
		}
	}
	return cnt2 * cnt3
}

// checkID checks IDs for letters that appear exactly 2 or 3 times.
func checkID(id string) (bool, bool) {
	runeMap := map[rune]int{}
	for _, r := range id {
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

// Part2 finds the two box IDs that differ by just one char,
// and returns their common chars.
func Part2(fileName string) string {
	ids := util.MustLoadData(fileName)
	for _, id1 := range ids {
		for _, id2 := range ids {
			if b, commonChars := differByOne(id1, id2); b {
				return commonChars
			}
		}
	}
	return ""
}

// The second value returned is a string of the chars
// that s1 and s2 have in common. It's meaningful only
// if the first value is true.
func differByOne(s1, s2 string) (bool, string) {
	var commonChars bytes.Buffer
	var diffCnt int
	if s1 == s2 {
		return false, ""
	}
	// Work with rune slices so the indexes are reliable.
	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r1) != len(r2) {
		return false, ""
	}
	for idx, val := range r1 {
		if val != r2[idx] {
			diffCnt++
			if diffCnt > 1 {
				return false, ""
			}
		} else {
			commonChars.WriteRune(val)
		}
	}
	return true, commonChars.String()
}
