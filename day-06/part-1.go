package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Coordinate struct {
	ID int
	x  int
	y  int
}

func part1() {
	// Assumes current working directory is `day-03/`!
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

	mapByClosest := map[string]int{}
	infiniteIDs := map[int]bool{}
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			closestCoord := findClosestCoord(x, y, coordinateList)
			if closestCoord.ID == -1 {
				continue
			}
			currentCoordString := fmt.Sprintf("%v,%v", x, y)
			mapByClosest[currentCoordString] = closestCoord.ID
			if findClosestCoord(x+10000, y, coordinateList) == closestCoord ||
				findClosestCoord(x, y+10000, coordinateList) == closestCoord {
				infiniteIDs[closestCoord.ID] = true
			}
		}
	}

	locationCount := map[int]int{}
	for _, closestCoord := range mapByClosest {
		if _, ok := infiniteIDs[closestCoord]; ok {
			continue
		}
		locationCount[closestCoord]++
	}

	largestLocation := 0
	var locationID int
	for ID, count := range locationCount {
		if count > largestLocation {
			largestLocation = count
			locationID = ID
		}
	}

	fmt.Println(locationID, largestLocation)
}

func findClosestCoord(x, y int, coordinateList map[int]Coordinate) Coordinate {
	shortestDistance := 10000000
	var closestCoord Coordinate

	for id, coordinate := range coordinateList {
		xDist := math.Abs(float64(coordinate.x - x))
		yDist := math.Abs(float64(coordinate.y - y))

		if int(xDist+yDist) < shortestDistance {
			shortestDistance = int(xDist + yDist)
			closestCoord = coordinateList[id]
		} else if int(xDist+yDist) == shortestDistance {
			closestCoord = Coordinate{
				ID: -1,
				x:  0,
				y:  0,
			}
		}
	}

	return closestCoord
}
