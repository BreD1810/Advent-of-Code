package days

import (
	"aoc-2021/util"
	"fmt"
	"strings"
)

type direction string

const (
	Forward = "forward"
	Up      = "up"
	Down    = "down"
)

type movement struct {
	dir direction
	val int
}

type position struct {
	horizontal int
	depth      int
	aim        int
}

func newPosition() position {
	return position{horizontal: 0, depth: 0, aim: 0}
}

func Day2() {
	file_contents := util.ReadFileLines("inputs/day02-1.txt")
	movements := getMovementList(file_contents)
	fmt.Printf("Part 1: %d\n", part1(movements))
	fmt.Printf("Part 2: %d\n", part2(movements))
}

func getMovementList(lines []string) []movement {
	lines_len := len(lines)
	movements := make([]movement, lines_len)

	for i := 0; i < lines_len; i++ {
		movements[i] = parseMovement(lines[i])
	}

	return movements
}

func parseMovement(s string) movement {
	splits := strings.Split(s, " ")
	dir := splits[0]
	val := util.GetIntFromString(splits[1])
	return movement{dir: direction(dir), val: val}
}

func part1(movements []movement) int {
	pos := newPosition()

	for i := 0; i < len(movements); i++ {
		movement := movements[i]
		switch movement.dir {
		case Forward:
			pos.horizontal += movement.val
		case Down:
			pos.depth += movement.val
		case Up:
			pos.depth -= movement.val
		}
	}

	return pos.depth * pos.horizontal
}

func part2(movements []movement) int {
	pos := newPosition()

	for i := 0; i < len(movements); i++ {
		movement := movements[i]
		switch movement.dir {
		case Forward:
			pos.horizontal += movement.val
			pos.depth += pos.aim * movement.val
		case Down:
			pos.aim += movement.val
		case Up:
			pos.aim -= movement.val
		}
	}

	return pos.depth * pos.horizontal
}
