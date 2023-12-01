package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"strings"
	"unicode"
)

func Day1() {
	fileContents := util.ReadFileLines("inputs/day1.txt")
	fmt.Printf("Part 1: %d\n", day1Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day1Part2(fileContents))
}

func day1Part1(fileContents []string) int {
	return solver(fileContents)
}

func solver(lines []string) int {
	total := 0
	for _, l := range lines {
		firstDigit, lastDigit := -1, -1
		for _, c := range l {
			if unicode.IsDigit(c) {
				digitValue := int(c - '0')
				if firstDigit == -1 {
					firstDigit = digitValue
				}
				lastDigit = digitValue
			}
		}
		total += (10 * firstDigit) + lastDigit
	}
	return total
}

func day1Part2(fileContents []string) int {
	var lines []string
	for _, l := range fileContents {
		fmt.Println(l)
		l = strings.ReplaceAll(l, "one", "one1one")
		l = strings.ReplaceAll(l, "two", "two2two")
		l = strings.ReplaceAll(l, "three", "three3three")
		l = strings.ReplaceAll(l, "four", "four4four")
		l = strings.ReplaceAll(l, "five", "five5five")
		l = strings.ReplaceAll(l, "six", "six6six")
		l = strings.ReplaceAll(l, "seven", "seven7seven")
		l = strings.ReplaceAll(l, "eight", "eight8eight")
		l = strings.ReplaceAll(l, "nine", "nine9nine")
		fmt.Println(l)
		lines = append(lines, l)
	}
	return solver(lines)
}
