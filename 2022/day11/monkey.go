package main

import (
	"fmt"
	"math"
	"strconv"
)

type monkey struct {
	Controller       *controller
	StartingItems    []int
	OperationOperant string
	OperationValue   string
	TestDivider      int
	TestPositive     int
	TestNegative     int
	Inpsections      int
}

func (m *monkey) Inspect(reducedWorryLvl bool) {
	for _, item := range m.StartingItems {
		m.Inpsections++
		new := m.Operation(item)
		if reducedWorryLvl {
			new = m.Bored(new)
		}
		m.Throw(new)
	}
	m.StartingItems = []int{}
}

func (m *monkey) Throw(item int) {
	m.Controller.Dispatch(item, m.Test(item))
}

func (m *monkey) Append(item int) {
	m.StartingItems = append(m.StartingItems, item)
}

func (m *monkey) Operation(old int) (new int) {
	val, _ := strconv.Atoi(m.OperationValue)
	if m.OperationValue == "old" {
		val = old
	}

	switch m.OperationOperant {
	case "*":
		new = old * val
	case "+":
		new = old + val
	case "/":
		new = old / val
	case "-":
		new = old - val
	}

	return new
}

func (m *monkey) Bored(old int) int {
	return int(math.Floor(float64(old) / 3))
}

func (m *monkey) Test(item int) (nextMonkey int) {
	if item%m.TestDivider == 0 {
		return m.TestPositive
	} else {
		return m.TestNegative
	}
}

func (m *monkey) toString() {
	fmt.Println(m)
}
