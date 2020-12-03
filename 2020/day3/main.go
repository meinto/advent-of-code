package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const tree = "#"
const space = "."

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("cannot read input")
	}

	slopeParts := strings.Split(string(bytes), "\n")

	slopes := []struct {
		x int
		y int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	productOfTrees := 1
	for _, slope := range slopes {
		trees := countTreesOnSlope(slopeParts, slope.x, slope.y)
		productOfTrees *= trees
		fmt.Printf("trees on slope [%d,%d]: %d\n", slope.x, slope.y, trees)
	}
	fmt.Println("Product of trees:", productOfTrees)
}

func countTreesOnSlope(slopeParts []string, x, y int) (countedTrees int) {
	currentX := x
	currentY := y

	for currentY < len(slopeParts) {
		line := slopeParts[currentY]
		char := line[currentX]

		currentX += x
		if currentX >= len(line) {
			currentX = currentX - len(line)
		}
		currentY += y

		if string(char) == tree {
			countedTrees++
		}
	}
	return countedTrees
}
