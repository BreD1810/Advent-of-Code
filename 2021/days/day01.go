package days

import (
	"aoc-2021/util"
	"fmt"
)

func Run() {
	file_contents := util.ReadFileLines("inputs/day01-1.txt")
	fmt.Printf("Part 1: %d\n", countIncreases(file_contents))
	fmt.Printf("Part 2: %d\n", countWindowIncreases(file_contents))
}

func countIncreases(values []string) int {
	counter := 0
	prev_val := util.GetIntFromString(values[0])

	for i := 1; i < len(values); i++ {
		cur_val := util.GetIntFromString(values[i])
		if cur_val > prev_val {
			counter++
		}
		prev_val = cur_val
	}
	return counter
}

func countWindowIncreases(values []string) int {
	counter := 0
	prev_sum := sumArray(values[0:3])
	for i := 1; i < len(values)-2; i++ {
		cur_sum := sumArray(values[i : i+3])
		if cur_sum > prev_sum {
			counter++
		}
		prev_sum = cur_sum
	}
	return counter
}

func sumArray(values []string) int {
	sum := 0
	for _, val := range values {
		sum += util.GetIntFromString(val)
	}
	return sum
}
