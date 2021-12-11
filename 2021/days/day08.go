package days

import (
	"aoc-2021/util"
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day8() {
	fileContents := util.ReadFileLines("inputs/day08-1.txt")
	fmt.Printf("Part 1: %d\n", day8Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day8Part2(fileContents))
}

func day8Part1(lines []string) int {
	outputs := make([]string, len(lines))
	for i, line := range lines {
		outputs[i] = strings.Split(line, " | ")[1]
	}

	count := 0
	for _, output := range outputs {
		digits := strings.Fields(output)

		for _, digit := range digits {
			segCount := len(digit)
			if segCount == 2 || segCount == 3 || segCount == 4 || segCount == 7 {
				count++
			}
		}
	}

	return count
}

func day8Part2(lines []string) int {
	total := 0

	for _, line := range lines {
		splits := strings.Split(line, " | ")
		inputs := strings.Fields(splits[0])
		outputs := strings.Fields(splits[1])

		for n, seg := range inputs {
			inputs[n] = sortSeg(seg)
		}
		for n, seg := range outputs {
			outputs[n] = sortSeg(seg)
		}

		combined := make([]string, 0)
		combined = append(combined, inputs...)
		combined = append(combined, outputs...)

		digitsToSeg := make(map[int]string)
		segToDigits := make(map[string]int)
		for _, seg := range combined {
			switch len(seg) {
			case 2:
				digitsToSeg[1] = seg
				segToDigits[seg] = 1
			case 3:
				digitsToSeg[7] = seg
				segToDigits[seg] = 7
			case 4:
				digitsToSeg[4] = seg
				segToDigits[seg] = 4
			case 7:
				digitsToSeg[8] = seg
				segToDigits[seg] = 8
			}
		}

		locations := make(map[string]rune)
		for _, seg := range combined {
			if len(seg) == 6 {
				diff := runeDifference(digitsToSeg[1], seg)
				if len(diff) == 1 {
					digitsToSeg[6] = seg
					segToDigits[seg] = 6
					locations["top right"] = diff[0]
					break
				}
			}
		}

		for _, seg := range combined {
			if len(seg) == 5 {
				diff := runeDifference(digitsToSeg[1], seg)
				if len(diff) == 1 {
					if diff[0] == locations["top right"] {
						digitsToSeg[5] = seg
						segToDigits[seg] = 5
					} else {
						digitsToSeg[2] = seg
						segToDigits[seg] = 2
					}
				} else if len(diff) == 0 {
					digitsToSeg[3] = seg
					segToDigits[seg] = 3
				}
			}
		}

		for _, seg := range combined {
			if len(seg) == 6 && seg != digitsToSeg[6] {
				diff := runeDifference(digitsToSeg[5], seg)
				if len(diff) == 0 {
					digitsToSeg[9] = seg
					segToDigits[seg] = 9
				} else if len(diff) == 1 {
					digitsToSeg[0] = seg
					segToDigits[seg] = 0
				}
			}
		}

		var b bytes.Buffer
		for _, output := range outputs {
			intVal := strconv.Itoa(segToDigits[output])
			b.WriteString(intVal)
		}
		total += util.GetIntFromString(b.String())
	}

	return total
}

func sortSeg(s string) string {
	runes := make([]rune, 0)
	for _, c := range s {
		runes = append(runes, c)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func runeDifference(s1 string, s2 string) []rune {
	diff := make([]rune, 0)
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			diff = append(diff, r)
		}
	}
	return diff
}
