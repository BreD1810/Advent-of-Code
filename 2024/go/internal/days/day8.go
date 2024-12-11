package days

import (
	"aoc-2024/internal/util"
	"fmt"

	"github.com/mowshon/iterium"
)

func Day8() {
	fileContents := util.ReadFile2DRune("inputs/day8.txt")
	fmt.Printf("Part 1: %d\n", day8Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day8Part2(fileContents))
}

func day8Part1(input [][]rune) int {
	maxY := len(input)
	maxX := len(input[0])

	freqMap := make(map[rune][]util.Coordinate)
	for y, row := range input {
		for x, val := range row {
			if val == '.' {
				continue
			}

			if _, ok := freqMap[val]; !ok {
				freqMap[val] = []util.Coordinate{{X: x, Y: y}}
			} else {
				freqMap[val] = append(freqMap[val], util.Coordinate{X: x, Y: y})
			}
		}
	}

	antinodeMap := make(map[util.Coordinate]struct{})

	for _, positions := range freqMap {
		if len(positions) < 2 {
			continue
		}

		combinations := iterium.Combinations(positions, 2)
		combos, err := combinations.Slice()
		if err != nil {
			panic(err)
		}

		fmt.Printf("combos: %v\n", combos)

		for n, c := range combos {
			fmt.Printf("node: %q\n", n)
			first := c[0]
			second := c[1]
			xDiff := first.X - second.X
			yDiff := first.Y - second.Y
			fmt.Printf("x diff: %d, y diff: %d\n", xDiff, yDiff)

			firstAnti := util.Coordinate{X: first.X + xDiff, Y: first.Y + yDiff}
			if firstAnti.X >= 0 && firstAnti.Y >= 0 && firstAnti.X < maxX && firstAnti.Y < maxY {
				fmt.Printf("adding antinode: %v\n", firstAnti)
				antinodeMap[firstAnti] = struct{}{}
			}

			secondAnti := util.Coordinate{X: second.X - xDiff, Y: second.Y - yDiff}
			if secondAnti.X >= 0 && secondAnti.Y >= 0 && secondAnti.X < maxX && secondAnti.Y < maxY {
				fmt.Printf("adding antinode: %v\n", secondAnti)
				antinodeMap[secondAnti] = struct{}{}
			}
		}
	}

	return len(antinodeMap)
}

func day8Part2(input [][]rune) int {
	maxY := len(input)
	maxX := len(input[0])

	freqMap := make(map[rune][]util.Coordinate)
	for y, row := range input {
		for x, val := range row {
			if val == '.' {
				continue
			}

			coord := util.Coordinate{X: x, Y: y}
			if _, ok := freqMap[val]; !ok {
				freqMap[val] = []util.Coordinate{coord}
			} else {
				freqMap[val] = append(freqMap[val], coord)
			}
		}
	}

	antinodeMap := make(map[util.Coordinate]struct{})
	for _, positions := range freqMap {
		if len(positions) < 2 {
			continue
		}

		for _, pos := range positions {
			antinodeMap[pos] = struct{}{}
		}

		combinations := iterium.Combinations(positions, 2)
		combos, err := combinations.Slice()
		if err != nil {
			panic(err)
		}

		fmt.Printf("combos: %v\n", combos)

		for n, c := range combos {
			fmt.Printf("node: %q\n", n)
			first := c[0]
			second := c[1]
			xDiff := first.X - second.X
			yDiff := first.Y - second.Y
			fmt.Printf("x diff: %d, y diff: %d\n", xDiff, yDiff)

			firstAnti := util.Coordinate{X: first.X + xDiff, Y: first.Y + yDiff}
			for firstAnti.X >= 0 && firstAnti.Y >= 0 && firstAnti.X < maxX && firstAnti.Y < maxY {
				fmt.Printf("adding antinode: %v\n", firstAnti)
				antinodeMap[firstAnti] = struct{}{}
				firstAnti.X += xDiff
				firstAnti.Y += yDiff
			}

			secondAnti := util.Coordinate{X: second.X - xDiff, Y: second.Y - yDiff}
			for secondAnti.X >= 0 && secondAnti.Y >= 0 && secondAnti.X < maxX && secondAnti.Y < maxY {
				fmt.Printf("adding antinode: %v\n", secondAnti)
				antinodeMap[secondAnti] = struct{}{}
				secondAnti.X -= xDiff
				secondAnti.Y -= yDiff
			}
		}
	}

	return len(antinodeMap)
}
