package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"
)

func Day10() {
	fileContents := util.ReadFileLines("inputs/day10-1.txt")
	fmt.Printf("Part 1: %d\n", day10Part1(fileContents))

	fmt.Println("Part 2:")
	crt := day10Part2(fileContents)
	printCrt(crt)
}

func day10Part1(fileContents []string) int {
	signalStrengths := make([]int, 220)
	cycleNo, x := 1, 1

	for _, l := range fileContents {
		if cycleNo > 220 {
			break
		}

		if l == "noop" {
			signalStrengths[cycleNo-1] = cycleNo * x
			cycleNo++
			continue
		}

		splits := strings.Split(l, " ")
		toAdd := util.GetIntFromString(splits[1])
		for i := 0; i < 2; i++ {
			if cycleNo > 220 {
				break
			}
			signalStrengths[cycleNo-1] = cycleNo * x
			cycleNo++
		}
		x += toAdd
	}

	return signalStrengths[19] + signalStrengths[59] + signalStrengths[99] + signalStrengths[139] + signalStrengths[179] + signalStrengths[219]
}

func day10Part2(fileContents []string) []rune {
	crt := make([]rune, 240)
	for i := range crt {
		crt[i] = '.'
	}

	cycleNo, x := 0, 1

	for _, l := range fileContents {
		if cycleNo >= 240 {
			break
		}

		if l == "noop" {
			curPos := cycleNo % 40
			if curPos == x || curPos == x-1 || curPos == x+1 {
				crt[cycleNo] = '#'
			}
			cycleNo++
			continue
		}

		splits := strings.Split(l, " ")
		toAdd := util.GetIntFromString(splits[1])
		for i := 0; i < 2; i++ {
			if cycleNo >= 240 {
				break
			}
			curPos := cycleNo % 40
			if curPos == x || curPos == x-1 || curPos == x+1 {
				crt[cycleNo] = '#'
			}
			cycleNo++
		}
		x += toAdd
	}

	return crt
}

func printCrt(crt []rune) {
	for i, c := range crt {
		if i%40 == 0 && i != 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", c)
	}
}
