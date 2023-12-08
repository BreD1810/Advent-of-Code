package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
)

func Day8() {
	fileContents := util.ReadFileLines("inputs/day8-actual.txt")
	fmt.Printf("Part 1: %d\n", day8Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day8Part2(fileContents))
}

type LR int

const (
	left LR = iota
	right
)

func day8Part1(fileContents []string) int {
	lrs := parseLRs(fileContents[0])

	lrNodeMap := make(map[string][]string)

	for _, l := range fileContents[2:] {
		source := l[0:3]
		lft := l[7:10]
		rght := l[12:15]
		lrNodeMap[source] = []string{lft, rght}
	}

	found := false
	steps := 0
	curNodeSource := "AAA"
	for !found {
		if lrs[steps%len(lrs)] == left {
			curNodeSource = lrNodeMap[curNodeSource][0]
		} else {
			curNodeSource = lrNodeMap[curNodeSource][1]
		}
		if curNodeSource == "ZZZ" {
			found = true
		}
		steps++
	}

	return steps
}

func parseLRs(s string) []LR {
	cs := []rune(s)
	lrs := make([]LR, len(cs))
	for i, c := range cs {
		switch c {
		case 'L':
			lrs[i] = left
		case 'R':
			lrs[i] = right
		default:
			log.Fatalf("failed to parse LR %c\n", c)
		}
	}
	return lrs
}

func day8Part2(fileContents []string) int {
	lrs := parseLRs(fileContents[0])

	lrNodeMap := make(map[string][]string)

	var ghostLocs []string
	for _, l := range fileContents[2:] {
		source := l[0:3]
		if l[2] == 'A' {
			ghostLocs = append(ghostLocs, source)
		}
		lft := l[7:10]
		rght := l[12:15]
		lrNodeMap[source] = []string{lft, rght}
	}

	foundCounts := make([]int, len(ghostLocs))
	steps := 0
	for {
		steps++
		if lrs[(steps-1)%len(lrs)] == left {
			newGhostLocs := make([]string, len(ghostLocs))
			for i, l := range ghostLocs {
				newGhostLocs[i] = lrNodeMap[l][0]
			}
			ghostLocs = newGhostLocs
		} else {
			newGhostLocs := make([]string, len(ghostLocs))
			for i, l := range ghostLocs {
				newGhostLocs[i] = lrNodeMap[l][1]
			}
			ghostLocs = newGhostLocs
		}

		for i, l := range ghostLocs {
			if l[2] == 'Z' && foundCounts[i] == 0 {
				foundCounts[i] = steps
			}
		}

		allFound := true
		for _, found := range foundCounts {
			if found == 0 {
				allFound = false
				break
			}
		}
		if allFound {
			break
		}
	}

	return util.GetLeastCommonMultiple(foundCounts)
}
