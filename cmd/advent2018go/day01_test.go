package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 592
	result, err := Part1("../../data/day01.txt")
	if err != nil {
		t.Errorf("Unexpected error - [%s]", err)
	}
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}
