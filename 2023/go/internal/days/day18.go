package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type diggerMove struct {
	direction movement
	distance  int64
}

func Day18() {
	fileContents := util.ReadFileLines("inputs/day18-example.txt")
	fmt.Printf("Part 1: %d\n", day18Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day18Part2(fileContents))
}

func day18Part1(fileContents []string) int64 {
	moves := parseDiggerMoves(fileContents)
	return getLagoonSize(moves)
}

func day18Part2(fileContents []string) int64 {
	moves := parseDiggerMovesFromColour(fileContents)
	return getLagoonSize(moves)
}

func parseDiggerMoves(inp []string) []diggerMove {
	moves := make([]diggerMove, len(inp))
	for i, l := range inp {
		ss := strings.Split(l, " ")
		dir := udlrToMovement(ss[0])
		dist := int64(util.GetIntFromString(ss[1]))
		moves[i] = diggerMove{direction: dir, distance: dist}
	}
	return moves
}

func parseDiggerMovesFromColour(inp []string) []diggerMove {
	moves := make([]diggerMove, len(inp))
	for i, l := range inp {
		ss := strings.Split(l, " ")
		colour := ss[2]
		dir := hexToMovement(string(colour[7]))
		dist, _ := strconv.ParseInt(colour[2:7], 16, 64)
		moves[i] = diggerMove{direction: dir, distance: dist}
	}
	return moves
}

func udlrToMovement(s string) movement {
	switch s {
	case "U":
		return moveUp
	case "D":
		return moveDown
	case "L":
		return moveLeft
	case "R":
		return moveRight
	}
	log.Fatalf("couldn't parse %s to movement\n", s)
	return moveUp
}

func hexToMovement(hex string) movement {
	switch hex {
	case "0":
		return moveRight
	case "1":
		return moveDown
	case "2":
		return moveLeft
	case "3":
		return moveUp
	}
	log.Fatalf("couldn't parse hex %s to movement\n", hex)
	return moveUp
}

func getLagoonSize(moves []diggerMove) int64 {
	curLoc := util.Coordinate64{X: 0, Y: 0}
	var vertices []util.Coordinate64
	perimeter := int64(0)

	for _, m := range moves {
		switch m.direction {
		case moveUp:
			curLoc.Y -= m.distance
		case moveDown:
			curLoc.Y += m.distance
		case moveRight:
			curLoc.X += m.distance
		case moveLeft:
			curLoc.X -= m.distance
		}
		vertices = append(vertices, curLoc)
		perimeter += m.distance
	}

	return shoelaceArea(vertices) + perimeter/2 + 1
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelaceArea(vertices []util.Coordinate64) int64 {
	area := int64(0)

	j := len(vertices) - 1
	for i := 0; i < len(vertices); i++ {
		area += (vertices[j].Y + vertices[i].Y) * (vertices[j].X - vertices[i].X)
		j = i
	}

	return area / 2
}
