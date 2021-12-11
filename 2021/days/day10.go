package days

import (
	"aoc-2021/util"
	"fmt"
	"sort"
	"strings"
)

func Day10() {
	fileContents := util.ReadFileLines("inputs/day10-1.txt")
	// fileContents := util.ReadFileLines("inputs/test-10.txt")
	fmt.Printf("Part 1: %d\n", day10Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day10Part2(fileContents))
}

func day10Part1(lines []string) int {
	charScores := make(map[rune]int)
	charScores[')'] = 3
	charScores[']'] = 57
	charScores['}'] = 1197
	charScores['>'] = 25137

	total := 0
	for _, line := range lines {
		invalidCloser := getInvalidCloser(line)
		if invalidCloser != '\n' {
			total += charScores[invalidCloser]
		}
	}

	return total
}

func day10Part2(lines []string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		completionScore := getCompletionScore(line)
		if completionScore != 0 {
			scores = append(scores, completionScore)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func getInvalidCloser(line string) rune {
	// Easier than []rune as no contains method
	openChars := "([{<"
	closeChars := ")]}>"

	lineStack := make([]rune, 0)
	for _, c := range line {
		if strings.ContainsRune(openChars, c) {
			lineStack = append(lineStack, c)
		} else if strings.ContainsRune(closeChars, c) {
			actual := lineStack[len(lineStack)-1]
			expected := getOpenerForCloser(c, openChars, closeChars)
			if actual != expected {
				return c
			}
			lineStack = lineStack[:len(lineStack)-1]
		} else {
			fmt.Printf("Couldn't find character %q\n", c)
			return '0'
		}
	}
	return '\n'
}

func getOpenerForCloser(c rune, openChars string, closeChars string) rune {
	return rune(openChars[strings.IndexRune(closeChars, c)])
}

func getCompletionScore(line string) int {
	openChars := "([{<"
	closeChars := ")]}>"

	charScores := make(map[rune]int)
	charScores['('] = 1
	charScores['['] = 2
	charScores['{'] = 3
	charScores['<'] = 4

	lineStack := make([]rune, 0)
	for _, c := range line {
		if strings.ContainsRune(openChars, c) {
			lineStack = append(lineStack, c)
		} else if strings.ContainsRune(closeChars, c) {
			actual := lineStack[len(lineStack)-1]
			expected := getOpenerForCloser(c, openChars, closeChars)
			if actual != expected {
				return 0
			}
			lineStack = lineStack[:len(lineStack)-1]
		} else {
			fmt.Printf("Couldn't find character %q\n", c)
			return 0
		}
	}

	total := 0
	for i := len(lineStack) - 1; i >= 0; i-- {
		total *= 5
		total += charScores[lineStack[i]]
	}
	return total
}
