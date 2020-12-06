package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("cannot read file")
	}
	defer f.Close()

	rows := 128
	seats := 8
	plainRows := create(rows)
	seatsPerRow := create(seats)

	seatIDs := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		remainingRows := plainRows
		for _, char := range line[:7] {
			remainingRows, err = reduce(remainingRows, char)
			if err != nil {
				log.Fatal(err)
			}
		}

		remainingSeats := seatsPerRow
		for _, char := range line[7:] {
			remainingSeats, err = reduce(remainingSeats, char)
			if err != nil {
				log.Fatal(err)
			}
		}

		row := remainingRows[0]
		seat := remainingSeats[0]
		seatIDs = append(seatIDs, row*8+seat)
	}

	fmt.Println("max seat id:", max(seatIDs))
	fmt.Println("my seat is:", findMySeat(rows, seats, seatIDs))
}

func findMySeat(plainRows, seatsPerRow int, takenSeatIDs []int) (mySeatID int) {
	seats := []int{}
	for r := 0; r < plainRows; r++ {
		for s := 0; s < seatsPerRow; s++ {
			seatID := r*8 + s
			seats = append(seats, seatID)
			if len(seats) >= 3 {
				seatBefore := isSeatTaken(seats[len(seats)-3], takenSeatIDs)
				myPossibleSeat := isSeatTaken(seats[len(seats)-2], takenSeatIDs)
				seatAfter := isSeatTaken(seats[len(seats)-1], takenSeatIDs)

				if seatBefore && !myPossibleSeat && seatAfter {
					return seats[len(seats)-2]
				}
			}
		}
	}
	return -1
}

func isSeatTaken(seatID int, takenSeatIDs []int) bool {
	for _, takenSeatID := range takenSeatIDs {
		if takenSeatID == seatID {
			return true
		}
	}
	return false
}

func reduce(input []int, r rune) ([]int, error) {
	if len(input)%2 != 0 {
		return []int{}, errors.New("point number index")
	}
	index := len(input) / 2

	switch r {
	case 'F':
		fallthrough
	case 'L':
		return input[:index], nil
	case 'B':
		fallthrough
	case 'R':
		return input[index:], nil
	default:
		return []int{}, errors.New("missing control charactor")
	}
}

func create(size int) (output []int) {
	for i := 0; i < size; i++ {
		output = append(output, i)
	}
	return output
}

func max(in []int) (max int) {
	for _, i := range in {
		if i > max {
			max = i
		}
	}
	return max
}
