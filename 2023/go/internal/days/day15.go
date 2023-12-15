package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"strings"
)

type hashmapInstruction struct {
	boxNo int
	add   bool
	lenseInfo
}

type lenseInfo struct {
	label       string
	focalLength int
}

func Day15() {
	fileContents := util.ReadFileLine("inputs/day15-example.txt")
	fmt.Printf("Part 1: %d\n", day15Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day15Part2(fileContents))
}

func day15Part1(fileContents string) int {
	ss := strings.Split(fileContents, ",")
	total := 0
	for _, s := range ss {
		total += runHolidayAsciiStringHelper(s)
	}
	return total
}

func day15Part2(fileContents string) int {
	ss := strings.Split(fileContents, ",")
	boxes := make([][]lenseInfo, 256)

	instructions := parseHashmapInstructions(ss)
	for _, ins := range instructions {
		if ins.add {
			b := boxes[ins.boxNo]

			lensLoc := -1
			for i, lens := range b {
				if lens.label == ins.label {
					lensLoc = i
					break
				}
			}

			if lensLoc == -1 {
				b = append(b, lenseInfo{label: ins.label, focalLength: ins.focalLength})
			} else {
				b[lensLoc].focalLength = ins.focalLength
			}

			boxes[ins.boxNo] = b
		} else {
			b := boxes[ins.boxNo]
			lensLoc := -1
			for i, lens := range b {
				if lens.label == ins.label {
					lensLoc = i
					break
				}
			}

			if lensLoc != -1 {
				var newBox []lenseInfo
				for i, lens := range b {
					if i == lensLoc {
						continue
					}
					newBox = append(newBox, lens)
				}
				boxes[ins.boxNo] = newBox
			}
		}
	}

	total := 0
	for i, b := range boxes {
		for j, lens := range b {
			total += (1 + i) * (1 + j) * lens.focalLength
		}
	}
	return total
}

func runHolidayAsciiStringHelper(s string) int {
	val := 0
	for _, c := range s {
		val += int(c)
		val *= 17
		val = val % 256
	}
	return val
}

func parseHashmapInstructions(ss []string) []hashmapInstruction {
	ins := make([]hashmapInstruction, len(ss))
	for i, s := range ss {
		if strings.ContainsRune(s, '=') {
			splits := strings.Split(s, "=")
			label := splits[0]
			boxNo := runHolidayAsciiStringHelper(label)
			focalLength := util.GetIntFromString(splits[1])
			ins[i] = hashmapInstruction{boxNo: boxNo, add: true, lenseInfo: lenseInfo{label: label, focalLength: focalLength}}
		} else {
			label := s[:len(s)-1]
			boxNo := runHolidayAsciiStringHelper(label)
			ins[i] = hashmapInstruction{boxNo: boxNo, add: false, lenseInfo: lenseInfo{label: label}}
		}
	}
	return ins
}
