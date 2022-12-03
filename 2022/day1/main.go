package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var elfs []int
	ec := 1 // elf counter
	for scanner.Scan() {
		line := scanner.Text()

		ei := ec - 1 // elf index

		if line != "" {
			if calories, err := strconv.Atoi(line); err == nil {
				if len(elfs) < ec {
					elfs = append(elfs, calories)
				} else {
					elfs[ei] += calories
				}
			} else {
				log.Fatal(err)
			}
		} else {
			ec++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("most calories: %d\n", elfs[0])
		topCount := 3
		topSum := 0
		for i := 0; i < topCount; i++ {
			topSum += elfs[i]
		}
		fmt.Printf("most calories of top three: %d\n", topSum)
	}
}
