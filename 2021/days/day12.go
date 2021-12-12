package days

import (
	"aoc-2021/util"
	"fmt"
	"strings"
)

const (
	startCave = "start"
	endCave   = "end"
)

func Day12() {
	fileContents := util.ReadFileLines("inputs/day12-1.txt")
	// fileContents := util.ReadFileLines("inputs/test.txt")
	// fileContents := util.ReadFileLines("inputs/test-2.txt")
	// fileContents := util.ReadFileLines("inputs/test-3.txt")
	fmt.Printf("Part 1: %d\n", day12Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day12Part2(fileContents))
}

func day12Part1(lines []string) int {
	caveMap := map[string][]string{}

	for _, line := range lines {
		caves := strings.Split(line, "-")
		caveMap[caves[0]] = append(caveMap[caves[0]], caves[1])
		caveMap[caves[1]] = append(caveMap[caves[1]], caves[0])
	}

	pathCount := 0

	for _, cave := range caveMap[startCave] {
		startMap := map[string]bool{"start": true}
		pathCount += findPathCount(caveMap, cave, startMap)
	}

	return pathCount
}

func day12Part2(lines []string) int {
	caveMap := map[string][]string{}

	for _, line := range lines {
		caves := strings.Split(line, "-")
		caveMap[caves[0]] = append(caveMap[caves[0]], caves[1])
		caveMap[caves[1]] = append(caveMap[caves[1]], caves[0])
	}

	completePaths := map[string]bool{}
	for revisitCave := range caveMap {
		if revisitCave == startCave || revisitCave == endCave || strings.ToLower(revisitCave) != revisitCave {
			continue
		}

		for _, startCave := range caveMap[startCave] {
			findPathCountWithOneRevisit(caveMap, startCave, revisitCave, false, "start", completePaths)
		}
	}

	return len(completePaths)
}

func findPathCount(caves map[string][]string, startCave string, seenCaves map[string]bool) int {
	if startCave == endCave {
		return 1
	}

	count := 0
	seenCaves[startCave] = true
	for _, next := range caves[startCave] {
		_, visited := seenCaves[next]
		// Can't revisit small caves
		if visited && strings.ToLower(next) == next {
			continue
		}

		// Need to clone map so it's not passed by reference
		newSeenCaves := map[string]bool{}
		for k, v := range seenCaves {
			newSeenCaves[k] = v
		}

		count += findPathCount(caves, next, newSeenCaves)
	}

	return count
}

func findPathCountWithOneRevisit(caves map[string][]string, startCave string, toRevisit string, haveRevisited bool, path string, completePaths map[string]bool) {
	if startCave == endCave {
		completePaths[path] = true
		return
	}

	for _, next := range caves[startCave] {
		visited := strings.Contains(path, next)

		// Can only revisit small caves once
		visitingAgain := haveRevisited
		if visited && strings.ToLower(next) == next {
			if next == toRevisit {
				if haveRevisited {
					continue
				}

				visitingAgain = true
			} else {
				continue
			}
		}

		newPath := fmt.Sprintf("%s,%s", path, startCave)
		findPathCountWithOneRevisit(caves, next, toRevisit, visitingAgain, newPath, completePaths)
	}
}
