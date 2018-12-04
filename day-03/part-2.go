package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func part2() {
	// Assumes current working directory is `day-03/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	listOfClaims := strings.Split(string(fileContent), "\n")

	var processedClaims []fabricClaim
	for _, claim := range listOfClaims {
		if claim == "" {
			continue
		}

		// Function used for FieldsFunc to pull out only strings that we care about
		fields := func(c rune) bool {
			return !unicode.IsNumber(c)
		}

		// Produces an array with just the numbers from the puzzle input
		// Sample claim => #1 @ 749,666: 27x15
		// Becomes 		=> [1 749 666 27 15]
		claimData := strings.FieldsFunc(claim, fields)

		var claimDataNumbers []int
		for _, data := range claimData {
			dataInt, err := strconv.Atoi(data)
			if err != nil {
				panic(err)
			}
			claimDataNumbers = append(claimDataNumbers, dataInt)
		}

		processedClaim := fabricClaim{
			ID:        claimDataNumbers[0],
			leftStart: claimDataNumbers[1],
			topStart:  claimDataNumbers[2],
			width:     claimDataNumbers[3],
			height:    claimDataNumbers[4],
		}

		processedClaims = append(processedClaims, processedClaim)
	}

	fabric := make(map[squareInch]int)

	for _, claim := range processedClaims {
		for x := claim.leftStart + 1; x < claim.leftStart+1+claim.width; x++ {
			for y := claim.topStart + 1; y < claim.topStart+1+claim.height; y++ {
				squareInch := squareInch{x, y}
				fabric[squareInch]++
			}
		}
	}

	for _, claim := range processedClaims {
		hasOverlap := false
		for x := claim.leftStart + 1; x < claim.leftStart+1+claim.width; x++ {
			for y := claim.topStart + 1; y < claim.topStart+1+claim.height; y++ {
				squareInch := squareInch{x, y}
				if fabric[squareInch] > 1 {
					hasOverlap = true
				}
			}
		}
		if !hasOverlap {
			fmt.Println(claim.ID)
		}
	}
}
