package day03

import (
	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge,
// 2 inches from the top edge, 5 inches wide, and 4 inches tall.
type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

// Part1 does a checksum on the box IDs.
func Part1(fileName string) int {
	claims := util.MustLoadData(fileName)
	println(claims)
	return 7
}
