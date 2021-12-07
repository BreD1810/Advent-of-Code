package days

import (
	"aoc-2021/util"
	"fmt"
	"math"
	"strings"
)

func Day7() {
	fileContents := util.ReadFileLines("inputs/day07-1.txt")
	// fileContents := util.ReadFileLines("inputs/test.txt")
	fmt.Printf("Part 1: %d\n", day7Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day7Part2(fileContents))
}

func day7Part1(lines []string) int {
	input := lines[0]
	crabPosStrings := strings.Split(input, ",")

	maxPos := 0
	crabPos := make([]int, len(crabPosStrings))
	for i, crab := range crabPosStrings {
		crabPos[i] = util.GetIntFromString(crab)
		if crabPos[i] > maxPos {
			maxPos = crabPos[i]
		}
	}

	minFuel := math.MaxInt64
	for pos := 0; pos <= maxPos; pos++ {
		total := 0
		for _, crab := range crabPos {
			total += util.IntAbs(pos - crab)
		}

		if total < minFuel {
			minFuel = total
		}
	}

	return minFuel
}

func day7Part2(lines []string) int {
	input := lines[0]
	crabPosStrings := strings.Split(input, ",")

	maxPos := 0
	crabPos := make([]int, len(crabPosStrings))
	for i, crab := range crabPosStrings {
		crabPos[i] = util.GetIntFromString(crab)
		if crabPos[i] > maxPos {
			maxPos = crabPos[i]
		}
	}

	minFuel := math.MaxInt64
	for pos := 0; pos <= maxPos; pos++ {
		total := 0
		for _, crab := range crabPos {
			total += calcFuelCost(util.IntAbs(pos - crab))
		}

		if total < minFuel {
			minFuel = total
		}
	}

	return minFuel
}

func calcFuelCost(dx int) int {
	return (dx * (dx + 1)) / 2
}
