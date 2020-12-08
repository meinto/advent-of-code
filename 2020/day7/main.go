package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	bags := parse(file)

	search := "shiny gold"
	canContainCount := 0
	for _, b := range bags {
		if b.canContainColor(search, bags) {
			canContainCount++
		}
	}

	fmt.Println(canContainCount, "bags can contain a", search, "bag")
	fmt.Println("a", search, "bag contains", bags[search].contains(bags), "bags")
}

func parse(file *os.File) (bags map[string]*Bag) {
	bags = make(map[string]*Bag)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), ".")
		rule := strings.Split(line, " bags contain ")
		color := rule[0]

		bags[color] = &Bag{
			color: color,
			count: 1,
		}

		if !strings.HasPrefix(rule[1], "no") {
			canContainRules := strings.Split(rule[1], ", ")
			canContainBags := []*Bag{}
			for _, r := range canContainRules {
				r = strings.TrimSuffix(r, "bags")
				r = strings.TrimSuffix(r, "bag")
				ruleParts := strings.SplitN(r, " ", 2)
				count, _ := strconv.Atoi(ruleParts[0])
				color := strings.TrimPrefix(ruleParts[1], " ")
				color = strings.TrimSuffix(color, " ")
				canContainBags = append(canContainBags, &Bag{
					color: color,
					count: count,
				})
			}
			bags[color].canContain = canContainBags
		}
	}
	return bags
}

type Bag struct {
	color      string
	count      int
	canContain []*Bag
}

func (b *Bag) canContainColor(color string, bagMap map[string]*Bag) (canContain bool) {
	if b.canContain != nil {
		for _, containedBag := range b.canContain {
			if containedBag.color == color {
				canContain = true
			} else if _, ok := bagMap[containedBag.color]; ok {
				canContain = bagMap[containedBag.color].canContainColor(color, bagMap)
			}
			if canContain {
				return true
			}
		}
	}
	return false
}

func (b *Bag) contains(bagMap map[string]*Bag) (contain int) {
	if b.canContain != nil {
		for _, containedBag := range b.canContain {
			contain += containedBag.count + containedBag.count*bagMap[containedBag.color].contains(bagMap)
		}
	}
	return contain
}
