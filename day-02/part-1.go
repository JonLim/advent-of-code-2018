package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func part1() {
	start := time.Now()
	// Assumes current working directory is `day-02/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	listOfIDs := strings.Split(string(fileContent), "\n")

	repeatCount := make(map[int]int)
	for _, boxID := range listOfIDs {
		charCount := make(map[string]int)
		chars := strings.Split(boxID, "")
		for _, char := range chars {
			if _, ok := charCount[char]; ok {
				charCount[char]++
			} else {
				charCount[char] = 1
			}
		}

		uniqueOccurence := make(map[int]int)
		for _, count := range charCount {
			if _, ok := uniqueOccurence[count]; !ok {
				uniqueOccurence[count] = 1
			} else {
				uniqueOccurence[count]++
			}
		}

		for occurrence := range uniqueOccurence {
			if _, ok := repeatCount[occurrence]; !ok {
				repeatCount[occurrence] = 1
			} else {
				repeatCount[occurrence]++
			}
		}
	}

	fmt.Printf("CHECKSUM: %v\n", repeatCount[2]*repeatCount[3])
	fmt.Printf("Total time: %s\n", time.Since(start))
}
