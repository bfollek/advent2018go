package day05

import (
	"unicode"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

// Part1 returns the length of the reacted polymer.
func Part1(fileName string) int {
	polymer := util.MustLoadString(fileName)
	reacted := runReaction(polymer)
	return len(reacted)
}

// Part2 returns the length of the improved polymer.
func Part2(fileName string) int {
	_ = fileName
	return 2
}

func runReaction(polymer string) string {
	if polymer == "" {
		return "" // Guard against edge case
	}
	units := []rune(polymer)
	rv := []rune{units[0]}
	for i := 1; i < len(units); i++ {
		nextUnit := units[i]
		lastIndex := len(rv) - 1
		lastUnit := rv[lastIndex]
		if react(nextUnit, lastUnit) {
			rv = rv[:lastIndex] // Drop last unit from rv
		} else {
			rv = append(rv, nextUnit) // Add next unit to rv
		}
	}
	return string(rv)
}

func react(unit1, unit2 rune) bool {
	// Same char, different cases, e.g. a, A
	return (unit1 != unit2) && (unicode.ToUpper(unit1) == unicode.ToUpper(unit2))
}
