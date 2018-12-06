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
	processedUnits := processUnits(polymerUnits)

	fmt.Println("POLYMERS: ", len(processedUnits))
}

func processUnits(polymerUnits []string) []string {
	index := 0
	trimmedUnits := []string{}

	for index < len(polymerUnits) {
		currentUnit := polymerUnits[index]
		if index == len(polymerUnits)-1 {
			trimmedUnits = append(trimmedUnits, currentUnit)
			index++
			continue
		}
		nextUnit := polymerUnits[index+1]

		if currentUnit != nextUnit {
			isCurrentUpperCase := currentUnit == strings.ToUpper(currentUnit)
			isNextUpperCase := nextUnit == strings.ToUpper(nextUnit)

			if isCurrentUpperCase && !isNextUpperCase {
				if currentUnit == strings.ToUpper(nextUnit) {
					index += 2
					continue
				}
			} else if !isCurrentUpperCase && isNextUpperCase {
				if currentUnit == strings.ToLower(nextUnit) {
					index += 2
					continue
				}
			}
		}

		trimmedUnits = append(trimmedUnits, currentUnit)
		index++
	}

	if len(polymerUnits) == len(trimmedUnits) {
		return trimmedUnits
	}

	return processUnits(trimmedUnits)
}
