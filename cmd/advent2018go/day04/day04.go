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
	idToMinutes := mapData(fileName)
	id := mostMinutesAsleep(idToMinutes)
	min := minuteMostAsleep(idToMinutes[id])
	return util.MustAtoi(id) * min
}

// Part2 finds the guard who is most frequently asleep on the same minute,
// and returns that minute multiplied by the guard's ID.
func Part2(fileName string) int {
	idToMinutes := mapData(fileName)
	// We need the index into minutes to build our result.
	// We need the value at that index to find our result as we loop.
	var maxSoFarMinIndex int // Index into minutes, i.e. 00..59
	maxSoFarMinValue := -1   // Value at that index, i.e. minutes[maxSoFarMin]
	var maxSoFarID string
	for id, mins := range idToMinutes {
		min := minuteMostAsleep(mins)
		if mins[min] > maxSoFarMinValue {
			maxSoFarMinValue = mins[min]
			maxSoFarMinIndex = min
			maxSoFarID = id
		}
	}
	return util.MustAtoi(maxSoFarID) * maxSoFarMinIndex
}

func mapData(fileName string) map[string]minutes {
	data := util.MustLoadData(fileName)
	sort.Strings(data)
	idToMinutes := map[string]minutes{}
	var id, fallsAsleep, wakesUp string
	var captures []string
	for _, s := range data {
		if captures = reBeginsShift.FindStringSubmatch(s); captures != nil {
			id = captures[1]
		} else if captures = reFallsAsleep.FindStringSubmatch(s); captures != nil {
			fallsAsleep = captures[1]
		} else if captures = reWakesUp.FindStringSubmatch(s); captures != nil {
			wakesUp = captures[1]
			logNap(idToMinutes, id, fallsAsleep, wakesUp)
		} else {
			log.Fatal(fmt.Errorf("Unexpected data line: %s", s))
		}
	}
	return idToMinutes
}

func logNap(idToMinutes map[string]minutes, id string, fallsAsleep, wakesUp string) {
	for i := util.MustAtoi(fallsAsleep); i < util.MustAtoi(wakesUp); i++ {
		mins := idToMinutes[id]
		mins[i]++
		idToMinutes[id] = mins
	}
}

// Return the id of the guard who spends the most minutes asleep.
func mostMinutesAsleep(idToMinutes map[string]minutes) string {
	maxSoFar := -1
	var maxSoFarID string
	for id, mins := range idToMinutes {
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
// This is the max value in the array.
func minuteMostAsleep(mins minutes) int {
	maxSoFar := -1
	var maxSoFarMin int
	for i, n := range mins {
		if n > maxSoFar {
			maxSoFar = n
			maxSoFarMin = i
		}
	}
	return maxSoFarMin
}
