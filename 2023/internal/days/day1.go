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
	total := 0
	for _, l := range fileContents {
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
	total := 0
	for _, l := range fileContents {
		positions := map[int][]int{}
		for p := 0; p < 10; p++ {
			var digitPos []int
			s := util.DigitToString(p)

			wordIndex := strings.Index(l, s)
			digitIndex := strings.Index(l, fmt.Sprintf("%d", p))
			if wordIndex >= 0 {
				digitPos = append(digitPos, wordIndex)
			}
			if digitIndex >= 0 {
				digitPos = append(digitPos, digitIndex)
			}

			wordIndex = strings.LastIndex(l, s)
			digitIndex = strings.LastIndex(l, fmt.Sprintf("%d", p))
			if wordIndex >= 0 {
				digitPos = append(digitPos, wordIndex)
			}
			if digitIndex >= 0 {
				digitPos = append(digitPos, digitIndex)
			}

			positions[p] = digitPos
		}

		firstDigit, lastDigit := 0, 0
		firstPosition, lastPosition := len(l), -1
		for d, digitPositions := range positions {
			for _, position := range digitPositions {
				if position < firstPosition {
					firstDigit = d
					firstPosition = position
				}
				if position > lastPosition {
					lastDigit = d
					lastPosition = position
				}
			}
		}

		total += (10 * firstDigit) + lastDigit
	}
	return total
}
