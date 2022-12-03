package days

import (
	"aoc-2022/util"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func Day3() {
	fileContents := util.ReadFileLines("inputs/day03-1.txt")
	fmt.Printf("Part 1: %d\n", day3Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day3Part2(fileContents))
}

func day3Part1(fileContents []string) int {
	commonItems := make([]rune, len(fileContents))

	for i, l := range fileContents {
		midpoint := len(l) / 2
		s1 := l[:midpoint]
		s2 := l[midpoint:]
		commonItems[i] = getCommonRune(s1, s2)
	}

	return sumItems(commonItems)
}

func getCommonRune(s1, s2 string) rune {
	// fmt.Printf("S1: %s, S2: %s\n", s1, s2)
	for _, r := range s1 {
		if strings.ContainsRune(s2, r) {
			// fmt.Printf("Common rune: %q\n", r)
			return r
		}
	}
	log.Fatalf("No common runes between %s and %s", s1, s2)
	return 'a'
}

func sumItems(items []rune) int {
	var sum int
	for _, r := range items {
		// fmt.Printf("Rune is %q", r)
		if unicode.IsUpper(r) {
			// fmt.Printf(", value is %d\n", int(r)-38)
			sum += int(r) - 38
		} else {
			// fmt.Printf(", value is %d\n", int(r)-96)
			sum += int(r) - 96
		}
	}
	return sum
}

func day3Part2(fileContents []string) int {
	var badges []rune

	for i := 0; i < len(fileContents); i += 3 {
		s1 := fileContents[i]
		s2 := fileContents[i+1]
		s3 := fileContents[i+2]
		b := getBadge(s1, s2, s3)
		badges = append(badges, b)
	}

	return sumItems(badges)
}

func getBadge(s1, s2, s3 string) rune {
	for _, r := range s1 {
		if strings.ContainsRune(s2, r) && strings.ContainsRune(s3, r) {
			return r
		}
	}
	log.Fatalf("No common runes between %s and %s", s1, s2)
	return 'a'
}
