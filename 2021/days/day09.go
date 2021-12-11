package days

import (
	"aoc-2021/util"
	"fmt"
	"sort"
)

func Day9() {
	fileContents := util.ReadFileLines("inputs/day09-1.txt")
	fmt.Printf("Part 1: %d\n", day9Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day9Part2(fileContents))
}

func day9Part1(lines []string) int {
	rows := make([][]int, len(lines))
	for i, line := range lines {
		rows[i] = make([]int, len(line))
		for j, digit := range line {
			rows[i][j] = util.GetIntFromString(string(digit))
		}
	}

	sum := 0
	for y, row := range rows {
		for x, val := range row {
			if isLow(rows, x, y) {
				sum += val + 1
			}
		}
	}

	return sum
}

func day9Part2(lines []string) int {
	rows := make([][]int, len(lines))
	for i, line := range lines {
		rows[i] = make([]int, len(line))
		for j, digit := range line {
			rows[i][j] = util.GetIntFromString(string(digit))
		}
	}

	basinSizes := make([]int, 0)
	for y, row := range rows {
		for x := range row {
			if isLow(rows, x, y) {
				basinSizes = append(basinSizes, getBasinSize(rows, x, y, 0))
			}
		}
	}

	sort.Ints(basinSizes)
	noBasins := len(basinSizes)
	return basinSizes[noBasins-1] * basinSizes[noBasins-2] * basinSizes[noBasins-3]
}

func isLow(rows [][]int, x int, y int) bool {
	val := rows[y][x]
	if x > 0 && rows[y][x-1] <= val {
		return false
	}
	if x < len(rows[y])-1 && rows[y][x+1] <= val {
		return false
	}
	if y > 0 && rows[y-1][x] <= val {
		return false
	}
	if y < len(rows)-1 && rows[y+1][x] <= val {
		return false
	}
	return true
}

func getBasinSize(rows [][]int, x int, y int, size int) int {
	if rows[y][x] == 9 {
		return size
	}

	size++
	rows[y][x] = 9
	if x > 0 {
		size = getBasinSize(rows, x-1, y, size)
	}
	if x < len(rows[y])-1 {
		size = getBasinSize(rows, x+1, y, size)
	}
	if y > 0 {
		size = getBasinSize(rows, x, y-1, size)
	}
	if y < len(rows)-1 {
		size = getBasinSize(rows, x, y+1, size)
	}

	return size
}
