package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
	"math"
	"strings"
)

func Day4() {
	fileContents := util.ReadFileLines("inputs/day4-actual.txt")
	fmt.Printf("Part 1: %v\n", day4Part1(fileContents))
	fmt.Printf("Part 2: %v\n", day4Part2(fileContents))
}

func day4Part1(fileContents []string) float64 {
	total := float64(0)

	for _, l := range fileContents {
		ss := strings.Split(l, ": ")
		if len(ss) <= 1 {
			log.Fatalf("line '%s' is not in the right format", l)
		}

		numberHalves := strings.Split(ss[1], " | ")
		if len(ss) <= 1 {
			log.Fatalf("line '%s' is not in the right format", l)
		}

		var winningNums []int
		nums := strings.Fields(numberHalves[0])
		for _, num := range nums {
			winningNums = append(winningNums, util.GetIntFromString(num))
		}

		base := float64(2)
		exp := float64(-1)

		nums = strings.Fields(numberHalves[1])
		for _, num := range nums {
			curNum := util.GetIntFromString(num)
			for _, winNum := range winningNums {
				if curNum == winNum {
					exp += 1
				}
			}
		}

		if exp != -1 {
			total += math.Pow(base, exp)
		}
	}

	return total
}

func day4Part2(fileContents []string) int {
	cardTotals := map[int]int{}
	for i := 0; i < len(fileContents); i++ {
		cardTotals[i] = 1
	}

	for i, l := range fileContents {
		ss := strings.Split(l, ": ")
		if len(ss) <= 1 {
			log.Fatalf("line '%s' is not in the right format", l)
		}

		numberHalves := strings.Split(ss[1], " | ")
		if len(ss) <= 1 {
			log.Fatalf("line '%s' is not in the right format", l)
		}

		var winningNums []int
		nums := strings.Fields(numberHalves[0])
		for _, num := range nums {
			winningNums = append(winningNums, util.GetIntFromString(num))
		}

		winCount := 0
		nums = strings.Fields(numberHalves[1])
		for _, num := range nums {
			curNum := util.GetIntFromString(num)
			for _, winNum := range winningNums {
				if curNum == winNum {
					winCount += 1
				}
			}
		}

		if winCount != 0 {
			for j := 1; j <= winCount; j++ {
				cardTotals[i+j] += cardTotals[i]
			}
		}
	}

	total := 0
	for _, count := range cardTotals {
		total += count
	}
	return total
}
