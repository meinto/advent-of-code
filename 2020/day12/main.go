package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	NORTH_SOUTH = iota
	EAST_WEST
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	possibleDirections := []rune{'E', 'S', 'W', 'N'}
	northSouthStrings := map[int]string{-1: "south", 1: "north"}
	eastWestStrings := map[int]string{-1: "east", 1: "west"}

	coordinates := make(map[int]int)
	currentDirection := 'E'
	for _, line := range lines {
		action := rune(line[0])
		length, _ := strconv.Atoi(line[1:])

		if action == 'F' {
			action = currentDirection
		}

		switch action {
		case 'N':
			coordinates[NORTH_SOUTH] += length
		case 'S':
			coordinates[NORTH_SOUTH] -= length
		case 'W':
			coordinates[EAST_WEST] += length
		case 'E':
			coordinates[EAST_WEST] -= length

		case 'R':
			fallthrough
		case 'L':
			steps := length / 90
			if action == 'L' {
				steps *= -1
			}
			currentIndex := -1
			for i, possibleDirection := range possibleDirections {
				if possibleDirection == currentDirection {
					currentIndex = i
				}
			}
			modulo := steps % 4
			nextIndex := currentIndex + modulo
			if nextIndex >= len(possibleDirections) {
				nextIndex -= len(possibleDirections)
			} else if nextIndex < 0 {
				nextIndex += len(possibleDirections)
			}
			currentDirection = possibleDirections[nextIndex]
		}
	}

	northSouth := coordinates[0]
	eastWest := coordinates[1]
	fmt.Println(northSouthStrings[northSouth/abs(northSouth)], ":", abs(northSouth))
	fmt.Println(eastWestStrings[eastWest/abs(eastWest)], ":", abs(eastWest))
	fmt.Println("sum:", abs(northSouth)+abs(eastWest))
}

func abs(in int) int {
	return int(math.Abs(float64(in)))
}
