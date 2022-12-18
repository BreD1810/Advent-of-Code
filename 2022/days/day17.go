package days

import (
	"aoc-2022/util"
	"fmt"
)

func Day17() {
	fileContents := util.ReadFileLine("inputs/day17-1.txt")
	fmt.Printf("Part 1: %d\n", day17Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day17Part2(fileContents))
}

const (
	flatRock int = iota
	plusRock
	lRock
	vertRock
	squareRock
)

type fallingRock struct {
	name   string
	height int
	width  int
	shape  [][]bool
}

func day17Part1(fileContents string) int {
	noRocks := 2022
	// noRocks := 10
	jets := parseJets(fileContents)
	// fmt.Println(jets)
	rockTypes := []fallingRock{
		{
			name:   "flat",
			height: 1,
			width:  4,
			shape:  [][]bool{{true, true, true, true}},
		},
		{
			name:   "plus",
			height: 3,
			width:  3,
			shape: [][]bool{
				{false, true, false},
				{true, true, true},
				{false, true, false},
			},
		},
		{
			name:   "l",
			height: 3,
			width:  3,
			shape: [][]bool{
				{true, true, true},
				{false, false, true},
				{false, false, true},
			},
		},
		{
			name:   "vert",
			height: 4,
			width:  1,
			shape: [][]bool{
				{true},
				{true},
				{true},
				{true},
			},
		},
		{
			name:   "square",
			height: 2,
			width:  2,
			shape: [][]bool{
				{true, true},
				{true, true},
			},
		},
	}
	return getRocksFinalHeight(noRocks, jets, rockTypes)
}

// true is right, false is left
func parseJets(s string) []bool {
	jets := make([]bool, len(s)-1)

	for i, c := range s {
		switch c {
		case '<':
			jets[i] = false
		case '>':
			jets[i] = true
		}
	}

	return jets
}

func getRocksFinalHeight(noRocks int, jets []bool, rockTypes []fallingRock) int {
	var chamber [][]bool
	// chamber := make([][]bool, 4)
	// for i := range chamber {
	// 	chamber[i] = make([]bool, 7)
	// }
	var highestPos int
	var curJet int

	for i := 0; i < noRocks; i++ {
		if len(chamber) < highestPos+7 { // Highest is the vertRock
			for j := 0; j < 7; j++ {
				chamber = append(chamber, make([]bool, 7))
			}
		}

		curRock := rockTypes[i%len(rockTypes)]
		curPos := util.Coordinate{X: 2, Y: highestPos + 3}
		// printChamberWithRock(chamber, curRock, curPos)
		fmt.Println()
		// fmt.Printf("Currently dropping a %s at {%d, %d}, highest is %d\n", curRock.name, curPos.X, curPos.Y, highestPos)
		stopped := false
		for !stopped {
			curJet = curJet % len(jets)
			// fmt.Printf("Jet is %v\n", jets[curJet])
			if jets[curJet] && canMoveRight(chamber, curPos, curRock) {
				curPos.X++
			} else if !jets[curJet] && canMoveLeft(chamber, curPos, curRock) {
				curPos.X--
			}
			curJet++

			if canMoveDown(chamber, curPos, curRock) {
				curPos.Y--
			} else {
				chamber = placeRock(chamber, curPos, curRock)
				stopped = true
				// printChamber(chamber)
			}
		}

		if curPos.Y+curRock.height > highestPos {
			highestPos = curPos.Y + curRock.height
		}
	}

	// printChamber(chamber)

	return highestPos
}

func printChamber(chamber [][]bool) {
	for y := len(chamber) - 1; y >= 0; y-- {
		fmt.Printf("%3d", y)
		for _, v := range chamber[y] {
			if v {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
}
func printChamberWithRock(chamber [][]bool, rock fallingRock, pos util.Coordinate) {
	for y := len(chamber) - 1; y >= 0; y-- {
		fmt.Printf("%3d", y)
		for x := range chamber[y] {
			if x >= pos.X && x < pos.X+rock.width && y >= pos.Y && y < pos.Y+rock.height {
				fmt.Printf("%c", 'O')
			} else {
				if chamber[y][x] {
					fmt.Printf("%c", '#')
				} else {
					fmt.Printf("%c", '.')
				}
			}
		}
		fmt.Println()
	}
}

func placeRock(chamber [][]bool, curPos util.Coordinate, rock fallingRock) [][]bool {
	for y := 0; y < rock.height; y++ {
		for x := 0; x < rock.width; x++ {
			if rock.shape[y][x] {
				chamber[curPos.Y+y][curPos.X+x] = true
			}
		}
	}
	return chamber
}

func canMoveDown(chamber [][]bool, curPos util.Coordinate, rock fallingRock) bool {
	if curPos.Y == 0 {
		return false
	}
	newY := curPos.Y - 1
	for y := 0; y < rock.height; y++ {
		for x := 0; x < rock.width; x++ {
			if rock.shape[y][x] && chamber[newY+y][curPos.X+x] {
				return false
			}
		}
	}

	return true
}

func canMoveRight(chamber [][]bool, curPos util.Coordinate, rock fallingRock) bool {
	newX := curPos.X + 1
	if newX+rock.width > 7 {
		return false
	}

	for y := 0; y < rock.height; y++ {
		for x := 0; x < rock.width; x++ {
			if rock.shape[y][x] && chamber[curPos.Y+y][newX+x] {
				return false
			}
		}
	}

	return true
}

func canMoveLeft(chamber [][]bool, curPos util.Coordinate, rock fallingRock) bool {
	if curPos.X == 0 {
		return false
	}
	newX := curPos.X - 1

	for y := 0; y < rock.height; y++ {
		for x := 0; x < rock.width; x++ {
			if rock.shape[y][x] && chamber[curPos.Y+y][newX+x] {
				return false
			}
		}
	}

	return true
}

func day17Part2(fileContents string) int {
	return 0
}
