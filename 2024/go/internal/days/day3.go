package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	fileContents := util.ReadFileLines("inputs/day3.txt")
	fmt.Printf("Part 1: %d\n", day3Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day3Part2(fileContents))
}

func day3Part1(lines []string) int {
	var total int
	r := regexp.MustCompile(`(mul\(\d+,\d+\))`)

	for _, l := range lines {
		matches := r.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			total += getMulTotal(m)
		}
	}

	return total
}

func getMulTotal(match []string) int {
	s := match[0]
	s = strings.TrimPrefix(s, "mul(")
	s = strings.TrimSuffix(s, ")")
	ss := strings.Split(s, ",")
	a, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(err)
	}
	return a * b
}

func day3Part2(lines []string) int {
	input := strings.Join(lines, "")

	mulReg := regexp.MustCompile(`(mul\(\d+,\d+\))`)
	doReg := regexp.MustCompile(`(do\(\))`)
	dontReg := regexp.MustCompile(`(don't\(\))`)
	var mulIndexes, doIndexes, dontIndexes []int

	// 1st number is overall match start index
	// 2nd number is overall match end index + 1
	// 3rd number is the inner match start index
	// 4th number is the inner match end index + 1
	for _, m := range mulReg.FindAllStringSubmatchIndex(input, -1) {
		mulIndexes = append(mulIndexes, m[2])
	}
	fmt.Printf("mul indexes: %v\n", mulIndexes)
	mulInfo := mulReg.FindAllStringSubmatch(input, -1)
	for _, m := range doReg.FindAllStringSubmatchIndex(input, -1) {
		doIndexes = append(doIndexes, m[2])
	}
	fmt.Printf("do indexes: %v\n", doIndexes)
	for _, m := range dontReg.FindAllStringSubmatchIndex(input, -1) {
		dontIndexes = append(dontIndexes, m[2])
	}
	fmt.Printf("dont indexes: %v\n", dontIndexes)

	var total int
	enabled := true
	var mulI, doI, dontI int
	for i := 0; i < len(input); i++ {
		if doI < len(doIndexes) && doIndexes[doI] == i {
			enabled = true
			doI++
		}
		if dontI < len(dontIndexes) && dontIndexes[dontI] == i {
			enabled = false
			dontI++
		}

		if mulI < len(mulIndexes) && mulIndexes[mulI] == i {
			if enabled {
				curMul := mulInfo[mulI]
				total += getMulTotal(curMul)
			}
			mulI++
		}
	}

	// Get list of do and don'ts
	return total
}
