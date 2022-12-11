package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type controller struct {
	Monkeys []*monkey
	Round   int
}

func (c *controller) Play(maxRounds int, reduceWorryLvl func(int) int, logRounds bool) {
	if c.Round < maxRounds {
		c.Round++
		for _, monkey := range c.Monkeys {
			monkey.Inspect(reduceWorryLvl)
		}
		if logRounds {
			fmt.Printf("round %d\n", c.Round)
			for _, monkey := range c.Monkeys {
				monkey.toString()
			}
		}
		c.Play(maxRounds, reduceWorryLvl, logRounds)
	}
}

func (c *controller) Dispatch(item int, nextMonkey int) {
	c.Monkeys[nextMonkey].Append(item)
}

func (c *controller) GetMostActive() (inspections []int) {
	for _, monkey := range c.Monkeys {
		inspections = append(inspections, monkey.Inpsections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections
}

func (c *controller) ReduceWorryLvl_1(old int) int {
	return int(math.Floor(float64(old) / 3))
}

func (c *controller) ReduceWorryLvl_2(old int) int {
	divisor := 1
	for _, monkey := range c.Monkeys {
		divisor *= monkey.TestDivider
	}
	return old % divisor
}

func NewController(filepath string) (c *controller) {
	c = &controller{Round: 0}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentMonkey := -1
	var monkeys []*monkey
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)

		if strings.HasPrefix(text, "Monkey") {
			currentMonkey++
			monkeys = append(monkeys, &monkey{Controller: c})
		}

		if strings.HasPrefix(text, "Starting items") {
			itemsString := strings.TrimSpace(strings.Split(text, ":")[1])
			itemStringList := strings.Split(itemsString, ", ")
			var itemList []int
			for _, itemString := range itemStringList {
				item, _ := strconv.Atoi(itemString)
				itemList = append(itemList, item)
			}
			monkeys[currentMonkey].StartingItems = itemList
		}

		if strings.HasPrefix(text, "Operation") {
			r := regexp.MustCompile(`new\s=\s\w+\s([-+\*\/])\s(.*)`)
			matches := r.FindStringSubmatch(text)
			monkeys[currentMonkey].OperationOperant = matches[1]
			monkeys[currentMonkey].OperationValue = matches[2]
		}

		if strings.HasPrefix(text, "Test") {
			r := regexp.MustCompile(`divisible by ([0-9]+)`)
			matches := r.FindStringSubmatch(text)
			devider, _ := strconv.Atoi(matches[1])
			monkeys[currentMonkey].TestDivider = devider
		}

		if strings.HasPrefix(text, "If true") {
			r := regexp.MustCompile(`throw to monkey ([0-9]+)`)
			matches := r.FindStringSubmatch(text)
			devider, _ := strconv.Atoi(matches[1])
			monkeys[currentMonkey].TestPositive = devider
		}

		if strings.HasPrefix(text, "If false") {
			r := regexp.MustCompile(`throw to monkey ([0-9]+)`)
			matches := r.FindStringSubmatch(text)
			devider, _ := strconv.Atoi(matches[1])
			monkeys[currentMonkey].TestNegative = devider
		}
	}

	c.Monkeys = monkeys
	return c
}
