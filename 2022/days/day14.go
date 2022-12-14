package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"
)

func Day14() {
	fileContents := util.ReadFileLines("inputs/day14-1.txt")
	fmt.Printf("Part 1: %d\n", day14Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day14Part2(fileContents))
}

func day14Part1(fileContents []string) int {
	rockPathCoords, maxX, maxY := parseRockPaths(fileContents)

	layout := makeRockLayout(rockPathCoords, maxX, maxY)
	// util.Print2DRuneArray(layout)
	return simulateSand(layout)
}

func parseRockPaths(fileContents []string) ([][]util.Coordinate, int, int) {
	rockPaths := make([][]util.Coordinate, len(fileContents))
	maxX, maxY := 0, 0

	for i, l := range fileContents {
		var curPathCoords []util.Coordinate
		coordStrings := strings.Split(l, " -> ")
		for _, s := range coordStrings {
			sSplit := strings.Split(s, ",")
			x := util.GetIntFromString(sSplit[0])
			if x > maxX {
				maxX = x
			}
			y := util.GetIntFromString(sSplit[1])
			if y > maxY {
				maxY = y
			}
			curPathCoords = append(curPathCoords, util.Coordinate{X: x, Y: y})
		}
		rockPaths[i] = curPathCoords
	}

	return rockPaths, maxX + 1, maxY + 1
}

func makeRockLayout(rockPathCoords [][]util.Coordinate, maxX, maxY int) [][]rune {
	layout := make([][]rune, maxY)
	for y := 0; y < maxY; y++ {
		row := make([]rune, maxX)
		for x := 0; x < maxX; x++ {
			row[x] = '.'
		}
		layout[y] = row
	}

	for _, coordPath := range rockPathCoords {
		firstCoord := coordPath[1]
		var secondCoord util.Coordinate

		layout[firstCoord.Y][firstCoord.X] = '#'
		for n := 1; n < len(coordPath); n++ {
			firstCoord = coordPath[n-1]
			secondCoord = coordPath[n]

			if firstCoord.X == secondCoord.X {
				smallestY, largestY := firstCoord.Y, secondCoord.Y
				if firstCoord.Y > secondCoord.Y {
					smallestY, largestY = secondCoord.Y, firstCoord.Y
				}
				for y := smallestY; y <= largestY; y++ {
					layout[y][firstCoord.X] = '#'
				}
			} else {
				smallestX, largestX := firstCoord.X, secondCoord.X
				if firstCoord.X > secondCoord.X {
					smallestX, largestX = secondCoord.X, firstCoord.X
				}
				for x := smallestX; x <= largestX; x++ {
					layout[firstCoord.Y][x] = '#'
				}
			}
		}
	}

	return layout
}

func simulateSand(layout [][]rune) int {
	var sandCount int

	for {
		sandPos := util.Coordinate{X: 500, Y: 0}
		stopped := false
		if layout[0][500] == '#' {
			return sandCount
		}
		for !stopped {
			if sandPos.Y+1 < len(layout) && layout[sandPos.Y+1][sandPos.X] == '.' {
				sandPos.Y++
			} else if sandPos.Y+1 < len(layout) && sandPos.X-1 >= 0 && layout[sandPos.Y+1][sandPos.X-1] == '.' {
				sandPos.Y++
				sandPos.X--
			} else if sandPos.Y+1 < len(layout) && sandPos.X+1 < len(layout[sandPos.Y+1]) && layout[sandPos.Y+1][sandPos.X+1] == '.' {
				sandPos.Y++
				sandPos.X++
			} else if sandPos.Y+1 == len(layout) {
				return sandCount
			} else {
				layout[sandPos.Y][sandPos.X] = '#'
				stopped = true
			}
		}
		sandCount++
	}
}

func day14Part2(fileContents []string) int {
	rockPathCoords, maxX, maxY := parseRockPaths(fileContents)

	layout := makeRockLayout(rockPathCoords, maxX+1000, maxY+2)
	addFloor(layout)
	// util.Print2DRuneArray(layout)
	return simulateSand(layout)
}

func addFloor(layout [][]rune) {
	floorLevel := len(layout) - 1
	floor := layout[floorLevel]
	for i := range floor {
		layout[floorLevel][i] = '#'
	}
}
