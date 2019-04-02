package day03

import "testing"

func TestPart1(t *testing.T) {
	expected := 109716
	result := Part1("testdata/day03.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := "124"
	result := Part2("testdata/day03.txt")
	if expected != result {
		t.Errorf("Error - expected [%s], got [%s]", expected, result)
	}
}
