package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
)

func Day10() {
	fileContents := util.ReadFileLines("inputs/day10-actual.txt")
	fmt.Printf("Part 1: %d\n", day10Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day10Part2(fileContents))
}

type pipe int

const (
	vertPipe   pipe = iota // |
	horizPipe              // -
	nePipe                 // L
	nwPipe                 // J
	swPipe                 // 7
	sePipe                 // F
	groundPipe             // .
	startPipe              // S
)

func day10Part1(fileContents []string) int {
	startCoord, pipes := parsePipeMap(fileContents)

	u := findLoopLength(util.Coordinate{X: startCoord.X, Y: startCoord.Y - 1}, pipes, make(map[util.Coordinate]bool), 0)
	d := findLoopLength(util.Coordinate{X: startCoord.X, Y: startCoord.Y + 1}, pipes, make(map[util.Coordinate]bool), 0)
	l := findLoopLength(util.Coordinate{X: startCoord.X - 1, Y: startCoord.Y}, pipes, make(map[util.Coordinate]bool), 0)
	r := findLoopLength(util.Coordinate{X: startCoord.X + 1, Y: startCoord.Y}, pipes, make(map[util.Coordinate]bool), 0)

	return max(u, d, l, r) / 2
}

func day10Part2(fileContents []string) int {
	startCoord, pipes := parsePipeMap(fileContents)

	pipePath := make(map[util.Coordinate]bool)
	length := findLoopLength(util.Coordinate{X: startCoord.X, Y: startCoord.Y - 1}, pipes, pipePath, 0)

	pipePathD := make(map[util.Coordinate]bool)
	if d := findLoopLength(util.Coordinate{X: startCoord.X, Y: startCoord.Y + 1}, pipes, pipePathD, 0); d > length {
		length = d
		pipePath = pipePathD
	}

	pipePathL := make(map[util.Coordinate]bool)
	if l := findLoopLength(util.Coordinate{X: startCoord.X - 1, Y: startCoord.Y}, pipes, pipePathL, 0); l > length {
		length = l
		pipePath = pipePathL
	}

	pipePathR := make(map[util.Coordinate]bool)
	if r := findLoopLength(util.Coordinate{X: startCoord.X + 1, Y: startCoord.Y}, pipes, pipePathR, 0); r > length {
		length = r
		pipePath = pipePathR
	}

	pipes[startCoord.Y][startCoord.X] = vertPipe

	count := 0
	for y, pipeRow := range pipes {
		inside := false
		lastStart := horizPipe
		for x, p := range pipeRow {
			if !pipePath[util.Coordinate{X: x, Y: y}] {
				if inside {
					count++
				}
				continue
			}
			if p == horizPipe {
				continue
			}
			if p == vertPipe || p == sePipe || p == nePipe {
				inside = !inside
				lastStart = p
			}
			if p == nwPipe && lastStart == nePipe {
				inside = !inside
				lastStart = horizPipe
			}
			if p == swPipe && lastStart == sePipe {
				inside = !inside
				lastStart = horizPipe
			}
		}
	}

	return count
}

func parsePipeMap(inp []string) (util.Coordinate, [][]pipe) {
	pipes := make([][]pipe, len(inp))
	var startCoord util.Coordinate

	for y, l := range inp {
		pipeRow := make([]pipe, len(l))
		for x, c := range l {
			p := runeToPipe(c)
			pipeRow[x] = p
			if p == startPipe {
				startCoord = util.Coordinate{X: x, Y: y}
			}
		}
		pipes[y] = pipeRow
	}

	return startCoord, pipes
}

func runeToPipe(r rune) pipe {
	switch r {
	case '|':
		return vertPipe
	case '-':
		return horizPipe
	case 'L':
		return nePipe
	case 'J':
		return nwPipe
	case '7':
		return swPipe
	case 'F':
		return sePipe
	case '.':
		return groundPipe
	case 'S':
		return startPipe
	}
	log.Fatalf("could not parse part of pipe map: %c\n", r)
	return groundPipe
}

func findLoopLength(p util.Coordinate, pipes [][]pipe, visited map[util.Coordinate]bool, stepCount int) int {
	if visited[p] || p.X < 0 || p.Y < 0 || p.X >= len(pipes[0]) || p.Y >= len(pipes) || pipes[p.Y][p.X] == groundPipe {
		return stepCount
	}
	stepCount++
	visited[p] = true

	switch pipes[p.Y][p.X] {
	case vertPipe:
		u := findLoopLength(util.Coordinate{X: p.X, Y: p.Y - 1}, pipes, visited, stepCount)
		d := findLoopLength(util.Coordinate{X: p.X, Y: p.Y + 1}, pipes, visited, stepCount)
		return max(u, d)
	case horizPipe:
		l := findLoopLength(util.Coordinate{X: p.X - 1, Y: p.Y}, pipes, visited, stepCount)
		r := findLoopLength(util.Coordinate{X: p.X + 1, Y: p.Y}, pipes, visited, stepCount)
		return max(l, r)
	case nePipe:
		r := findLoopLength(util.Coordinate{X: p.X + 1, Y: p.Y}, pipes, visited, stepCount)
		u := findLoopLength(util.Coordinate{X: p.X, Y: p.Y - 1}, pipes, visited, stepCount)
		return max(r, u)
	case nwPipe:
		u := findLoopLength(util.Coordinate{X: p.X, Y: p.Y - 1}, pipes, visited, stepCount)
		l := findLoopLength(util.Coordinate{X: p.X - 1, Y: p.Y}, pipes, visited, stepCount)
		return max(u, l)
	case swPipe:
		l := findLoopLength(util.Coordinate{X: p.X - 1, Y: p.Y}, pipes, visited, stepCount)
		d := findLoopLength(util.Coordinate{X: p.X, Y: p.Y + 1}, pipes, visited, stepCount)
		return max(l, d)
	case sePipe:
		r := findLoopLength(util.Coordinate{X: p.X + 1, Y: p.Y}, pipes, visited, stepCount)
		d := findLoopLength(util.Coordinate{X: p.X, Y: p.Y + 1}, pipes, visited, stepCount)
		return max(r, d)
	case startPipe:
		return stepCount
	}
	log.Fatalf("unrecognised pipe: %v\n", pipes[p.Y][p.X])
	return 0
}
