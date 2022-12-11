package main

import "fmt"

func main() {
	example := NewController("./example.txt")
	example.Play(20, true, false)
	interactionsExample := example.GetMostActive()
	fmt.Printf("result example after 20 rounds (reduced worry): %d\n\n\n", interactionsExample[0]*interactionsExample[1])
	// example = NewController("./example.txt")
	// example.Play(10000, false, false)
	// interactionsExample = example.GetMostActive()
	// fmt.Printf("result example after 20 rounds (no reduced worry): %d\n\n\n", interactionsExample[0]*interactionsExample[1])

	c := NewController("./input.txt")
	c.Play(20, true, false)
	interactions := c.GetMostActive()
	fmt.Printf("result example after 20 rounds (reduced worry): %d\n\n\n", interactions[0]*interactions[1])
	// fmt.Printf("result example after 10000 rounds (no reduced worry): %d\n\n\n", interactions[0]*interactions[1])
}
