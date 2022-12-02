package days

import (
	"aoc-2022/util"
	"fmt"
	"log"
	"strings"
)

type move int

const (
	rock move = iota + 1
	paper
	scissors
)

const (
	lose int = 0
	draw int = 3
	win  int = 6
)

func Day2() {
	fileContents := util.ReadFileLines("inputs/day02-1.txt")
	fmt.Printf("Part 1: %d\n", day2Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day2Part2(fileContents))
}

func day2Part1(fileContents []string) int {
	opMs, myMs := parseInput(fileContents)
	var total int

	for i, opM := range opMs {
		total += getScore(opM, myMs[i])
	}

	return total
}

func parseInput(ls []string) ([]move, []move) {
	op, me := make([]move, len(ls)), make([]move, len(ls))

	for i, l := range ls {
		s := strings.Split(l, " ")
		op[i] = parseOpponentMove(s[0])
		me[i] = parseMyMove(s[1])
	}

	return op, me
}

func parseOpponentMove(s string) move {
	switch s {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	}
	log.Fatalf("Unable to parse opponent move %s", s)
	return rock
}

func getScore(opM move, myM move) int {
	shapescore := int(myM)
	// draw
	if opM == myM {
		return 3 + shapescore
	}

	// my win
	if (opM == rock && myM == paper) ||
		(opM == paper && myM == scissors) ||
		(opM == scissors && myM == rock) {
		return 6 + shapescore
	}

	// lose
	return 0 + shapescore
}

func parseMyMove(s string) move {
	switch s {
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	}
	log.Fatalf("Unable to parse my move %s", s)
	return rock
}

func day2Part2(fileContents []string) int {
	opMs, ends := parseInput2(fileContents)
	var total int

	for i, opM := range opMs {
		total += getScore2(opM, ends[i])
	}

	return total
}

func parseInput2(ls []string) ([]move, []int) {
	op, ends := make([]move, len(ls)), make([]int, len(ls))

	for i, l := range ls {
		s := strings.Split(l, " ")
		op[i] = parseOpponentMove(s[0])
		ends[i] = parseEnd(s[1])
	}

	return op, ends
}

func parseEnd(s string) int {
	switch s {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	log.Fatalf("Unable to parse ending %s", s)
	return 0
}

func getScore2(opM move, end int) int {
	// draw
	if end == 3 {
		return 3 + int(opM)
	}

	// my win
	if end == 6 {
		switch opM {
		case rock:
			return end + int(paper)
		case paper:
			return end + int(scissors)
		case scissors:
			return end + int(rock)
		}
	}

	// lose
	switch opM {
	case rock:
		return int(scissors)
	case paper:
		return int(rock)
	case scissors:
		return int(paper)
	}
	log.Fatal("Couldn't work out move")
	return 0
}
