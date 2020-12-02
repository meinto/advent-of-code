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

	count := flag.Int("count", 2, "how much numbers should sum to 2020")
	search := flag.Int("seach", 2020, "how much numbers should sum to 2020")
	flag.Parse()

	loop(arr, 0, *count, []int{}, *search)
}

func loop(input []string, startIndex, count int, collection []int, search int) {
	for i := startIndex; i < len(input); i++ {
		s := input[i]
		number, _ := strconv.Atoi(s)
		if len(collection) < count-1 {
			loop(input, i+1, count, append(collection, number), search)
		} else {
			tmp := append(collection, number)
			if sum(tmp) == search {
				log.Printf(ouput(tmp, " + "), sum(tmp))
				log.Printf(ouput(tmp, " * "), multiply(tmp))
			}
		}
	}
}

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
