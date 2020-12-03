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

	validForFirstRule, validForSecondRule := validate(file)
	fmt.Println("valid for first validation rule:", validForFirstRule)
	fmt.Println("valid for second validation rule:", validForSecondRule)
}

func validate(file *os.File) (validForFirstRule, validForSecondRule int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstNumber, err := strconv.Atoi(findMatch(line, `^([0-9]*)-`))
		checkError(err)
		secondNumber, err := strconv.Atoi(findMatch(line, `-([0-9]+)`))
		checkError(err)
		char := findMatch(line, `([a-z]):`)
		pass := findMatch(line, `: ([a-z]*)$`)

		if firstValidationRule(pass, char, firstNumber, secondNumber) {
			validForFirstRule++
		}

		if secondValidationRule(pass, char, firstNumber, secondNumber) {
			validForSecondRule++
		}
	}
	return validForFirstRule, validForSecondRule
}

func firstValidationRule(pass, char string, min, max int) bool {
	countedChars := strings.Count(pass, char)
	if countedChars >= min && countedChars <= max {
		return true
	}
	return false
}

func secondValidationRule(pass, char string, firstIndex, secondIndex int) bool {
	firstCharMatch := false
	secondCharMatch := false
	if firstIndex <= len(pass) {
		firstCharMatch = string(pass[firstIndex-1]) == char
	}
	if secondIndex <= len(pass) {
		secondCharMatch = string(pass[secondIndex-1]) == char
	}
	if (firstCharMatch == true || secondCharMatch == true) && firstCharMatch != secondCharMatch {
		return true
	}
	return false
}

/**
 *	HELPERS
 */
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func findMatch(line, pattern string) string {
	re := regexp.MustCompile(pattern)
	if len(re.FindStringSubmatch(line)) < 2 {
		log.Fatal("cannot find submatch in line:", line, "\n", "with pattern:", pattern)
	}
	return re.FindStringSubmatch(line)[1]
}
