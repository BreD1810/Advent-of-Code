package days

import (
	"aoc-2021/util"
	"fmt"
)

type binaryCount struct {
	noZeros int
	noOnes  int
}

func Day3() {
	fileContents := util.ReadFileLines("inputs/day03-1.txt")
	fmt.Printf("Part 1: %d\n", day3Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day3Part2(fileContents))
}

func day3Part1(vals []string) int {
	noDigits := len(vals[0])
	binaryCounts := make([]binaryCount, noDigits)

	for i := 0; i < noDigits; i++ {
		binaryCounts[i] = binaryCount{0, 0}
	}

	for _, val := range vals {
		for i := 0; i < noDigits; i++ {
			switch val[i] {
			case '1':
				binaryCounts[i].noOnes++
			case '0':
				binaryCounts[i].noZeros++
			default:
				fmt.Println("Error: Character is not 0 or 1")
			}
		}
	}

	gammaRate := binaryCountsToGamma(binaryCounts)
	epsilonRate := binaryCountsToEpsilon(binaryCounts)
	return gammaRate * epsilonRate
}

func day3Part2(vals []string) int64 {
	generatorRating := getGeneratorRating(vals)
	scrubberRating := getScrubberRating(vals)

	return generatorRating * scrubberRating
}

func binaryCountsToGamma(counts []binaryCount) int {
	binaryPower := 0
	total := 0
	for i := len(counts) - 1; i >= 0; i-- {
		if counts[i].noOnes > counts[i].noZeros {
			total += util.PowWithInts(2, binaryPower)
		}
		binaryPower += 1
	}

	return total
}

func binaryCountsToEpsilon(counts []binaryCount) int {
	binaryPower := 0
	total := 0

	for i := len(counts) - 1; i >= 0; i-- {
		if counts[i].noOnes < counts[i].noZeros {
			total += util.PowWithInts(2, binaryPower)
		}
		binaryPower += 1
	}

	return total
}

func getGeneratorRating(vals []string) int64 {
	count := 0
	for len(vals) > 1 && count < len(vals[0]) {
		binaryCount := binaryCount{0, 0}
		for _, val := range vals {
			switch val[count] {
			case '1':
				binaryCount.noOnes++
			case '0':
				binaryCount.noZeros++
			default:
				fmt.Println("Error: Character is not 0 or 1")
			}
		}

		if binaryCount.noOnes >= binaryCount.noZeros {
			vals = filterVals(vals, count, '1')
		} else {
			vals = filterVals(vals, count, '0')
		}

		count++
	}

	if len(vals) != 1 {
		fmt.Println("Not just 1 left for generator rating")
	}

	return util.BinaryStringToDecimal(vals[0])
}

func getScrubberRating(vals []string) int64 {
	count := 0
	for len(vals) > 1 && count < len(vals[0]) {
		binaryCount := binaryCount{0, 0}
		for _, val := range vals {
			switch val[count] {
			case '1':
				binaryCount.noOnes++
			case '0':
				binaryCount.noZeros++
			default:
				fmt.Println("Error: Character is not 0 or 1")
			}
		}

		if binaryCount.noOnes < binaryCount.noZeros {
			vals = filterVals(vals, count, '1')
		} else {
			vals = filterVals(vals, count, '0')
		}

		count++
	}

	if len(vals) != 1 {
		fmt.Println("Not just 1 left for scrubber rating")
	}

	return util.BinaryStringToDecimal(vals[0])
}

func filterVals(vals []string, index int, value byte) []string {
	var res []string
	for _, val := range vals {
		if val[index] == value {
			res = append(res, val)
		}
	}
	return res
}
