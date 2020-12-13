package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	adapters := getAdapters(lines)
	adapterChain := NewAdapterChain(adapters).sortAsc()

	deviceJoltDiff := 3
	deviceJolt := adapterChain.maxJolt() + deviceJoltDiff

	adapterChain.calcDifferences(0, deviceJolt)

	distinctChainCount := adapterChain.calcDistinctConnections(0, deviceJolt)

	fmt.Println(adapterChain.getDiffCount(1), "times 1-jolt difference")
	fmt.Println(adapterChain.getDiffCount(3), "times 3-jolt difference")
	fmt.Println("product:", adapterChain.getDiffCount(1)*(adapterChain.getDiffCount(3)))
	fmt.Println("distinct chains:", distinctChainCount)
}

type AdapterChain struct {
	adapters    []Adapter
	differences map[int]int
	maxDistance int
}

func NewAdapterChain(adapters []Adapter) *AdapterChain {
	return &AdapterChain{adapters, make(map[int]int), 3}
}

func (ac *AdapterChain) calcDistinctConnections(inputJolt, deviceJolt int) int {
	connectionMap := make(map[int]int)
	connectionMap[deviceJolt] = 0
	ac.sortDesc()
	for _, a := range ac.adapters {
		connectionCount := ac.calcConnectionsForAdapter(a.outputJolt, connectionMap)
		connectionMap[a.outputJolt] = connectionCount
	}
	connectionCount := ac.calcConnectionsForAdapter(inputJolt, connectionMap)
	return connectionCount
}

func (ac *AdapterChain) calcConnectionsForAdapter(outputJolt int, connectionMap map[int]int) (connectionCount int) {
	for i := 1; i <= ac.maxDistance; i++ {
		if connections, ok := connectionMap[outputJolt+i]; ok {
			if connections == 0 {
				connectionCount = 1
			}
			connectionCount += connections
		}
	}
	return connectionCount
}

func (ac *AdapterChain) calcDifferences(inputJolt, outputJolt int) {
	for i, curr := range ac.adapters {
		diff := -1
		if i == 0 {
			diff = curr.outputJolt - inputJolt
		} else {
			prev := ac.adapters[i-1]
			diff = curr.outputJolt - prev.outputJolt
		}
		ac.incrementDiff(diff)
	}
	deviceDiff := outputJolt - ac.maxJolt()
	ac.incrementDiff(deviceDiff)
}

func (ac *AdapterChain) incrementDiff(diff int) {
	if _, ok := ac.differences[diff]; !ok {
		ac.differences[diff] = 1
	} else {
		ac.differences[diff]++
	}
}

func (ac *AdapterChain) getDiffCount(diff int) int {
	return ac.differences[diff]
}

func (ac *AdapterChain) sortAsc() *AdapterChain {
	sort.Slice(ac.adapters, func(a, b int) bool {
		return ac.adapters[a].outputJolt < ac.adapters[b].outputJolt
	})
	return ac
}

func (ac *AdapterChain) sortDesc() *AdapterChain {
	sort.Slice(ac.adapters, func(a, b int) bool {
		return ac.adapters[a].outputJolt > ac.adapters[b].outputJolt
	})
	return ac
}

func (ac *AdapterChain) maxJolt() (max int) {
	for _, a := range ac.adapters {
		if a.outputJolt > max {
			max = a.outputJolt
		}
	}
	return max
}

type Adapter struct {
	outputJolt int
}

func NewAdapter(outputJolt int) Adapter {
	return Adapter{outputJolt}
}

func getAdapters(lines []string) (adapters []Adapter) {
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		adapters = append(adapters, NewAdapter(num))
	}
	return adapters
}
