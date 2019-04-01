package day02

import "testing"

func TestPart1(t *testing.T) {
	expected := 8715
	result := Part1("testdata/day02.txt")
	if expected != result {
		t.Errorf("Error - expected [%d], got [%d]", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := "fvstwblgqkhpuixdrnevmaycd"
	result := Part2("testdata/day02.txt")
	if expected != result {
		t.Errorf("Error - expected [%s], got [%s]", expected, result)
	}
}
