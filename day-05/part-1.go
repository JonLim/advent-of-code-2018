package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	// Assumes current working directory is `day-03/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	polymerUnits := strings.Split(string(fileContent), "")

	currentLength := len(polymerUnits)
	trimmedPolymerSequence := trimPolymerSequence(polymerUnits)
	counter := 0
	fmt.Println("CURR: ", currentLength, ", TRIM: ", len(trimmedPolymerSequence))
	for currentLength != len(trimmedPolymerSequence) {
		currentLength = len(trimmedPolymerSequence)
		trimmedPolymerSequence = trimPolymerSequence(trimmedPolymerSequence)
		counter++
		fmt.Println("CURR: ", currentLength, ", TRIM: ", len(trimmedPolymerSequence))
	}
	fmt.Println("ITER: ", counter)
	fmt.Println(len(trimmedPolymerSequence))
	// fmt.Println(currentPolymerSequence)
}

func trimPolymerSequence(polymerSequence []string) (trimmedPolymerSequence []string) {
	lastPolymerUnit := polymerSequence[0]
	trimPair := false

	for _, polymerUnit := range polymerSequence[1:] {
		trimPair = false
		if lastPolymerUnit == "" {
			lastPolymerUnit = polymerUnit
			continue
		}

		isUpperCase := polymerUnit == strings.ToUpper(polymerUnit)
		if isUpperCase {
			if lastPolymerUnit == strings.ToLower(polymerUnit) {
				trimPair = true
			}
		} else {
			if lastPolymerUnit == strings.ToUpper(polymerUnit) {
				trimPair = true
			}
		}

		if !trimPair {
			trimmedPolymerSequence = append(trimmedPolymerSequence, lastPolymerUnit)
			trimmedPolymerSequence = append(trimmedPolymerSequence, polymerUnit)
			lastPolymerUnit = ""
			continue
		}

		lastPolymerUnit = polymerUnit
	}

	return trimmedPolymerSequence
}
