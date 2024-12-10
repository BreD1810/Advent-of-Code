package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"maps"
)

func Day6() {
	fileContents := util.ReadFile2DRune("inputs/day6.txt")
	fmt.Printf("Part 1: %d\n", day6Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day6Part2(fileContents))
}

func day6Part1(input [][]rune) int {
	obstacles := make(map[util.Coordinate]struct{})
	var guardPos util.Coordinate

	for y, row := range input {
		for x, col := range row {
			if col == '#' {
				obstacles[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
			if col == '^' {
				guardPos.X = x
				guardPos.Y = y
			}
		}
	}

	visited := getGuardVisted(guardPos, obstacles, len(input[0]), len(input))

	return len(visited)
}

func getGuardVisted(guardPos util.Coordinate, obstacles map[util.Coordinate]struct{}, xMax, yMax int) map[util.Coordinate]struct{} {
	curDirection := util.MoveUp
	visited := make(map[util.Coordinate]struct{})
	visited[guardPos] = struct{}{}
done:
	for {
		switch curDirection {
		case util.MoveUp:
			if _, ok := obstacles[util.Coordinate{X: guardPos.X, Y: guardPos.Y - 1}]; ok {
				curDirection = curDirection.CycleClockwise()
				continue
			}
			guardPos = util.Coordinate{X: guardPos.X, Y: guardPos.Y - 1}
		case util.MoveRight:
			if _, ok := obstacles[util.Coordinate{X: guardPos.X + 1, Y: guardPos.Y}]; ok {
				curDirection = curDirection.CycleClockwise()
				continue
			}
			guardPos = util.Coordinate{X: guardPos.X + 1, Y: guardPos.Y}
		case util.MoveDown:
			if _, ok := obstacles[util.Coordinate{X: guardPos.X, Y: guardPos.Y + 1}]; ok {
				curDirection = curDirection.CycleClockwise()
				continue
			}
			guardPos = util.Coordinate{X: guardPos.X, Y: guardPos.Y + 1}
		case util.MoveLeft:
			if _, ok := obstacles[util.Coordinate{X: guardPos.X - 1, Y: guardPos.Y}]; ok {
				curDirection = curDirection.CycleClockwise()
				continue
			}
			guardPos = util.Coordinate{X: guardPos.X - 1, Y: guardPos.Y}
		default:
			fmt.Printf("Unknown direction: %v\n", curDirection)
			panic("Unknown direction to move in")
		}

		if guardPos.Y < 0 || guardPos.Y > yMax || guardPos.X < 0 || guardPos.X > xMax {
			break done
		}
		visited[guardPos] = struct{}{}
	}

	return visited
}

func printGuardMap(xMax, yMax int, visited, obstacles map[util.Coordinate]struct{}) {
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if _, ok := obstacles[util.Coordinate{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else if _, ok := visited[util.Coordinate{X: x, Y: y}]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

type visitedWithDirection struct {
	loc       util.Coordinate
	direction util.Movement
}

func day6Part2(input [][]rune) int {
	startingObstacles := make(map[util.Coordinate]struct{})
	var startingGuardPos util.Coordinate

	for y, row := range input {
		for x, col := range row {
			if col == '#' {
				startingObstacles[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
			if col == '^' {
				startingGuardPos.X = x
				startingGuardPos.Y = y
			}
		}
	}

	guardWillVisit := getGuardVisted(startingGuardPos, startingObstacles, len(input[0]), len(input))

	var loopCount int
	for pos := range guardWillVisit {
		if startingGuardPos.X == pos.X && startingGuardPos.Y == pos.Y {
			continue
		}
		guardPos := util.Coordinate{X: startingGuardPos.X, Y: startingGuardPos.Y}
		curDirection := util.MoveUp

		obstacles := maps.Clone(startingObstacles)
		obstacles[util.Coordinate{X: pos.X, Y: pos.Y}] = struct{}{}

		visited := map[visitedWithDirection]struct{}{{loc: guardPos, direction: curDirection}: {}}

	done:
		for {
			visited[visitedWithDirection{loc: guardPos, direction: curDirection}] = struct{}{}

			switch curDirection {
			case util.MoveUp:
				if _, ok := obstacles[util.Coordinate{X: guardPos.X, Y: guardPos.Y - 1}]; ok {
					curDirection = curDirection.CycleClockwise()
					continue
				}
				guardPos = util.Coordinate{X: guardPos.X, Y: guardPos.Y - 1}
			case util.MoveRight:
				if _, ok := obstacles[util.Coordinate{X: guardPos.X + 1, Y: guardPos.Y}]; ok {
					curDirection = curDirection.CycleClockwise()
					continue
				}
				guardPos = util.Coordinate{X: guardPos.X + 1, Y: guardPos.Y}
			case util.MoveDown:
				if _, ok := obstacles[util.Coordinate{X: guardPos.X, Y: guardPos.Y + 1}]; ok {
					curDirection = curDirection.CycleClockwise()
					continue
				}
				guardPos = util.Coordinate{X: guardPos.X, Y: guardPos.Y + 1}
			case util.MoveLeft:
				if _, ok := obstacles[util.Coordinate{X: guardPos.X - 1, Y: guardPos.Y}]; ok {
					curDirection = curDirection.CycleClockwise()
					continue
				}
				guardPos = util.Coordinate{X: guardPos.X - 1, Y: guardPos.Y}
			default:
				fmt.Printf("Unknown direction: %v\n", curDirection)
				panic("Unknown direction to move in")
			}

			if guardPos.Y < 0 || guardPos.Y > len(input) || guardPos.X < 0 || guardPos.X > len(input[0]) {
				break done
			}
			if _, ok := visited[visitedWithDirection{loc: guardPos, direction: curDirection}]; ok {
				loopCount++
				break done
			}
		}
	}

	return loopCount
}
