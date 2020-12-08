package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	fmt.Println("accumulator till infinite loop:", accumulatorTillFirstLoop(lines))
	fmt.Println("accumulator with fixed program:", accumulatorWithFixedProgram(lines))
}

func accumulatorTillFirstLoop(lines []string) int {
	accumulator := 0
	visited := []int{}
	currentIndex := 0
	for currentIndex < len(lines) {
		if contains(visited, currentIndex) {
			break
		}

		visited = append(visited, currentIndex)
		line := lines[currentIndex]
		command := strings.Split(line, " ")[0]
		strnum := strings.Split(line, " ")[1]
		num, _ := strconv.Atoi(strnum)

		if command == "acc" {
			accumulator += num
		}

		if command == "jmp" {
			currentIndex += num
			continue
		}

		currentIndex++
	}
	return accumulator
}

func accumulatorWithFixedProgram(lines []string) (accumulator int) {
	modifiedIndexs := []int{}
	fixed := false
	for !fixed {
		accumulator = 0
		visited := []int{}
		currentIndex := 0
		modified := false
		fixed = true
		for currentIndex < len(lines) {
			if contains(visited, currentIndex) {
				fixed = false
				break
			}

			visited = append(visited, currentIndex)
			line := lines[currentIndex]
			command := strings.Split(line, " ")[0]
			strnum := strings.Split(line, " ")[1]
			num, _ := strconv.Atoi(strnum)

			if command == "acc" {
				accumulator += num
			}

			if !modified && !contains(modifiedIndexs, currentIndex) {
				if command == "nop" && num > 0 {
					command = "jmp"
					modified = true
					modifiedIndexs = append(modifiedIndexs, currentIndex)
				} else if command == "jmp" {
					command = "nop"
					modified = true
					modifiedIndexs = append(modifiedIndexs, currentIndex)
				}
			}

			if command == "jmp" {
				currentIndex += num
				continue
			}

			currentIndex++
		}
	}
	return accumulator
}

func contains(collection []int, search int) bool {
	for _, c := range collection {
		if c == search {
			return true
		}
	}
	return false
}
