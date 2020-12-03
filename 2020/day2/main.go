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

		firstNumber, err := strconv.Atoi(findMatch(line, `^([0-9]*)-`))
		checkError(err)
		secondNumber, err := strconv.Atoi(findMatch(line, `-([0-9]+)`))
		checkError(err)
		char := findMatch(line, `([a-z]):`)
		pass := findMatch(line, `: ([a-z]*)$`)

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

func findMatch(line, pattern string) string {
	re := regexp.MustCompile(pattern)
	if len(re.FindStringSubmatch(line)) < 2 {
		log.Fatal("cannot find submatch in line:", line, "\n", "with pattern:", pattern)
	}
	return re.FindStringSubmatch(line)[1]
}
