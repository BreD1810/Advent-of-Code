package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day10() {
	fileContents := util.ReadFileIntLines("inputs/day10.txt")
	fmt.Printf("Part 1: %d\n", day10Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day10Part2(fileContents))
}

func day10Part1(input [][]int) int {
	trailheads := make(map[util.Coordinate]struct{})
	trailheadEnds := make(map[util.Coordinate]map[util.Coordinate]struct{})

	for y, row := range input {
		for x, val := range row {
			if val == 0 {
				trailheads[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
		}
	}

	for start := range trailheads {
		ends := getTrailheadEnds(input, start, 0)
		endMap := make(map[util.Coordinate]struct{})
		for _, end := range ends {
			endMap[end] = struct{}{}
		}
		trailheadEnds[start] = endMap
	}

	var total int
	for _, endMap := range trailheadEnds {
		total += len(endMap)
	}

	return total
}

func getTrailheadEnds(input [][]int, loc util.Coordinate, expectedVal int) []util.Coordinate {
	newLocs := []util.Coordinate{}
	if loc.X < 0 || loc.Y < 0 || loc.X >= len(input[0]) || loc.Y >= len(input) || input[loc.Y][loc.X] != expectedVal {
		return newLocs
	}
	if input[loc.Y][loc.X] == 9 {
		return []util.Coordinate{loc}
	}

	// Up
	newLocs = append(newLocs, getTrailheadEnds(input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, expectedVal+1)...)
	//Down
	newLocs = append(newLocs, getTrailheadEnds(input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, expectedVal+1)...)
	// Left
	newLocs = append(newLocs, getTrailheadEnds(input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, expectedVal+1)...)
	// Right
	newLocs = append(newLocs, getTrailheadEnds(input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, expectedVal+1)...)

	return newLocs
}

func day10Part2(input [][]int) int {
	trailheads := make(map[util.Coordinate]struct{})

	for y, row := range input {
		for x, val := range row {
			if val == 0 {
				fmt.Printf("adding trailhead %v\n", util.Coordinate{X: x, Y: y})
				trailheads[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
		}
	}

	var total int
	for start := range trailheads {
		fmt.Printf("starting at: %v\n", start)
		score := getTrailPathCount(input, start, 0)
		fmt.Printf("has score: %v\n", score)
		total += score
	}

	return total
}

func getTrailPathCount(input [][]int, loc util.Coordinate, expectedVal int) int {
	var total int
	if loc.X < 0 || loc.Y < 0 || loc.X >= len(input[0]) || loc.Y >= len(input) || input[loc.Y][loc.X] != expectedVal {
		return total
	}
	if input[loc.Y][loc.X] == 9 {
		return 1
	}

	// Up
	total += getTrailPathCount(input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, expectedVal+1)
	//Down
	total += getTrailPathCount(input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, expectedVal+1)
	// Left
	total += getTrailPathCount(input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, expectedVal+1)
	// Right
	total += getTrailPathCount(input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, expectedVal+1)

	return total
}
