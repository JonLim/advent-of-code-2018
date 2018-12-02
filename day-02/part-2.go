package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2() {
	// Assumes current working directory is `day-02/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	listOfIDs := strings.Split(string(fileContent), "\n")

	for _, boxID := range listOfIDs {
		if boxID == "" {
			continue
		}

		box1, box2 := matchBoxID(boxID, listOfIDs)
		if box1 != "" && box2 != "" {
			fmt.Printf("BOX 1: %v\nBOX 2: %v\n", box1, box2)
			break
		}
	}
}

func matchBoxID(boxID string, listOfIDs []string) (string, string) {
	boxIDChars := strings.Split(boxID, "")
	for _, boxIDToCheck := range listOfIDs {
		if boxID == boxIDToCheck || boxIDToCheck == "" {
			continue
		}

		fmt.Printf("CURRENT: %v - CHECKING: %v\n", boxID, boxIDToCheck)

		var mismatchCount int
		boxIDToCheckChars := strings.Split(boxIDToCheck, "")
		for charPos := range boxIDChars {
			if boxIDChars[charPos] != boxIDToCheckChars[charPos] {
				if mismatchCount == 0 {
					mismatchCount++
				} else if mismatchCount >= 1 {
					break
				}
			}
			if charPos == len(boxIDChars)-1 && boxIDChars[charPos] == boxIDToCheckChars[charPos] {
				return boxID, boxIDToCheck
			}
		}
	}

	return boxID, ""
}
