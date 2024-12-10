package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day4() {
	fileContents := util.ReadFile2DRune("inputs/day4.txt")
	fmt.Printf("Part 1: %d\n", day4Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day4Part2(fileContents))
}

func day4Part1(input [][]rune) int {
	var count int
	for y := range input {
		for x, curRune := range input[y] {
			if curRune == 'X' {
				count += countXmas(input, x, y)
			}
		}
	}
	return count
}

func countXmas(input [][]rune, x, y int) int {
	var count int
	// Check left
	if x >= 3 && input[y][x-1] == 'M' && input[y][x-2] == 'A' && input[y][x-3] == 'S' {
		count++
	}
	// Check right
	if x < len(input[y])-3 && input[y][x+1] == 'M' && input[y][x+2] == 'A' && input[y][x+3] == 'S' {
		count++
	}
	// Check up
	if y >= 3 && input[y-1][x] == 'M' && input[y-2][x] == 'A' && input[y-3][x] == 'S' {
		count++
	}
	// Check down
	if y < len(input)-3 && input[y+1][x] == 'M' && input[y+2][x] == 'A' && input[y+3][x] == 'S' {
		count++
	}
	// Check up right
	if x < len(input[y])-3 && y >= 3 && input[y-1][x+1] == 'M' && input[y-2][x+2] == 'A' && input[y-3][x+3] == 'S' {
		count++
	}

	// Check up left
	if x >= 3 && y >= 3 && input[y-1][x-1] == 'M' && input[y-2][x-2] == 'A' && input[y-3][x-3] == 'S' {
		count++
	}
	// Check down right
	if x < len(input[y])-3 && y < len(input)-3 && input[y+1][x+1] == 'M' && input[y+2][x+2] == 'A' && input[y+3][x+3] == 'S' {
		count++
	}
	// Check down left
	if x >= 3 && y < len(input)-3 && input[y+1][x-1] == 'M' && input[y+2][x-2] == 'A' && input[y+3][x-3] == 'S' {
		count++
	}

	return count
}

func day4Part2(input [][]rune) int {
	var count int
	for y := range input {
		for x, _ := range input[y] {
			if isCrossMas(input, x, y) {
				count += 1
			}
		}
	}
	return count
}

func isCrossMas(input [][]rune, x, y int) bool {
	if y >= len(input)-2 || x >= len(input[y])-2 {
		return false
	}
	if input[y][x] != 'M' && input[y][x] != 'S' {
		return false
	}
	if input[y][x] == 'M' && (input[y+1][x+1] != 'A' || input[y+2][x+2] != 'S') {
		return false
	}
	if input[y][x] == 'S' && (input[y+1][x+1] != 'A' || input[y+2][x+2] != 'M') {
		return false
	}
	if input[y+2][x] != 'M' && input[y+2][x] != 'S' {
		return false
	}
	if input[y+2][x] == 'M' && (input[y+1][x+1] != 'A' || input[y][x+2] != 'S') {
		return false
	}
	if input[y+2][x] == 'S' && (input[y+1][x+1] != 'A' || input[y][x+2] != 'M') {
		return false
	}
	return true
}
