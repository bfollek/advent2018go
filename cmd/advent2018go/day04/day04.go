package day04

import (
	"fmt"
	"log"
	"regexp"
	"sort"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

type tID string
type tMinutes [60]int

var reBeginsShift *regexp.Regexp
var reFallsAsleep *regexp.Regexp
var reWakesUp *regexp.Regexp

func init() {
	/*
		[1518-11-05 00:03] Guard #99 begins shift
		[1518-11-05 00:45] falls asleep
		[1518-11-05 00:55] wakes up
		Timestamps are written using year-month-day hour:minute format.
		Hour is always midnight, so it's irrelevant.
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
	//todo convert id to string
	return 902
}

func mapData(fileName string) map[id]minutes {
	data := util.MustLoadData(fileName)
	sort.Strings(data)
	m := map[tID]tMinutes{}
	var id, fallsAsleep, wakesUp string
	for _, s := range data {
		if captures := reBeginsShift.FindStringSubmatch(s); captures != nil {
			id = captures[1]
		} else if captures := reFallsAsleep.FindStringSubmatch(s); captures != nil {
			fallsAsleep = captures[1]
		} else if captures := reWakesUp.FindStringSubmatch(s); captures != nil {
			wakesUp = captures[1]
			logNap(m, id, fallsAsleep, wakesUp)
		} else {
			log.Fatal(fmt.Errorf("Unexpected data line: %s", s))
		}
	}
	return m
}

func logNap(m map[tID]tMinutes, id, fallsAsleep, wakesUp string) {
	for i := util.MustAtoi(fallsAsleep); i < util.MustAtoi(wakesUp); i++ {
		tmp := m[tID(id)][i]
		m[tID(id)][i] = tmp
	}
}
