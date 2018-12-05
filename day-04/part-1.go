package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func part1() {
	// Assumes current working directory is `day-03/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	guardLogs := strings.Split(string(fileContent), "\n")
	sort.Strings(guardLogs)

	var selectedGuardID int
	var minutesSlept time.Time

	sleepSchedule := map[int][]int{}

	for _, guardLog := range guardLogs {
		if guardLog == "" {
			continue
		}

		dateTime, _ := time.Parse("2006-01-02 15:04", strings.Split(guardLog[1:], "]")[0])
		logText := strings.Split(guardLog, "] ")[1]

		if strings.Contains(logText, "Guard") {
			selectedGuardID, _ = strconv.Atoi(strings.Split(strings.Split(logText, "#")[1], " ")[0])
		} else if strings.Contains(logText, "falls asleep") {
			minutesSlept = dateTime
		} else if strings.Contains(logText, "wakes up") {
			for duration := minutesSlept; duration.Before(dateTime); duration = duration.Add(time.Minute) {
				sleepSchedule[selectedGuardID] = append(sleepSchedule[selectedGuardID], duration.Minute())
			}
		}
	}

	var highestMinutes, sleepiestGuard int
	for guardID, minuteRange := range sleepSchedule {
		if len(minuteRange) > highestMinutes {
			highestMinutes = len(minuteRange)
			sleepiestGuard = guardID
		}
	}

	minuteTracker := map[int]int{}
	for _, minuteSlept := range sleepSchedule[sleepiestGuard] {
		minuteTracker[minuteSlept]++
	}

	var bestMinute, highestOccurrences int
	for minute, occurence := range minuteTracker {
		if occurence > highestOccurrences {
			bestMinute = minute
			highestOccurrences = occurence
		}
	}

	fmt.Println("BEST GUARD: ", sleepiestGuard)
	fmt.Println("BEST MINUTE: ", bestMinute)
	fmt.Println("OCCURRENCES: ", highestOccurrences)
}
