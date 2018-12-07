package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func part2() {
	fileContent, err := ioutil.ReadFile("puzzle-input.txt")
	if err != nil {
		fmt.Println(err)
	}

	coordinates := strings.Split(string(fileContent), "\n")

	xStart := 10000
	xEnd := 0
	yStart := 10000
	yEnd := 0
	coordinateList := map[int]Coordinate{}
	for index, coordinate := range coordinates {
		if coordinate == "" {
			continue
		}

		var x, y int
		fmt.Sscanf(coordinate, "%d, %d", &x, &y)

		if x < xStart {
			xStart = x
		}

		if x > xEnd {
			xEnd = x
		}

		if y < yStart {
			yStart = y
		}

		if y > yEnd {
			yEnd = y
		}

		coordinateList[index] = Coordinate{index, x, y}
	}

	mapByCloseEnough := map[string]bool{}
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			totalDistance := findTotalDistance(x, y, coordinateList)
			if totalDistance < 10000 {
				coordString := fmt.Sprintf("%d,%d", x, y)
				mapByCloseEnough[coordString] = true
			}
		}
	}

	fmt.Println("COORDS: (", xStart, yStart, "), (", xEnd, yEnd, ")")
	fmt.Println(len(mapByCloseEnough))
}

func findTotalDistance(x, y int, coordinateList map[int]Coordinate) int {
	totalDistance := 0
	for _, coordinate := range coordinateList {
		xDist := math.Abs(float64(coordinate.x - x))
		yDist := math.Abs(float64(coordinate.y - y))

		totalDistance += int(xDist + yDist)
	}

	return totalDistance
}
