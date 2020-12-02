package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	validValidation1 := 0
	validValidation2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`^([0-9]*)-`)
		match := re.FindStringSubmatch(line)[1]
		firstNumber, err := strconv.Atoi(match)
		checkError(err)

		re = regexp.MustCompile(`-([0-9]+)`)
		match = re.FindStringSubmatch(line)[1]
		secondNumber, err := strconv.Atoi(match)
		checkError(err)

		re = regexp.MustCompile(`([a-z]):`)
		char := re.FindStringSubmatch(line)[1]

		re = regexp.MustCompile(`: ([a-z]*)$`)
		pass := re.FindStringSubmatch(line)[1]

		countedChars := strings.Count(pass, char)
		if countedChars >= firstNumber && countedChars <= secondNumber {
			validValidation1++
		}

		firstCharMatch := false
		secondCharMatch := false
		if firstNumber <= len(pass) {
			firstCharMatch = string(pass[firstNumber-1]) == char
		}
		if firstNumber <= len(pass) {
			secondCharMatch = string(pass[secondNumber-1]) == char
		}
		if (firstCharMatch == true || secondCharMatch == true) && firstCharMatch != secondCharMatch {
			validValidation2++
		}
	}
	fmt.Println("valid for first validation rule:", validValidation1)
	fmt.Println("valid for second validation rule:", validValidation2)
}

/**
 *	HELPERS
 */
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
