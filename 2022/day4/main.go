package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fullyOverlappingCounter := 0
	partiallyOverlappingCounter := 0
	for scanner.Scan() {
		text := scanner.Text()
		pair := strings.Split(text, ",")
		from1, to1, sections1 := getSections(pair[0])
		from2, to2, sections2 := getSections(pair[1])

		fullyOverlapping := false
		partiallyOverlapping := false
		if len(sections1) > len(sections2) {
			fullyOverlapping = isFullyOverlapping(from1, from2, to1, to2)
			partiallyOverlapping = isPartiallyOverlapping(from1, from2, to1, to2)
		} else {
			fullyOverlapping = isFullyOverlapping(from2, from1, to2, to1)
			partiallyOverlapping = isPartiallyOverlapping(from2, from1, to2, to1)
		}

		if fullyOverlapping {
			fullyOverlappingCounter++
		}

		if partiallyOverlapping {
			partiallyOverlappingCounter++
		}
	}

	fmt.Printf("result first puzzle: fully overlapping: %d\n", fullyOverlappingCounter)
	fmt.Printf("result first puzzle: partially overlapping: %d\n", partiallyOverlappingCounter)
}

func getSections(sectionDefinition string) (from, to int, sections []int) {
	tuple := strings.Split(sectionDefinition, "-")
	from, _ = strconv.Atoi(tuple[0])
	to, _ = strconv.Atoi(tuple[1])
	for i := from; i <= to; i++ {
		sections = append(sections, i)
	}
	return from, to, sections
}

func isFullyOverlapping(from1, from2, to1, to2 int) bool {
	if from1 <= from2 && to1 >= to2 {
		return true
	}
	return false
}

func isPartiallyOverlapping(from1, from2, to1, to2 int) bool {
	if (from1 <= from2 && to1 >= from2) || (from1 <= to2 && to1 >= to2) {
		return true
	}
	return false
}
