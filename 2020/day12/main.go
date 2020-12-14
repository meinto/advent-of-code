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
	if northSouth > 0 {
		fmt.Println("north:", northSouth)
	}
	if northSouth <= 0 {
		fmt.Println("south:", -northSouth)
	}
	if eastWest > 0 {
		fmt.Println("west:", eastWest)
	}
	if eastWest <= 0 {
		fmt.Println("east:", -eastWest)
	}
	fmt.Println("sum:", math.Abs(float64(northSouth))+math.Abs(float64(eastWest)))
}
