package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"
)

func Day4() {
	fileContents := util.ReadFileLines("inputs/day04-1.txt")
	fmt.Printf("Part 1: %d\n", day4Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day4Part2(fileContents))
}

func day4Part1(fileContents []string) int {
	var containsCount int

	for _, pairString := range fileContents {
		splits := strings.Split(pairString, ",")

		elfOne := strings.Split(splits[0], "-")
		elfOneStart := util.GetIntFromString(elfOne[0])
		elfOneEnd := util.GetIntFromString(elfOne[1])
		elfTwo := strings.Split(splits[1], "-")
		elfTwoStart := util.GetIntFromString(elfTwo[0])
		elfTwoEnd := util.GetIntFromString(elfTwo[1])

		if (elfOneStart >= elfTwoStart && elfOneEnd <= elfTwoEnd) ||
			(elfOneStart <= elfTwoStart && elfOneEnd >= elfTwoEnd) {
			containsCount++
			continue
		}
	}

	return containsCount
}

func day4Part2(fileContents []string) int {
	var overlapCount int

	for _, pairString := range fileContents {
		splits := strings.Split(pairString, ",")

		elfOne := strings.Split(splits[0], "-")
		elfOneStart := util.GetIntFromString(elfOne[0])
		elfOneEnd := util.GetIntFromString(elfOne[1])
		elfTwo := strings.Split(splits[1], "-")
		elfTwoStart := util.GetIntFromString(elfTwo[0])
		elfTwoEnd := util.GetIntFromString(elfTwo[1])

		if elfOneStart <= elfTwoEnd && elfOneEnd >= elfTwoStart {
			overlapCount++
			continue
		}
	}

	return overlapCount
}
