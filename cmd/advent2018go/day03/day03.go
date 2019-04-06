package day03

import (
	"regexp"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

type claim struct {
	id       string
	left     int
	top      int
	width    int
	height   int
	overlaps bool // Does this claim overlap any other claims?
}

type inch struct {
	x int
	y int
}

var reClaim *regexp.Regexp

func init() {
	// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge,
	// 2 inches from the top edge, 5 inches wide, and 4 inches tall.
	reClaim = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
}

// Part1 counts how many square inches of fabric are within two or more claims.
func Part1(fileName string) int {
	claims := parseClaims(fileName)
	m := mapClaims(claims)
	cnt := 0
	for _, values := range m {
		if len(values) > 1 {
			cnt++
		}
	}
	return cnt
}

func parseClaims(fileName string) []*claim {
	data := util.MustLoadStringSlice(fileName)
	claims := []*claim{}
	for _, s := range data {
		captures := reClaim.FindStringSubmatch(s)
		c := claim{id: captures[1], left: util.MustAtoi(captures[2]), top: util.MustAtoi(captures[3]),
			width: util.MustAtoi(captures[4]), height: util.MustAtoi(captures[5])}
		claims = append(claims, &c)
	}
	return claims
}

// mapClaims creates a map where each key is a square inch from a claim,
// and each value is a slice of the claims that include that square inch.
func mapClaims(claims []*claim) map[inch][]*claim {
	m := map[inch][]*claim{}
	for _, c := range claims {
		for x := c.left; x < c.left+c.width; x++ {
			for y := c.top; y < c.top+c.height; y++ {
				in := inch{x, y}
				m[in] = append(m[in], c)
			}
		}
	}
	return m
}

// Part2 finds the one claim that doesn't overlap any other, and returns its ID.
func Part2(fileName string) string {
	claims := parseClaims(fileName)
	m := mapClaims(claims)
	for _, values := range m {
		// If an inch has multiple claims, set all those claims to overlaps = true.
		if len(values) > 1 {
			for _, c := range values {
				c.overlaps = true
			}
		}
	}
	// The claim that doesn't overlap is the winner.
	for _, c := range claims {
		if !c.overlaps {
			return c.id
		}
	}
	return ""
}
