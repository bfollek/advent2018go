package day02

import "testing"

func TestPart1(t *testing.T) {
	expected := 8715
	result, err := Part1("testdata/day02.txt")
	if err != nil {
		t.Errorf("Unexpected error - [%s]", err)
	}
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}
