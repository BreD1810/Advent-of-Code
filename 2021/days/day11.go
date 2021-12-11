package days

import (
	"aoc-2021/util"
	"fmt"
)

const (
	gridX      = 10
	gridY      = 10
	part1Steps = 100
)

func Day11() {
	fileContents := util.ReadFileLines("inputs/day11-1.txt")
	fmt.Printf("Part 1: %d\n", day11Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day11Part2(fileContents))
}

func day11Part1(lines []string) int {
	octos := make([][]int, len(lines))
	for i, line := range lines {
		octos[i] = make([]int, len(line))
		for j, digit := range line {
			octos[i][j] = util.GetIntFromString(string(digit))
		}
	}

	totalFlashes := 0
	for i := 0; i < part1Steps; i++ {
		// fmt.Println("Step ", i+1)
		flashes := 0
		flashes, octos = runStep(octos)
		// printOctos(octos)
		totalFlashes += flashes
	}

	return totalFlashes
}

func day11Part2(lines []string) int {
	octos := make([][]int, len(lines))
	for i, line := range lines {
		octos[i] = make([]int, len(line))
		for j, digit := range line {
			octos[i][j] = util.GetIntFromString(string(digit))
		}
	}

	count := 1
	for {
		// fmt.Println("Step ", i+1)
		_, octos = runStep(octos)
		// printOctos(octos)
		if checkSync(octos) {
			return count
		}
		count++
	}
}

func checkSync(octos [][]int) bool {
	for _, octoLine := range octos {
		for _, octo := range octoLine {
			if octo != 0 {
				return false
			}
		}
	}

	return true
}

func runStep(octos [][]int) (int, [][]int) {
	// Increase by 1
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[i]); j++ {
			octos[i][j]++
		}
	}

	// Increase all neighbours by 1 if 9
	flashes := 0
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos[i]); j++ {
			flashes += flashOcto(octos, i, j)
		}
	}

	return flashes, octos
}

func flashOcto(octos [][]int, i, j int) int {
	noFlashes := 0
	if octos[i][j] > 9 {
		octos[i][j] = 0
		noFlashes++

		flashNeighbour := func(y, x int) {
			if y >= 0 && x >= 0 && y < len(octos) && x < len(octos[y]) && octos[y][x] != 0 {
				octos[y][x]++
				noFlashes += flashOcto(octos, y, x)
			}
		}

		flashNeighbour(i-1, j)
		flashNeighbour(i-1, j+1)
		flashNeighbour(i, j+1)
		flashNeighbour(i+1, j+1)
		flashNeighbour(i+1, j)
		flashNeighbour(i+1, j-1)
		flashNeighbour(i, j-1)
		flashNeighbour(i-1, j-1)
	}

	return noFlashes
}

func printOctos(octos [][]int) {
	for _, octoLine := range octos {
		for _, octo := range octoLine {
			fmt.Printf("%d", octo)
		}
		fmt.Println()
	}
	fmt.Println()
}
