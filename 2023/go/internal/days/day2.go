package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"regexp"
	"strings"
)

type cubeCounts struct {
	red   int
	green int
	blue  int
}

var (
	redRegex   = regexp.MustCompile(`(\d+) red`)
	greenRegex = regexp.MustCompile(`(\d+) green`)
	blueRegex  = regexp.MustCompile(`(\d+) blue`)
)

func Day2() {
	fileContents := util.ReadFileLines("inputs/day2.txt")
	fmt.Printf("Part 1: %d\n", day2Part1(fileContents, cubeCounts{red: 12, green: 13, blue: 14}))
	fmt.Printf("Part 2: %d\n", day2Part2(fileContents))
}

func day2Part1(fileContents []string, colours cubeCounts) int {
	sum := 0

	for i, l := range fileContents {
		ss := strings.Split(l, ";")
		isPossible := true
		for _, s := range ss {
			redMatches := redRegex.FindStringSubmatch(s)
			if len(redMatches) > 1 && util.GetIntFromString(redMatches[1]) > colours.red {
				isPossible = false
				break
			}

			greenMatches := greenRegex.FindStringSubmatch(s)
			if len(greenMatches) > 1 && util.GetIntFromString(greenMatches[1]) > colours.green {
				isPossible = false
				break
			}

			blueMatches := blueRegex.FindStringSubmatch(s)
			if len(blueMatches) > 1 && util.GetIntFromString(blueMatches[1]) > colours.blue {
				isPossible = false
				break
			}
		}
		if isPossible {
			sum += i + 1
		}
	}

	return sum
}

func day2Part2(fileContents []string) int {
	powerSum := 0

	for _, l := range fileContents {
		colourCounts := cubeCounts{red: 0, green: 0, blue: 0}
		ss := strings.Split(l, ";")
		for _, s := range ss {
			redMatches := redRegex.FindStringSubmatch(s)
			if len(redMatches) > 1 && util.GetIntFromString(redMatches[1]) > colourCounts.red {
				colourCounts.red = util.GetIntFromString(redMatches[1])
			}

			greenMatches := greenRegex.FindStringSubmatch(s)
			if len(greenMatches) > 1 && util.GetIntFromString(greenMatches[1]) > colourCounts.green {
				colourCounts.green = util.GetIntFromString(greenMatches[1])
			}

			blueMatches := blueRegex.FindStringSubmatch(s)
			if len(blueMatches) > 1 && util.GetIntFromString(blueMatches[1]) > colourCounts.blue {
				colourCounts.blue = util.GetIntFromString(blueMatches[1])
			}
		}
		powerSum += colourCounts.red * colourCounts.green * colourCounts.blue
	}

	return powerSum
}
