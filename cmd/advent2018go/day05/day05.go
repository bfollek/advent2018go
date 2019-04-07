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
	reacted := []rune{}
	for _, nextUnit := range polymer {
		if len(reacted) > 0 {
			lastIndex := len(reacted) - 1
			lastUnit := reacted[lastIndex]
			if react(nextUnit, lastUnit) {
				reacted = reacted[:lastIndex] // Delete last unit from reacted
				continue
			}
		}
		// If we get here, either there was no reaction or reacted is empty.
		// (reacted can be empty more than once because we delete from it.)
		// In either case, add next unit to reacted.
		reacted = append(reacted, nextUnit)
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
