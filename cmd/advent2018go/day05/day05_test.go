package day05

import "testing"

func TestRunReaction(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	expected := "dabCBAcaDA"
	result := runReaction(input)
	if expected != result {
		t.Errorf("Error - expected [%s], got [%s]", expected, result)
	}
}

func TestPart1(t *testing.T) {
	expected := 10450
	result := Part1("testdata/day05.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 4624
// 	result := Part2("testdata/day05.txt")
// 	if expected != result {
// 		t.Errorf("Error - expected [%d], got [%d]", expected, result)
// 	}
// }
