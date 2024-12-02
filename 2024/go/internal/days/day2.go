package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day2() {
	fileContents := util.ReadFileIntLines("inputs/day2.txt")
	fmt.Printf("Part 1: %d\n", day2Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day2Part2(fileContents))
}

func day2Part1(reports [][]int) int {
	var safe int

	for _, r := range reports {
		if isSafe(r) {
			safe++
		}
	}

	return safe
}

func isSafe(r []int) bool {
	currentIsSafe := true

	increasing := r[0] < r[1] // This may not work if 2 values next to each other are the same?

	for i := 1; i < len(r); i++ {
		var diff int
		if increasing {
			diff = r[i] - r[i-1]
		}
		if !increasing {
			diff = r[i-1] - r[i]
		}
		if diff < 1 || diff > 3 {
			currentIsSafe = false
			break
		}
	}

	return currentIsSafe
}

func day2Part2(reports [][]int) int {
	var safe int

	for _, curReport := range reports {
		if isSafe(curReport) {
			safe += 1
			continue
		}

		for i := 0; i < len(curReport); i++ {
			newReports := make([]int, len(curReport))
			copy(newReports, curReport)
			if isSafe(append(newReports[:i], newReports[i+1:]...)) {
				safe++
				break
			}
		}
	}

	return safe
}

// func day2Part2(reports [][]int) int {
// 	var safe int
//
// 	for n, r := range reports {
// 		// if n != 353 {
// 		// 	continue
// 		// }
// 		currentIsSafe := true
// 		fmt.Printf("%d, %d\n", r[0], r[1])
//
// 		prevVal := r[0]
// 		var increasing, dampened bool
// 		// increasing := prevVal < r[1] // This may not work if 2 values next to each other are the same?
//
// 		for i := 1; i < len(r); i++ {
// 			if prevVal == r[0] {
// 				fmt.Printf("increasing: %t\n", increasing)
// 				increasing = prevVal < r[1] // This may not work if 2 values next to each other are the same?
// 			}
// 			var diff int
// 			if increasing {
// 				diff = r[i] - prevVal
// 			}
// 			if !increasing {
// 				diff = prevVal - r[i]
// 			}
// 			// fmt.Printf("diff: %d\n", diff)
// 			if diff < 1 || diff > 3 {
// 				if i == 1 {
// 					fmt.Printf("current line: %d\n", n)
// 				}
// 				if dampened {
// 					currentIsSafe = false
// 					fmt.Printf("%d, %d ", r[i-1], r[i])
// 					fmt.Println("not safe")
// 					break
// 				}
// 				dampened = true
// 				continue
// 			}
// 			prevVal = r[i]
// 		}
//
// 		if currentIsSafe {
// 			safe += 1
// 		}
// 	}
//
// 	return safe
// }
