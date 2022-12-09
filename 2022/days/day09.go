package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"
)

type direction int

const (
	Up direction = iota
	Down
	Left
	Right
)

type ropeMove struct {
	dir direction
	v   int
}

type ropePos struct {
	x int
	y int
}

func Day9() {
	fileContents := util.ReadFileLines("inputs/day09-1.txt")
	fmt.Printf("Part 1: %d\n", day9Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day9Part2(fileContents))
}

func day9Part1(fileContents []string) int {
	moves := getRopeMoves(fileContents)
	head, tail := ropePos{0, 0}, ropePos{0, 0}
	tailPosMap := map[ropePos]struct{}{tail: {}}

	for _, m := range moves {
		v := m.v
		for v > 0 {
			switch m.dir {
			case Up:
				head.y++
			case Down:
				head.y--
			case Right:
				head.x++
			case Left:
				head.x--
			}
			v--
			tail = adjustTail(tail, head)
			tailPosMap[tail] = struct{}{}
		}
	}

	return len(tailPosMap)
}

func getRopeMoves(fileContents []string) []ropeMove {
	moves := make([]ropeMove, len(fileContents))

	for i, l := range fileContents {
		splits := strings.Split(l, " ")
		var d direction

		switch splits[0] {
		case "U":
			d = Up
		case "D":
			d = Down
		case "R":
			d = Right
		case "L":
			d = Left
		}

		amount := util.GetIntFromString(splits[1])

		moves[i] = ropeMove{dir: d, v: amount}
	}

	return moves
}

func adjustTail(tail, head ropePos) (newTail ropePos) {
	newTail = tail
	ropeDiff := ropePos{head.x - tail.x, head.y - tail.y}

	// Change in x
	switch ropeDiff {
	// case ropePos{0, 0}, ropePos{-1, 0}, ropePos{1, 0}, ropePos{0, -1}, ropePos{0, 1}:
	// 	return
	case ropePos{2, 0}, ropePos{1, 2}, ropePos{1, -2}, ropePos{2, 1}, ropePos{2, -1}, ropePos{2, 2}, ropePos{2, -2}:
		newTail.x++
	case ropePos{-2, 0}, ropePos{-1, 2}, ropePos{-1, -2}, ropePos{-2, 1}, ropePos{-2, -1}, ropePos{-2, -2}, ropePos{-2, 2}:
		newTail.x--
	}

	// Change in y
	switch ropeDiff {
	// case ropePos{0, 2}, ropePos{1, 2}, ropePos{-1, 2}, ropePos{2, 1}, ropePos{-2, 1}:
	case ropePos{0, 2}, ropePos{1, 2}, ropePos{-1, 2}, ropePos{-2, 1}, ropePos{2, 1}, ropePos{2, 2}, ropePos{-2, 2}:
		newTail.y++
	case ropePos{0, -2}, ropePos{1, -2}, ropePos{-1, -2}, ropePos{-2, -2}, ropePos{-2, -1}, ropePos{2, -1}, ropePos{2, -2}:
		newTail.y--
	}

	return
}

func day9Part2(fileContents []string) int {
	moves := getRopeMoves(fileContents)
	knots := make([]ropePos, 10)
	tailPosMap := map[ropePos]struct{}{knots[9]: {}}

	for _, m := range moves {
		v := m.v
		for v > 0 {
			switch m.dir {
			case Up:
				knots[0].y++
			case Down:
				knots[0].y--
			case Right:
				knots[0].x++
			case Left:
				knots[0].x--
			}
			v--
			for i := 0; i < 9; i++ {
				knots[i+1] = adjustTail(knots[i+1], knots[i])
			}
			tailPosMap[knots[9]] = struct{}{}
		}
	}

	return len(tailPosMap)
}
