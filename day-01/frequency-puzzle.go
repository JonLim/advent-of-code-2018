package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Assumes current working directory is `day-01/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var resultingFrequency float64
	frequencyChangeList := strings.Split(string(fileContent), "\n")
	for _, frequencyChangeString := range frequencyChangeList {
		if frequencyChangeString == "" {
			continue
		}

		frequencyChange, err := strconv.ParseFloat(frequencyChangeString, 64)
		if err != nil {
			fmt.Println(err)
		}

		resultingFrequency += frequencyChange
	}

	fmt.Printf("CALIBRATED FREQUENCY: %v\n", resultingFrequency)
}
