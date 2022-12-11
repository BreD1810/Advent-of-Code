package days

import (
	"aoc-2022/util"
	"fmt"
	"log"
	"strings"
)

func Day11() {
	fileContents := util.ReadFileLines("inputs/day011-1.txt")
	fmt.Printf("Part 1: %d\n", day11Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day11Part2(fileContents))
}

type monkey struct {
	items         []int
	operationType monkeyOperation
	operationVal  int
	testDivisor   int
	trueDest      int
	falseDest     int
}

type monkeyOperation int

const (
	add monkeyOperation = iota
	multiply
	square
)

func day11Part1(fileContents []string) int {
	monkeys, _ := parseMonkeys(fileContents)
	inspections := countInspections(monkeys, 20)
	util.SortIntSliceDecending(inspections)
	return inspections[0] * inspections[1]
}

func parseMonkeys(fileContents []string) ([]monkey, int) {
	monkeys := make([]monkey, (len(fileContents)/7)+1)
	var monkeyCounter int
	biggestMod := 1

	for _, l := range fileContents {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "Monkey") {
			monkeyCounter = int(l[7] - '0')
			continue
		}

		if strings.HasPrefix(l, "Starting items") {
			l = strings.TrimPrefix(l, "Starting items: ")
			itemStrings := strings.Split(l, ", ")
			items := make([]int, len(itemStrings))
			for i, v := range itemStrings {
				items[i] = util.GetIntFromString(v)
			}
			monkeys[monkeyCounter].items = items
			continue
		}

		if strings.HasPrefix(l, "Operation") {
			operationRune := l[21]
			switch operationRune {
			case '+':
				monkeys[monkeyCounter].operationType = add
			case '*':
				monkeys[monkeyCounter].operationType = multiply
			default:
				log.Fatalf("Unknown operation %c\n", operationRune)
			}

			splits := strings.Split(l, string(operationRune)+" ")
			if splits[1] == "old" {
				monkeys[monkeyCounter].operationType = square
			} else {
				monkeys[monkeyCounter].operationVal = util.GetIntFromString(splits[1])
			}
			continue
		}

		if strings.HasPrefix(l, "Test") {
			l = strings.TrimPrefix(l, "Test: divisible by ")
			v := util.GetIntFromString(l)
			monkeys[monkeyCounter].testDivisor = v
			biggestMod *= v
			continue
		}

		if strings.HasPrefix(l, "If true") {
			l = strings.TrimPrefix(l, "If true: throw to monkey ")
			monkeys[monkeyCounter].trueDest = util.GetIntFromString(l)
			continue
		}

		if strings.HasPrefix(l, "If false") {
			l = strings.TrimPrefix(l, "If false: throw to monkey ")
			monkeys[monkeyCounter].falseDest = util.GetIntFromString(l)
			continue
		}
	}

	return monkeys, biggestMod
}

// countInspections goes through noRounds rounds, and returns the number of inspections for each monkey
func countInspections(monkeys []monkey, noRounds int) []int {
	inspections := make([]int, len(monkeys))

	for r := 0; r < noRounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			curMonkey := &monkeys[m]
			for _, worry := range curMonkey.items {
				switch curMonkey.operationType {
				case add:
					worry += curMonkey.operationVal
				case multiply:
					worry *= curMonkey.operationVal
				case square:
					worry *= worry
				}

				worry = worry / 3
				inspections[m] += 1

				if worry%curMonkey.testDivisor == 0 {
					monkeys[curMonkey.trueDest].items = append(monkeys[curMonkey.trueDest].items, worry)
				} else {
					monkeys[curMonkey.falseDest].items = append(monkeys[curMonkey.falseDest].items, worry)
				}
			}
			curMonkey.items = make([]int, 0)
		}
	}

	return inspections
}

func countInspections2(monkeys []monkey, noRounds int, biggestMod int) []int {
	inspections := make([]int, len(monkeys))

	for r := 0; r < noRounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			curMonkey := &monkeys[m]
			for _, worry := range curMonkey.items {
				switch curMonkey.operationType {
				case add:
					worry += curMonkey.operationVal
				case multiply:
					worry *= curMonkey.operationVal
				case square:
					worry *= worry
				}

				inspections[m] += 1

				if worry%curMonkey.testDivisor == 0 {
					monkeys[curMonkey.trueDest].items = append(monkeys[curMonkey.trueDest].items, worry%biggestMod)
				} else {
					monkeys[curMonkey.falseDest].items = append(monkeys[curMonkey.falseDest].items, worry%biggestMod)
				}
			}
			curMonkey.items = make([]int, 0)
		}
	}

	return inspections
}

func day11Part2(fileContents []string) int {
	monkeys, biggestMod := parseMonkeys(fileContents)
	inspections := countInspections2(monkeys, 10000, biggestMod)
	util.SortIntSliceDecending(inspections)
	return inspections[0] * inspections[1]
}
