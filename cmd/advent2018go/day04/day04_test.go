package day04

import "testing"

func TestPart1(t *testing.T) {
	expected := 72925
	result := Part1("testdata/day04.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}
