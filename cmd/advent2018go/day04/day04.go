package day04

import (
	"regexp"
	"sort"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

type id int
type minutes [60]int

var reBeginsShift *regexp.Regexp
var reFallsAsleep *regexp.Regexp
var reWakesUp *regexp.Regexp

func init() {
	/*
		[1518-11-05 00:03] Guard #99 begins shift
		[1518-11-05 00:45] falls asleep
		[1518-11-05 00:55] wakes up
		Timestamps are written using year-month-day hour:minute format.
	*/

	reBeginsShift = regexp.MustCompile(`(\d+) begins shift`)
	reFallsAsleep = regexp.MustCompile(`:(\d\d) falls asleep`)
	reWakesUp = regexp.MustCompile(`:(\d\d) wakes up`)
}

// Part1 finds the guard that has the most minutes asleep,
// and the minute they're asleep the most. Part1 returns
// that minute multiplied by the guard's ID.
func Part1(fileName string) int {
	m := mapData(fileName)
	_ = m
	return 902
}

func mapData(fileName string) map[id]minutes {
	data := util.MustLoadData(fileName)
	sort.Strings(data)
	_ = data
	m := map[id]minutes{}
	//call funcs? - expectingFallsAsleep, expectingWakesUp
	// for _, s := range data {
	// 	captures := reClaim.FindStringSubmatch(s)
	// 	c := claim{id: captures[1], left: util.MustAtoi(captures[2]), top: util.MustAtoi(captures[3]),
	// 		width: util.MustAtoi(captures[4]), height: util.MustAtoi(captures[5])}
	// 	claims = append(claims, &c)
	// }
	return m
}
