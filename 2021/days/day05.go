package days

import (
	"aoc-2021/util"
	"fmt"
	"strings"
)

type line struct {
	start coordinate
	end   coordinate
}

type coordinate struct {
	x int
	y int
}

func Day5() {
	fileContents := util.ReadFileLines("inputs/day05-1.txt")
	// fileContents := util.ReadFileLines("inputs/test.txt")
	fmt.Printf("Part 1: %d\n", day5Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day5Part2(fileContents))
}

func day5Part1(entries []string) int {
	lines := make([]line, 0)

	for _, entry := range entries {
		coordinates := strings.Split(entry, " -> ")

		line := line{
			start: coordinateFromString(coordinates[0]),
			end:   coordinateFromString(coordinates[1]),
		}

		if line.start.x == line.end.x || line.start.y == line.end.y {
			lines = append(lines, line)
		}
	}

	points := make(map[coordinate]int)
	for _, line := range lines {
		if line.start.x == line.end.x {
			if line.start.y > line.end.y {
				line.start.y, line.end.y = line.end.y, line.start.y
			}
			for y := line.start.y; y <= line.end.y; y++ {
				point := coordinate{x: line.start.x, y: y}
				points[point]++
			}
		} else {
			if line.start.x > line.end.x {
				line.start.x, line.end.x = line.end.x, line.start.x
			}
			for x := line.start.x; x <= line.end.x; x++ {
				point := coordinate{x: x, y: line.start.y}
				points[point]++
			}
		}
	}

	return countOverlaps(points)
}

func day5Part2(entries []string) int {
	lines := make([]line, 0)

	for _, entry := range entries {
		coordinates := strings.Split(entry, " -> ")

		line := line{
			start: coordinateFromString(coordinates[0]),
			end:   coordinateFromString(coordinates[1]),
		}

		lines = append(lines, line)
	}

	points := make(map[coordinate]int)
	for _, line := range lines {
		start, end := line.start, line.end
		// Swap required for both horizontal and diagonal
		if start.y > end.y {
			start, end = end, start
		}

		if start.x == end.x {
			for y := start.y; y <= end.y; y++ {
				point := coordinate{x: start.x, y: y}
				points[point]++
			}
		} else if start.y == end.y {
			if start.x > end.x {
				start.x, end.x = end.x, start.x
			}
			for x := start.x; x <= end.x; x++ {
				point := coordinate{x: x, y: start.y}
				points[point]++
			}
		} else {
			if start.x > end.x { // Slopes up to the left
				for x, y := start.x, start.y; y <= end.y; x, y = x-1, y+1 {
					point := coordinate{x, y}
					points[point]++
				}
			} else { // Slopes up to the right
				for x, y := start.x, start.y; y <= end.y; x, y = x+1, y+1 {
					point := coordinate{x, y}
					points[point]++
				}
			}
		}
	}

	return countOverlaps(points)
}

func coordinateFromString(s string) coordinate {
	parts := strings.Split(s, ",")
	x := util.GetIntFromString(parts[0])
	y := util.GetIntFromString(parts[1])
	return coordinate{x, y}
}

func countOverlaps(m map[coordinate]int) int {
	overlaps := 0

	for _, count := range m {
		if count > 1 {
			overlaps++
		}
	}

	return overlaps
}
