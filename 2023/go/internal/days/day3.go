package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"regexp"
)

type partInfo struct {
	xMin  int
	xMax  int
	value int
}

func Day3() {
	fileContents := util.ReadFileLines("inputs/day3-actual.txt")
	fmt.Printf("Part 1: %d\n", day3Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day3Part2(fileContents))
}

func day3Part1(fileContents []string) int {
	symbolRegex := regexp.MustCompile(`[^\w.]`)
	numberRegex := regexp.MustCompile(`\d+`)
	parts := make(map[int][]partInfo)
	var symbols []util.Coordinate

	for yIndex, l := range fileContents {
		symbolMatches := symbolRegex.FindAllStringIndex(l, -1)
		for _, match := range symbolMatches {
			symbols = append(symbols, util.Coordinate{X: match[0], Y: yIndex})
		}

		parts[yIndex] = parseParts(numberRegex, l)
	}

	total := 0
	for _, symbol := range symbols {
		aboveSizes := getAdjacentPartSizes(parts[symbol.Y-1], symbol.X)
		for _, size := range aboveSizes {
			total += size
		}

		equalSizes := getAdjacentPartSizes(parts[symbol.Y], symbol.X)
		for _, size := range equalSizes {
			total += size
		}

		belowSizes := getAdjacentPartSizes(parts[symbol.Y+1], symbol.X)
		for _, size := range belowSizes {
			total += size
		}
	}

	return total
}

func parseParts(numberRegex *regexp.Regexp, l string) []partInfo {
	var parts []partInfo
	numberMatches := numberRegex.FindAllStringIndex(l, -1)

	for _, match := range numberMatches {
		start, end := match[0], match[1]
		v := l[start:end]
		pInfo := partInfo{
			xMin:  start,
			xMax:  end - 1,
			value: util.GetIntFromString(v),
		}
		parts = append(parts, pInfo)
	}
	return parts
}

func getAdjacentPartSizes(parts []partInfo, symbolX int) []int {
	var partSizes []int
	for _, part := range parts {
		if part.xMin >= symbolX-1 && part.xMin <= symbolX+1 {
			partSizes = append(partSizes, part.value)
			continue
		}
		if part.xMax >= symbolX-1 && part.xMax <= symbolX+1 {
			partSizes = append(partSizes, part.value)
			continue
		}
		if part.xMin < symbolX && part.xMax > symbolX {
			partSizes = append(partSizes, part.value)
		}
	}
	return partSizes
}

func day3Part2(fileContents []string) int {
	gearRegex := regexp.MustCompile(`\*`)
	numberRegex := regexp.MustCompile(`\d+`)
	parts := make(map[int][]partInfo)
	var gears []util.Coordinate

	for yIndex, l := range fileContents {
		gearMatches := gearRegex.FindAllStringIndex(l, -1)
		for _, match := range gearMatches {
			gears = append(gears, util.Coordinate{X: match[0], Y: yIndex})
		}

		parts[yIndex] = parseParts(numberRegex, l)
	}

	total := 0
	for _, gear := range gears {
		var partSizes []int
		partSizes = append(partSizes, getAdjacentPartSizes(parts[gear.Y-1], gear.X)...)
		partSizes = append(partSizes, getAdjacentPartSizes(parts[gear.Y], gear.X)...)
		partSizes = append(partSizes, getAdjacentPartSizes(parts[gear.Y+1], gear.X)...)
		if len(partSizes) >= 2 {
			gearTotal := 0
			for i, partSize := range partSizes {
				if i == 0 {
					gearTotal = partSize
				} else {
					gearTotal *= partSize
				}
			}
			total += gearTotal
		}
	}

	return total
}
