package day05

import "testing"

func TestRunReaction(t *testing.T) {
	type reactionTest struct {
		input    string
		expected string
	}
	var reactionTests = []reactionTest{
		{
			"dabAcCaCBAcCcaDA",
			"dabCBAcaDA",
		},
		{
			"aA",
			"",
		},
		{
			"",
			"",
		},
	}
	for _, rt := range reactionTests {
		if result := runReaction(rt.input); result != rt.expected {
			t.Errorf("runReaction: got [%s], expecting [%s]", result, rt.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 10450
	result := Part1("testdata/day05.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 4624
	result := Part2("testdata/day05.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}
