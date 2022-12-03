package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const (
		ROCK     int = 1
		PAPER        = 2
		SCISSORS     = 3

		LOOSE = 0
		DRAW  = 3
		WIN   = 6
	)
	shapes := map[string]int{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,

		// first meaning
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	gameResult := map[string]int{
		// second meaning
		"X": LOOSE,
		"Y": DRAW,
		"Z": WIN,
	}

	possibleGameResults := map[int]map[int]int{
		ROCK: map[int]int{
			ROCK:     DRAW,
			PAPER:    WIN,
			SCISSORS: LOOSE,
		},
		PAPER: map[int]int{
			PAPER:    DRAW,
			SCISSORS: WIN,
			ROCK:     LOOSE,
		},
		SCISSORS: map[int]int{
			SCISSORS: DRAW,
			ROCK:     WIN,
			PAPER:    LOOSE,
		},
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstScore := 0
	secondScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		opponentShape := shapes[parts[0]]

		// first meaning
		myShape := shapes[parts[1]]
		firstScore += myShape
		firstScore += possibleGameResults[opponentShape][myShape]

		// second meaning
		requiredGameResult := gameResult[parts[1]]
		var requiredShape int
		for shape, gameResult := range possibleGameResults[opponentShape] {
			if gameResult == requiredGameResult {
				requiredShape = shape
			}
		}
		secondScore += requiredShape
		secondScore += requiredGameResult
	}

	fmt.Printf("first result: total score: %d\n", firstScore)
	fmt.Printf("second result: total score: %d\n", secondScore)
}
