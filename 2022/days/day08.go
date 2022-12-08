package days

import (
	"aoc-2022/util"
	"fmt"
)

func Day8() {
	fileContents := util.ReadFileIntLines("inputs/day08-1.txt")
	fmt.Printf("Part 1: %d\n", day8Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day8Part2(fileContents))
}

func day8Part1(fileContents [][]int) int {
	// printGrid(fileContents)
	// fmt.Println()
	var hidden int

	for i, row := range fileContents {
		if i == 0 || i == len(fileContents)-1 {
			continue
		}
		for j, val := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}
			hiddenLeft, hiddenRight, hiddenUp, hiddenDown := false, false, false, false
			// LR
			for hor := 0; hor < j; hor++ {
				if fileContents[i][hor] >= val {
					hiddenLeft = true
					break
				}
			}
			for hor := j + 1; hor < len(row); hor++ {
				if fileContents[i][hor] >= val {
					hiddenRight = true
					break
				}
			}
			// UD
			for vert := 0; vert < i; vert++ {
				if fileContents[vert][j] >= val {
					hiddenUp = true
					break
				}
			}
			for vert := i + 1; vert < len(fileContents); vert++ {
				if fileContents[vert][j] >= val {
					hiddenDown = true
					break
				}
			}

			if hiddenLeft && hiddenRight && hiddenDown && hiddenUp {
				hidden++
			}
		}
	}

	l := len(fileContents)
	return (l * l) - hidden
}

func day8Part2(fileContents [][]int) int {
	var largestScenic int

	for i, row := range fileContents {
		for j, val := range row {
			count, curScenic := 0, 1
			// LR
			for hor := j - 1; hor >= 0; hor-- {
				if fileContents[i][hor] >= val {
					count++
					break
				}
				count++
			}
			curScenic *= count
			count = 0
			for hor := j + 1; hor < len(row); hor++ {
				if fileContents[i][hor] >= val {
					count++
					break
				}
				count++
			}
			curScenic *= count
			// UD
			count = 0
			for vert := i - 1; vert >= 0; vert-- {
				if fileContents[vert][j] >= val {
					count++
					break
				}
				count++
			}
			curScenic *= count
			count = 0
			for vert := i + 1; vert < len(fileContents); vert++ {
				if fileContents[vert][j] >= val {
					count++
					break
				}
				count++
			}
			curScenic *= count
			if curScenic > largestScenic {
				largestScenic = curScenic
			}
		}
	}

	return largestScenic
}

func printGrid(fileContents [][]int) {
	for _, row := range fileContents {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		fmt.Printf("\n")
	}
}
