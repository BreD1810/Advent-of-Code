package days

import (
	"aoc-2022/util"
	"fmt"
)

func Day6() {
	fileContents := util.ReadFileLine("inputs/day06-1.txt")
	fmt.Printf("Part 1: %d\n", day6Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day6Part2(fileContents))
}

func day6Part1(signal string) int {
	runes := []rune(signal)
	found := false
	var out int

	for out < len(runes)-4 {
		window := runes[out : out+4]
		for i := 0; i < 4; i++ {
			curRune := runes[out+i]
			if util.CountRuneOccurences(curRune, window) > 1 {
				break
			}
			if i == 3 {
				found = true
			}
		}
		if found {
			break
		}
		out++
	}

	return out + 4
}

func day6Part2(signal string) int {
	runes := []rune(signal)
	found := false
	var out int

	for out < len(runes)-14 {
		window := runes[out : out+14]
		for i := 0; i < 14; i++ {
			curRune := runes[out+i]
			if util.CountRuneOccurences(curRune, window) > 1 {
				break
			}
			if i == 13 {
				found = true
			}
		}
		if found {
			fmt.Println(string(window))
			break
		}
		out++
	}

	return out + 14
}
