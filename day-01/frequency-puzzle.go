package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	dupePtr        *bool
	recurseCounter int
)

func main() {
	// Command Line Flag, `-dupe`, to determine what our desired output is
	dupePtr = flag.Bool("dupe", false, "Flag for determining if we want to recursively look for a duplicate!")
	flag.Parse()

	// Assumes current working directory is `day-01/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	frequencyChangeList := strings.Split(string(fileContent), "\n")

	var resultingFrequency float64
	historicalFrequencyList := make(map[float64]bool)
	resultingFrequency = calibrateFrequency(resultingFrequency, frequencyChangeList, historicalFrequencyList)

	if *dupePtr {
		fmt.Printf("DUPE: %v\n", resultingFrequency)
	} else {
		fmt.Printf("CALIBRATED FREQUENCY: %v\n", resultingFrequency)
	}
	fmt.Printf("TIMES RUN: %v\n", recurseCounter)
}

func calibrateFrequency(currentFrequency float64, frequencyChangeList []string, historicalFrequencyList map[float64]bool) float64 {
	var duplicatedFrequency float64
	recurseCounter++
	for _, frequencyChangeString := range frequencyChangeList {
		if frequencyChangeString == "" {
			continue
		}

		frequencyChange, err := strconv.ParseFloat(frequencyChangeString, 64)
		if err != nil {
			fmt.Println(err)
		}

		currentFrequency += frequencyChange
		if *dupePtr {
			if _, ok := historicalFrequencyList[currentFrequency]; !ok {
				historicalFrequencyList[currentFrequency] = true
			} else {
				duplicatedFrequency = currentFrequency
				break
			}
		}
	}

	// Only recurse if `-dupe` flag is present when running
	if *dupePtr {
		if duplicatedFrequency == 0.0 {
			currentFrequency = calibrateFrequency(currentFrequency, frequencyChangeList, historicalFrequencyList)
		} else {
			return currentFrequency
		}
	}

	return currentFrequency
}
