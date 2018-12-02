package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

func part2() {
	start := time.Now()

	// Assumes current working directory is `day-02/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	listOfIDs := strings.Split(string(fileContent), "\n")

	var wg sync.WaitGroup
	wg.Add(len(listOfIDs) - 1)

	for currentIndex, boxID := range listOfIDs {
		if boxID == "" {
			continue
		}

		go func(currentIndex int, boxID string, listOfIDs []string) {
			defer wg.Done()
			boxMatch := matchBoxID(boxID, listOfIDs[currentIndex:])
			if boxMatch != "" {
				fmt.Printf("BOX 1: %v\nBOX 2: %v\n", boxID, boxMatch)
			}
		}(currentIndex, boxID, listOfIDs)
	}

	wg.Wait()

	fmt.Printf("Total time: %s\n", time.Since(start))
}

func matchBoxID(boxID string, listOfIDs []string) string {
	boxIDChars := strings.Split(boxID, "")
	for _, boxIDToCheck := range listOfIDs {
		if boxID == boxIDToCheck || boxIDToCheck == "" {
			continue
		}

		var hasMismatch bool
		var hasMoreThanOneMismatch bool
		boxIDToCheckChars := strings.Split(boxIDToCheck, "")
		for charPos := range boxIDChars {
			if boxIDChars[charPos] != boxIDToCheckChars[charPos] {
				if !hasMismatch {
					hasMismatch = true
				} else if hasMismatch {
					hasMoreThanOneMismatch = true
					break
				}
			}
		}
		if !hasMoreThanOneMismatch {
			return boxIDToCheck
		}
	}
	return ""
}
