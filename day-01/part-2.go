package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	// Assumes current working directory is `day-01/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	frequencyChangeList := strings.Split(string(fileContent), "\n")

	var resultingFrequency float64
	historicalFrequencyList := make(map[float64]bool)
	resultingFrequency = determineRepeatedFrequency(resultingFrequency, frequencyChangeList, historicalFrequencyList)

	fmt.Printf("DUPE: %v\n", resultingFrequency)
	fmt.Printf("TIMES RUN: %v\n", recurseCounter)
}

func determineRepeatedFrequency(currentFrequency float64, frequencyChangeList []string, historicalFrequencyList map[float64]bool) float64 {
	var duplicatedFrequency float64
	for _, frequencyChangeString := range frequencyChangeList {
		if frequencyChangeString == "" {
			continue
		}

		frequencyChange, err := strconv.ParseFloat(frequencyChangeString, 64)
		if err != nil {
			fmt.Println(err)
		}

		currentFrequency += frequencyChange
		if _, ok := historicalFrequencyList[currentFrequency]; !ok {
			historicalFrequencyList[currentFrequency] = true
		} else {
			duplicatedFrequency = currentFrequency
			break
		}
	}

	if duplicatedFrequency == 0.0 {
		currentFrequency = determineRepeatedFrequency(currentFrequency, frequencyChangeList, historicalFrequencyList)
	}
	recurseCounter++
	return currentFrequency
}
