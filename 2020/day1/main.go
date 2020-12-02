package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	arr := strings.Split(string(content), "\n")

	count := flag.Int("count", 2, "how much numbers should sum up to searched number")
	search := flag.Int("search", 2020, "search number")
	flag.Parse()

	matching := collectMatching(arr, 0, *count, *search, []int{})

	for _, match := range matching {
		log.Printf(ouput(match, " + "), sum(match))
		log.Printf(ouput(match, " * "), multiply(match))
	}
}

func collectMatching(input []string, startIndex, count, search int, collection []int) (matching [][]int) {
	for i := startIndex; i < len(input); i++ {
		number, _ := strconv.Atoi(input[i])
		if len(collection)+1 < count {
			matching = append(
				matching,
				collectMatching(input, i+1, count, search, append(collection, number))...,
			)
		} else {
			collection := append(collection, number)
			if sum(collection) == search {
				matching = append(matching, collection)
			}
		}
	}
	return matching
}

/**
 *	HELPERS
 */
func sum(input []int) (result int) {
	for _, i := range input {
		result += i
	}
	return result
}

func multiply(input []int) (result int) {
	result = 1
	for _, i := range input {
		result *= i
	}
	return result
}

func ouput(input []int, seperator string) (template string) {
	numbers := []string{}
	for _, i := range input {
		numbers = append(numbers, strconv.Itoa(i))
	}
	template = strings.Join(numbers, seperator)
	return template + " = %d"
}
