package day05

import (
	"bytes"
	"unicode"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

// Part1 returns the length of the reacted polymer.
func Part1(fileName string) int {
	polymer := util.MustLoadString(fileName)
	reacted := runReaction(polymer)
	return len(reacted)
}

// Part2 returns the length of the most improved polymer.
func Part2(fileName string) int {
	results := map[rune]int{}
	polymer := util.MustLoadString(fileName)
	for _, unit := range polymer {
		unit = unicode.ToLower(unit)
		if _, ok := results[unit]; ok { // Already tried this unit
			continue
		}
		improved := removeUnit(polymer, unit)
		reacted := runReaction(improved)
		results[unit] = len(reacted)
	}
	minSoFar := len(polymer) + 1
	for _, n := range results {
		if n < minSoFar {
			minSoFar = n
		}
	}
	return minSoFar
}

func runReaction(polymer string) string {
	units := []rune(polymer)
	reacted := []rune{}
	for i := 0; i < len(units); i++ {
		nextUnit := units[i]
		// Special case when reacted is empty. This can come up more than once,
		// because we sometimes delete from reacted.
		if len(reacted) == 0 {
			reacted = append(reacted, nextUnit)
			continue
		}
		lastIndex := len(reacted) - 1
		lastUnit := reacted[lastIndex]
		if react(nextUnit, lastUnit) {
			reacted = reacted[:lastIndex] // Delete last unit from reacted
		} else {
			reacted = append(reacted, nextUnit) // Add next unit to reacted
		}
	}
	return string(reacted)
}

func react(unit1, unit2 rune) bool {
	// Same char, different cases, e.g. a, A
	return (unit1 != unit2) && (unicode.ToUpper(unit1) == unicode.ToUpper(unit2))
}

func removeUnit(polymer string, unit rune) string {
	var buf bytes.Buffer
	for _, r := range polymer {
		if unit == unicode.ToLower(r) {
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}
