package days

import (
	"aoc-2021/util"
	"fmt"
	"strings"
)

const (
	resetVal   = 6
	defaultVal = 8
	noDays     = 80
	noDays2    = 256
)

func Day6() {
	fileContents := util.ReadFileLines("inputs/day06-1.txt")
	fmt.Printf("Part 1: %d\n", day6Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day6Part2(fileContents))
}

func day6Part1(lines []string) int {
	input := lines[0]
	fishStrings := strings.Split(input, ",")
	fishes := make([]int, len(fishStrings))
	for i, fish := range fishStrings {
		fishes[i] = util.GetIntFromString(fish)
	}

	for day := 1; day <= noDays; day++ {
		dayLen := len(fishes)
		for i := 0; i < dayLen; i++ {
			fishes[i]--
			if fishes[i] < 0 {
				fishes[i] = resetVal
				fishes = append(fishes, defaultVal)
			}
		}
	}

	return len(fishes)
}

func day6Part2(lines []string) int {
	input := lines[0]
	fishStrings := strings.Split(input, ",")
	fishMap := make(map[int]int)

	for _, fish := range fishStrings {
		currentFish := util.GetIntFromString(fish)
		fishMap[currentFish]++
	}

	// fmt.Printf("Initial state: %v\n", fishMap)

	for day := 1; day <= noDays2; day++ {
		spawners := fishMap[0]
		for i := 0; i < defaultVal; i++ {
			fishMap[i] = fishMap[i+1]
		}
		fishMap[resetVal] += spawners
		fishMap[defaultVal] = spawners
		// fmt.Printf("After %d days: %v\n", day, fishMap)
	}

	total := 0
	for _, count := range fishMap {
		total += count
	}
	return total
}
