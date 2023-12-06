package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"strings"
)

func Day6() {
	fileContents := util.ReadFileLines("inputs/day6-actual.txt")
	fmt.Printf("Part 1: %d\n", day6Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day6Part2(fileContents))
}

func day6Part1(fileContents []string) int {
	timeString := strings.TrimPrefix(fileContents[0], "Time:")
	timess := strings.Fields(timeString)
	times := make([]int, len(timess))
	for i, t := range timess {
		times[i] = util.GetIntFromString(t)
	}

	distanceString := strings.TrimPrefix(fileContents[1], "Distance:")
	distancess := strings.Fields(distanceString)
	distances := make([]int, len(distancess))
	for i, d := range distancess {
		distances[i] = util.GetIntFromString(d)
	}

	counts := make(map[int]int)
	for i, t := range times {
		timeCount := 0
		distToBeat := distances[i]
		for j := 1; j <= t; j++ {
			curDist := j * (t - j)
			if curDist > distToBeat {
				timeCount++
			}
		}
		counts[i] = timeCount
	}

	total := 1
	for _, v := range counts {
		total *= v
	}
	return total
}

func day6Part2(fileContents []string) int {
	timeString := strings.TrimPrefix(fileContents[0], "Time:")
	time := util.GetIntFromString(strings.ReplaceAll(timeString, " ", ""))

	distanceString := strings.TrimPrefix(fileContents[1], "Distance:")
	distance := util.GetIntFromString(strings.ReplaceAll(distanceString, " ", ""))

	winCount := 0
	for j := 1; j <= time; j++ {
		curDist := j * (time - j)
		if curDist > distance {
			winCount++
		}
	}

	return winCount
}
