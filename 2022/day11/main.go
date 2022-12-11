package main

import (
	"fmt"
)

func main() {
	example := NewController("./example.txt")
	example.Play(20, example.ReduceWorryLvl_1, false)
	interactionsExample := example.GetMostActive()
	fmt.Printf("result example after 20 rounds (reduced worry): %d\n", interactionsExample[0]*interactionsExample[1])
	example = NewController("./example.txt")
	example.Play(10000, example.ReduceWorryLvl_2, false)
	interactionsExample = example.GetMostActive()
	fmt.Printf("result example after 10000 rounds (no reduced worry): %d\n\n\n", interactionsExample[0]*interactionsExample[1])

	c := NewController("./input.txt")
	c.Play(20, c.ReduceWorryLvl_1, false)
	interactions := c.GetMostActive()
	fmt.Printf("result after 20 rounds (reduced worry): %d\n", interactions[0]*interactions[1])
	c = NewController("./input.txt")
	c.Play(10000, c.ReduceWorryLvl_2, false)
	interactions = c.GetMostActive()
	fmt.Printf("result after 10000 rounds (no reduced worry): %d\n\n\n", interactions[0]*interactions[1])
}
