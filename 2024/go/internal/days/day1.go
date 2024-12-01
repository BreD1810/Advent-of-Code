package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"strconv"
	"strings"
)

func Day1() {
	fileContents := util.ReadFileLines("inputs/day1.txt")
	fmt.Printf("Part 1: %d\n", day1Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day1Part2(fileContents))
}

func day1Part1(lines []string) int {
	left, right := getDay1LeftRight(lines)

	var sum int
	for i, l := range left {
		sum += util.IntAbs(l - right[i])
	}

	return sum
}

func day1Part2(lines []string) int {
	left, right := getDay1LeftRight(lines)

	var sum int
	for _, l := range left {
		var matchCount int
		for _, r := range right {
			if l == r {
				matchCount += 1
			}
		}

		sum += l * matchCount
	}

	return sum
}

func getDay1LeftRight(lines []string) ([]int, []int) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, l := range lines {
		ls := strings.Split(l, " ")
		if len(ls) != 4 {
			panic(fmt.Sprintf("line %d is the wrong format: %s", i, l))
		}

		lInt, err := strconv.Atoi(strings.TrimSpace(ls[0]))
		if err != nil {
			panic(err)
		}
		left[i] = lInt

		rInt, err := strconv.Atoi(strings.TrimSpace(ls[3]))
		if err != nil {
			panic(err)
		}
		right[i] = rInt
	}

	util.SortIntSliceDecending(left)
	util.SortIntSliceDecending(right)

	return left, right
}
