package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2() {
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	polymerUnits := strings.Split(string(fileContent), "")

	improvedProcessor(polymerUnits)
}

func improvedProcessor(polymerUnits []string) {
	shortestLength := len(polymerUnits)
	polymerString := strings.Join(polymerUnits, "")
	letters := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, letter := range letters {
		modifiedPolymer := strings.Replace(polymerString, letter, "", -1)
		modifiedPolymer = strings.Replace(modifiedPolymer, strings.ToUpper(letter), "", -1)
		modifiedUnits := strings.Split(modifiedPolymer, "")
		processedUnits := processUnits(modifiedUnits)
		if len(processedUnits) < shortestLength {
			shortestLength = len(processedUnits)
		}
	}

	fmt.Println("SHORTEST: ", shortestLength)
}
