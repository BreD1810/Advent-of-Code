package days

import (
	"aoc-2022/util"
	"fmt"
)

func Day1() {
	fileContents := util.ReadFileLines("inputs/day01-1.txt")
	fmt.Printf("Part 1: %d\n", day1Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day1Part2(fileContents))
}

func day1Part1(fileContents []string) int {
	var curTotal, maxTotal int

	for _, line := range fileContents {
		if line == "" {
			if maxTotal < curTotal {
				maxTotal = curTotal
			}
			curTotal = 0
		} else {
			cals := util.GetIntFromString(line)
			curTotal += cals
		}
	}
	if curTotal != 0 && maxTotal < curTotal {
		maxTotal = curTotal
	}

	return maxTotal
}

func day1Part2(fileContents []string) int {
	totals := make([]int, 0)
	var curTotal int

	for _, line := range fileContents {
		if line == "" {
			totals = append(totals, curTotal)
			curTotal = 0
		} else {
			cals := util.GetIntFromString(line)
			curTotal += cals
		}
	}
	if curTotal != 0 {
		totals = append(totals, curTotal)
	}

	util.SortIntSliceDecending(totals)

	return totals[0] + totals[1] + totals[2]
}
