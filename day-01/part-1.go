package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	recurseCounter int
)

func part1() {
	// Assumes current working directory is `day-01/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	frequencyChangeList := strings.Split(string(fileContent), "\n")
	resultingFrequency := calibrateFrequency(frequencyChangeList)

	fmt.Printf("CALIBRATED FREQUENCY: %v\n", resultingFrequency)
	fmt.Printf("TIMES RUN: %v\n", recurseCounter)
}

func calibrateFrequency(frequencyChangeList []string) float64 {
	var currentFrequency float64
	for _, frequencyChangeString := range frequencyChangeList {
		if frequencyChangeString == "" {
			continue
		}

		frequencyChange, err := strconv.ParseFloat(frequencyChangeString, 64)
		if err != nil {
			fmt.Println(err)
		}

		currentFrequency += frequencyChange
	}

	recurseCounter++
	return currentFrequency
}
