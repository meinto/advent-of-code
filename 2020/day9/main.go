package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	nums := convertInt(lines)
	invalidIndex := -1
	for i, num := range nums[25:] {
		isValid := validate(nums[i:25+i], num)
		if !isValid {
			invalidIndex = i + 25
			fmt.Println(num, "is not valid")
			break
		}
	}

	continuousRange := findContinuosRange(nums, invalidIndex)
	fmt.Println(strings.Join(convertString(continuousRange), " + "), "=", nums[invalidIndex])
	max, min := maxMin(continuousRange)
	fmt.Println("max of continuous range:", max)
	fmt.Println("min of continuous range:", min)
	fmt.Println("sum of max + min:", max+min)
}

func validate(nums []int, search int) bool {
	for i, first := range nums {
		for _, second := range nums[i+1:] {
			if first+second == search {
				return true
			}
		}
	}
	return false
}

func findContinuosRange(nums []int, invalidIndex int) (continuousRange []int) {
	invalidNum := nums[invalidIndex]
	for i, first := range nums {
		continuousRange = []int{first}
		for j, n := range nums[i+1:] {
			if j >= invalidIndex {
				break
			}
			continuousRange = append(continuousRange, n)
			sum := sum(continuousRange)
			if sum > invalidNum {
				break
			} else if sum == invalidNum {
				return continuousRange
			}
		}
	}
	return continuousRange
}

func sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

func convertInt(lines []string) (nums []int) {
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		nums = append(nums, num)
	}
	return nums
}

func convertString(nums []int) (lines []string) {
	for _, n := range nums {
		line := strconv.Itoa(n)
		lines = append(lines, line)
	}
	return lines
}

func maxMin(nums []int) (max, min int) {
	min = -1
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min || min < 0 {
			min = n
		}
	}
	return max, min
}
