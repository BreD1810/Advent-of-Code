package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
	"strings"
)

type springState struct {
	springs string
	numbers string
}

var springCache = make(map[springState]int)

func Day12() {
	fileContents := util.ReadFileLines("inputs/day12-actual.txt")
	fmt.Printf("Part 1: %d\n", day12Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day12Part2(fileContents))
}

func day12Part1(fileContents []string) int {
	total := 0
	for _, l := range fileContents {
		springs, countsString, _ := strings.Cut(l, " ")
		var counts []int
		for _, s := range strings.Split(countsString, ",") {
			counts = append(counts, util.GetIntFromString(s))
		}
		total += countPossibleSprings(springs, counts)
	}
	return total
}

func day12Part2(fileContents []string) int {
	total := 0
	for _, l := range fileContents {
		springs, countsString, _ := strings.Cut(l, " ")
		var counts []int
		for _, s := range strings.Split(countsString, ",") {
			counts = append(counts, util.GetIntFromString(s))
		}

		unfoldedSprings, unfoldedCounts := unfoldSprings(springs, counts)
		total += countPossibleSprings(unfoldedSprings, unfoldedCounts)
	}
	return total
}

func countPossibleSprings(springs string, counts []int) int {
	if len(springs) == 0 && len(counts) == 0 {
		return 1
	}

	if len(springs) == 0 {
		return 0
	}

	if val, ok := springCache[springState{springs, fmt.Sprintf("%#v", counts)}]; ok {
		return val
	}

	if springs[0] == '.' {
		res := countPossibleSprings(springs[1:], counts)
		springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = res
		return res
	}

	if springs[0] == '?' {
		res := countPossibleSprings(springs[1:], counts) + countPossibleSprings("#"+springs[1:], counts)
		springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = res
		return res
	}

	if springs[0] == '#' {
		if len(counts) == 0 {
			springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = 0
			return 0
		}

		n := counts[0]
		groupLength := strings.Index(springs, ".")
		if groupLength == -1 {
			groupLength = len(springs)
		}

		if groupLength < n { // Not enough # or ?s
			springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = 0
			return 0
		}

		remaining := springs[n:]
		if len(remaining) == 0 {
			res := countPossibleSprings(remaining, counts[1:])
			springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = res
			return res
		}

		if remaining[0] == '#' { // Too long
			springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = 0
			return 0
		}

		res := countPossibleSprings(remaining[1:], counts[1:])
		springCache[springState{springs, fmt.Sprintf("%#v", counts)}] = res
		return res
	}
	log.Fatal("shouldn't reach here")
	return 0
}

func unfoldSprings(springs string, counts []int) (string, []int) {
	newSprings := strings.Repeat(springs+"?", 4) + springs

	var newCounts []int
	for i := 0; i < 5; i++ {
		for j := 0; j < len(counts); j++ {
			newCounts = append(newCounts, counts[j])
		}
	}

	return newSprings, newCounts
}
