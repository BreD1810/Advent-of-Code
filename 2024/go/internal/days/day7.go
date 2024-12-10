package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/mowshon/iterium"
)

func Day7() {
	fileContents := util.ReadFileLines("inputs/day7.txt")
	fmt.Printf("Part 1: %d\n", day7Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day7Part2(fileContents))
}

func day7Part1(input []string) int {
	var answer int

	for n, line := range input {
		totalSplit := strings.Split(line, ": ")
		if len(totalSplit) != 2 {
			fmt.Printf("Line %d is the wrong format: %q\n", n, line)
			panic("Line is the wrong format")
		}

		_, err := strconv.Atoi(totalSplit[0])
		totalWanted, err := strconv.Atoi(totalSplit[0])
		if err != nil {
			panic(err)
		}

		componentSplit := strings.Split(totalSplit[1], " ")
		components := make([]int, len(componentSplit))
		for i, c := range componentSplit {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			components[i] = v
		}

		noOperators := len(components) - 1
		operatorCombos, err := iterium.Product([]int{0, 1}, noOperators).Slice()
		if err != nil {
			panic(err)
		}
		for _, combo := range operatorCombos {
			runningTotal := components[0]
			for i := 1; i < len(components); i++ {
				if combo[i-1] == 0 {
					runningTotal += components[i]
				} else {
					runningTotal *= components[i]
				}
			}
			if runningTotal == totalWanted {
				answer += totalWanted
				break
			}
		}
	}

	return answer
}

func day7Part2(input []string) int {
	var answer int

	existingProducts := make(map[int][][]int)

	for n, line := range input {
		totalSplit := strings.Split(line, ": ")
		if len(totalSplit) != 2 {
			fmt.Printf("Line %d is the wrong format: %q\n", n, line)
			panic("Line is the wrong format")
		}

		_, err := strconv.Atoi(totalSplit[0])
		totalWanted, err := strconv.Atoi(totalSplit[0])
		if err != nil {
			panic(err)
		}

		componentSplit := strings.Split(totalSplit[1], " ")
		components := make([]int, len(componentSplit))
		for i, c := range componentSplit {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			components[i] = v
		}

		noOperators := len(components) - 1
		var operatorCombos [][]int
		var ok bool
		if operatorCombos, ok = existingProducts[noOperators]; !ok {
			operatorCombos, err = iterium.Product([]int{0, 1, 2}, noOperators).Slice()
			if err != nil {
				panic(err)
			}
		}
		for _, combo := range operatorCombos {
			runningTotal := components[0]
			for i := 1; i < len(components); i++ {
				curVal := components[i]
				switch combo[i-1] {
				case 0:
					runningTotal += curVal
				case 1:
					runningTotal *= curVal
				default:
					runningTotal, err = strconv.Atoi(fmt.Sprintf("%d%d", runningTotal, curVal))
					if err != nil {
						panic(err)
					}
				}
			}
			if runningTotal == totalWanted {
				answer += totalWanted
				break
			}
		}
	}

	return answer
}

func calculateIntPadding(n int) int {
	p := 10
	for p < n {
		p *= 10
	}
	return p
}
