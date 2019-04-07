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
	polyRunes := []rune(polymer)
	rv := []rune{polyRunes[0]}
	for i := 1; i < len(polyRunes); i++ {
		next := polyRunes[i]
		lastIndex := len(rv) - 1
		last := rv[lastIndex]
		if react(next, last) {
			rv = rv[:lastIndex] // Drop last from rv
		} else {
			rv = append(rv, next) // Add next to rv
		}
	}
	return string(rv)
}

func react(r1, r2 rune) bool {
	// Same char, different cases, e.g. a, A
	return (r1 != r2) && (unicode.ToUpper(r1) == unicode.ToUpper(r2))
}
