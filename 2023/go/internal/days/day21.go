package days

import (
	"aoc-2023/internal/util"
	"fmt"
)

func Day21() {
	fileContents := util.ReadFileLines("inputs/day21-example.txt")
	fmt.Printf("Part 1: %d\n", day21Part1(fileContents, 64))
	fmt.Printf("Part 2: %d\n", day21Part2(fileContents, 26501365))
}

func day21Part1(fileContents []string, stepCount int) int {
	garden, start := parseGardenPlots(fileContents)
	return getGardenPlotCounts(garden, start, stepCount)
}

func day21Part2(fileContents []string, stepCount int) int {
	garden, start := parseGardenPlots(fileContents)
	return getGardenPlotCounts2(garden, start, stepCount)
}

func parseGardenPlots(inp []string) ([][]rune, util.Coordinate) {
	var start util.Coordinate
	garden := make([][]rune, len(inp))

	for y, l := range inp {
		garden[y] = make([]rune, len(l))
		for x, c := range l {
			if c == 'S' {
				start = util.Coordinate{X: x, Y: y}
				c = '.'
			}
			garden[y][x] = c
		}
	}

	return garden, start
}

func getGardenPlotCounts(garden [][]rune, startPoint util.Coordinate, stepCount int) int {
	q := []util.Coordinate{startPoint}
	moveDirections := []util.Coordinate{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
	stepsTaken := 0
	var visited map[util.Coordinate]struct{}

	for stepsTaken < stepCount {
		var newQ []util.Coordinate
		visited = make(map[util.Coordinate]struct{})

		for len(q) > 0 {
			curPos := q[0]
			q = q[1:]
			for _, d := range moveDirections {
				newPos := util.Coordinate{X: curPos.X + d.X, Y: curPos.Y + d.Y}
				if newPos.X < 0 || newPos.X >= len(garden[0]) || newPos.Y < 0 || newPos.Y >= len(garden) ||
					garden[newPos.Y][newPos.X] == '#' {
					continue
				}
				if _, ok := visited[newPos]; !ok {
					visited[newPos] = struct{}{}
					newQ = append(newQ, newPos)
				}
			}
		}
		q = newQ
		stepsTaken++
	}

	return len(visited)
}

func getGardenPlotCounts2(garden [][]rune, startPoint util.Coordinate, stepCount int) int {
	q := []util.Coordinate{startPoint}
	moveDirections := []util.Coordinate{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}
	stepsTaken := 0
	var visited map[util.Coordinate]struct{}
	var polyVals []int

	for stepsTaken < stepCount {
		var newQ []util.Coordinate
		visited = make(map[util.Coordinate]struct{})

		for len(q) > 0 {
			curPos := q[0]
			q = q[1:]
			for _, d := range moveDirections {
				newPos := util.Coordinate{X: curPos.X + d.X, Y: curPos.Y + d.Y}
				adjustedNewPos := util.Coordinate{X: ((newPos.X % len(garden[0])) + len(garden[0])) % len(garden[0]), Y: ((newPos.Y % len(garden)) + len(garden)) % len(garden)}
				if garden[adjustedNewPos.Y][adjustedNewPos.X] == '#' {
					continue
				}
				if _, ok := visited[newPos]; !ok {
					visited[newPos] = struct{}{}
					newQ = append(newQ, newPos)
				}
			}
		}
		q = newQ
		stepsTaken++
		if stepsTaken%len(garden) == stepCount%len(garden) {
			polyVals = append(polyVals, len(visited))
			if len(polyVals) == 3 {
				firstTerm := polyVals[0]
				secondTerm := (polyVals[1] - polyVals[0]) * (stepCount / len(garden))
				thirdTermConst := (polyVals[2] - polyVals[1]) - (polyVals[1] - polyVals[0])
				thirdTerm := thirdTermConst * ((stepCount / len(garden)) * ((stepCount / len(garden)) - 1) / 2)
				return firstTerm + secondTerm + thirdTerm
			}
		}
	}

	return len(visited)
}
