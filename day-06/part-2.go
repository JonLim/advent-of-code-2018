package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2() {
	// Assumes current working directory is `day-03/`!
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	polymerUnits := strings.Split(string(fileContent), "")
	fmt.Println(polymerUnits)
}
