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
	Type          int
	State         int
	NextState     int
	AdjacentSeats []*Place
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	seats := parse(lines)

	occupationChanged := true
	for occupationChanged {
		occupationChanged = false
		for _, row := range seats {
			for _, p := range row {
				if p.Type == SEAT {
					adjacentOccupiedSeats := 0
					for _, as := range p.AdjacentSeats {
						if as.State == OCCUPIED {
							adjacentOccupiedSeats++
						}
					}
					if adjacentOccupiedSeats == 0 && p.State != OCCUPIED {
						occupationChanged = true
						p.NextState = OCCUPIED
					} else if adjacentOccupiedSeats >= 4 && p.State != EMPTY {
						occupationChanged = true
						p.NextState = EMPTY
					}
				}
			}
		}
		for _, row := range seats {
			for _, s := range row {
				s.State = s.NextState
			}
		}
	}

	occupiedSeats := 0
	for _, row := range seats {
		for _, s := range row {
			if s.State == OCCUPIED {
				occupiedSeats++
			}
		}
	}
	fmt.Println("occupied seats:", occupiedSeats)
}

func parse(lines []string) map[int]map[int]*Place {
	seats := make(map[int]map[int]*Place)
	for y, line := range lines {
		for x, r := range line {
			var adjacentSeats []*Place
			var seat *Place
			if r == 'L' {
				for yi := y - 1; yi < y+2; yi++ {
					for xi := x - 1; xi < x+2; xi++ {
						if yi >= 0 && yi < len(lines) && xi >= 0 && xi < len(line) {
							if _, ok := seats[yi][xi]; !ok {
								char := lines[yi][xi]
								if char == 'L' {
									s := &Place{Type: SEAT, NextState: EMPTY, State: EMPTY}
									if seats[yi] == nil {
										seats[yi] = make(map[int]*Place)
									}
									seats[yi][xi] = s
								}
							}
							existingSeat, exists := seats[yi][xi]
							if yi == y && xi == x {
								seat = existingSeat
							} else if exists {
								adjacentSeats = append(adjacentSeats, existingSeat)
							}
						}
					}
				}
				seat.AdjacentSeats = adjacentSeats
			} else {
				if seats[y] == nil {
					seats[y] = make(map[int]*Place)
				}
				seats[y][x] = &Place{Type: BOTTOM, State: EMPTY, NextState: EMPTY}
			}
		}
	}
	return seats
}
