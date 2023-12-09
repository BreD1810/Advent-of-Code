package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"strings"
)

func Day9() {
	fileContents := util.ReadFileLines("inputs/day9-actual.txt")
	fmt.Printf("Part 1: %d\n", day9Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day9Part2(fileContents))
}

func day9Part1(fileContents []string) int {
	inputs := getOasisInputs(fileContents)
	total := 0
	for _, curInput := range inputs {
		total += extrapolateNewOasisValue(curInput)
	}
	return total
}

func day9Part2(fileContents []string) int {
	inputs := getOasisInputs(fileContents)
	total := 0
	for _, curInput := range inputs {
		total += extrapolateNewOasisValueBackwards(curInput)
	}
	return total
}

func getOasisInputs(lines []string) [][]int {
	inputs := make([][]int, len(lines))
	for i, l := range lines {
		ns := strings.Fields(l)
		curInput := make([]int, len(ns))
		for j, n := range ns {
			curInput[j] = util.GetIntFromString(n)
		}
		inputs[i] = curInput
	}
	return inputs
}

func extrapolateNewOasisValue(vals []int) int {
	newVal := vals[len(vals)-1]
	lastDiffs := vals

	for l := len(vals) - 1; l > 0; l-- {
		diffs := make([]int, l)
		for i := 0; i < len(lastDiffs)-1; i++ {
			diffs[i] = lastDiffs[i+1] - lastDiffs[i]
		}

		if isAllZeros(diffs) {
			return newVal
		}

		lastDiffs = diffs
		newVal += lastDiffs[len(lastDiffs)-1]
	}

	return newVal
}

func isAllZeros(vals []int) bool {
	for _, v := range vals {
		if v != 0 {
			return false
		}
	}
	return true
}

func extrapolateNewOasisValueBackwards(vals []int) int {
	zerothVals := []int{vals[0]}
	lastDiffs := vals

	for l := len(vals) - 1; l > 0; l-- {
		diffs := make([]int, l)
		for i := 0; i < len(lastDiffs)-1; i++ {
			diffs[i] = lastDiffs[i+1] - lastDiffs[i]
		}

		if isAllZeros(diffs) {
			break
		}

		zerothVals = append(zerothVals, diffs[0])
		lastDiffs = diffs
	}

	newVal := 0
	for i := len(zerothVals) - 1; i >= 0; i-- {
		newVal = zerothVals[i] - newVal
	}
	return newVal
}
