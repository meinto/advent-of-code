package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	OCCUPIED = iota
	EMPTY
)

const (
	SEAT = iota
	BOTTOM
)

type Place struct {
	Type      int
	State     int
	NextState int
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	parse(lines)

	seats := parse(lines)

	fmt.Println("occupied seats 1:", calcOccupiedSeats(lines, seats, 1, 4))
	fmt.Println("occupied seats 2:", calcOccupiedSeats(lines, seats, -1, 5))
}

func calcOccupiedSeats(lines []string, seats map[int]map[int]*Place, maxLookAround, maxAdjacentOccupiedSeats int) int {
	occupationChanged := true
	for occupationChanged {
		occupationChanged = false
		for y, line := range lines {
			for x := range line {
				p := seats[y][x]
				if p.Type == SEAT {
					lookAround := 1
					matchInDirection := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1}
					for {
						if maxLookAround > 0 && lookAround > maxLookAround {
							break
						} else if maxLookAround < 0 &&
							y+lookAround > len(lines) &&
							y-lookAround < 0 &&
							x+lookAround > len(line) &&
							x-lookAround < 0 {
							break
						}
						direction := 0
						for yi := y - lookAround; yi <= y+lookAround; yi += lookAround {
							for xi := x - lookAround; xi <= x+lookAround; xi += lookAround {
								if yi >= 0 &&
									yi < len(lines) &&
									xi >= 0 &&
									xi < len(line) &&
									(yi != y || xi != x) {
									place := seats[yi][xi]
									if place.Type == SEAT && place.State == OCCUPIED {
										if matchInDirection[direction] < 0 {
											matchInDirection[direction] = 1
										}
									} else if place.Type == SEAT && place.State == EMPTY {
										if matchInDirection[direction] < 0 {
											matchInDirection[direction] = 0
										}
									}
								}
								direction++
							}
						}
						lookAround++
					}

					adjacentOccupiedSeats := 0
					for _, match := range matchInDirection {
						if match > 0 {
							adjacentOccupiedSeats += match
						}
					}

					if adjacentOccupiedSeats == 0 && p.State != OCCUPIED {
						occupationChanged = true
						p.NextState = OCCUPIED
					} else if adjacentOccupiedSeats >= maxAdjacentOccupiedSeats && p.State != EMPTY {
						occupationChanged = true
						p.NextState = EMPTY
					}
				}
			}
		}

		// printSeats(lines, seats)

		for y, line := range lines {
			for x := range line {
				s := seats[y][x]
				s.State = s.NextState
			}
		}
	}

	occupiedSeats := 0
	for _, row := range seats {
		for _, s := range row {
			if s.State == OCCUPIED {
				occupiedSeats++

				// reset seat
				s.State = EMPTY
				s.NextState = EMPTY
			}
		}
	}
	return occupiedSeats
}

func parse(lines []string) map[int]map[int]*Place {
	seats := make(map[int]map[int]*Place)
	for y, line := range lines {
		if seats[y] == nil {
			seats[y] = make(map[int]*Place)
		}
		for x, r := range line {
			if r == 'L' {
				seats[y][x] = &Place{Type: SEAT, NextState: EMPTY, State: EMPTY}
			} else {
				seats[y][x] = &Place{Type: BOTTOM, State: EMPTY, NextState: EMPTY}
			}
		}
	}
	return seats
}

func printSeats(lines []string, seats map[int]map[int]*Place) {
	fmt.Println()
	fmt.Println()
	for y, line := range lines {
		fmt.Println()
		for x := range line {
			s := seats[y][x]
			if s.Type == SEAT && s.State == OCCUPIED {
				fmt.Print("#")
			}
			if s.Type == SEAT && s.State == EMPTY {
				fmt.Print("L")
			}
			if s.Type == BOTTOM {
				fmt.Print(".")
			}
		}
	}
}
