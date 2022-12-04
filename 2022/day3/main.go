package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet += strings.ToUpper(alphabet)

	fmt.Printf("first result: %d\n", priorityOfRucksacks(alphabet))
	fmt.Printf("second result: %d\n", priorityOfGroups(alphabet))
}

func openFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func priorityOfRucksacks(alphabet string) (result int) {
	file := openFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		halfIndex := len(line) / 2
		compartment1 := line[:halfIndex]
		compartment2 := line[halfIndex:]

		found := false
		for _, itemType1 := range compartment1 {
			if found {
				break
			}
			for _, itemType2 := range compartment2 {
				if found {
					break
				}
				if itemType1 == itemType2 && !found {
					result += strings.IndexRune(alphabet, itemType1) + 1
					found = true
				}
			}
		}
	}
	return result
}

func priorityOfGroups(alphabet string) (result int) {
	file := openFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	groupSize := 3
	var group []string
	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, line)
		if len(group) == groupSize {
			groupOccurences := make(map[rune]int)
			for _, rucksack := range group {
				rucksackOccurences := make(map[rune]int)
				for _, r := range rucksack {
					if _, ok := rucksackOccurences[r]; !ok {
						groupOccurences[r]++
						if groupOccurences[r] == groupSize {
							result += strings.IndexRune(alphabet, r) + 1
						}
					}
					rucksackOccurences[r] = 1
				}
			}
			group = group[:0]
		}
	}
	return result
}
