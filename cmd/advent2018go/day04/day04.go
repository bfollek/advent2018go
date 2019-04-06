package day04

import (
	"fmt"
	"log"
	"regexp"
	"sort"

	"github.com/bfollek/advent2018go/cmd/advent2018go/util"
)

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
		Hour is always midnight, so it's irrelevant.
	*/

	reBeginsShift = regexp.MustCompile(`(\d+) begins shift`)
	reFallsAsleep = regexp.MustCompile(`:(\d\d)] falls asleep`)
	reWakesUp = regexp.MustCompile(`:(\d\d)] wakes up`)
}

// Part1 finds the guard that has the most minutes asleep,
// and the minute they're asleep the most. Part1 returns
// that minute multiplied by the guard's ID.
func Part1(fileName string) int {
	m := mapData(fileName)
	id := mostMinutesAsleep(m)
	min := minuteMostAsleep(m[id])
	return util.MustAtoi(id) * min
}

func mapData(fileName string) map[string]minutes {
	data := util.MustLoadData(fileName)
	sort.Strings(data)
	m := map[string]minutes{}
	var id, fallsAsleep, wakesUp string
	var captures []string
	for _, s := range data {
		if captures = reBeginsShift.FindStringSubmatch(s); captures != nil {
			id = captures[1]
		} else if captures = reFallsAsleep.FindStringSubmatch(s); captures != nil {
			fallsAsleep = captures[1]
		} else if captures = reWakesUp.FindStringSubmatch(s); captures != nil {
			wakesUp = captures[1]
			logNap(m, id, fallsAsleep, wakesUp)
		} else {
			log.Fatal(fmt.Errorf("Unexpected data line: %s", s))
		}
	}
	return m
}

func logNap(m map[string]minutes, id string, fallsAsleep, wakesUp string) {
	for i := util.MustAtoi(fallsAsleep); i < util.MustAtoi(wakesUp); i++ {
		mins := m[id]
		mins[i]++
		m[id] = mins
	}
}

// Return the id of the guard who spends the most minutes asleep.
func mostMinutesAsleep(m map[string]minutes) string {
	maxSoFar := -1
	maxSoFarID := ""
	for id, mins := range m {
		sum := 0
		for _, n := range mins {
			sum += n
		}
		if sum > maxSoFar {
			maxSoFar = sum
			maxSoFarID = id
		}
	}
	return maxSoFarID
}

// Given a minutes type, return the minute where the guard was most asleep.
func minuteMostAsleep(mins minutes) int {
	maxSoFar := -1
	maxSoFarMin := -1
	for i, n := range mins {
		if n > maxSoFar {
			maxSoFar = n
			maxSoFarMin = i
		}
	}
	return maxSoFarMin
}
