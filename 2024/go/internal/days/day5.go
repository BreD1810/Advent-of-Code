package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	fileContents := util.ReadFileLines("inputs/day5.txt")
	fmt.Printf("Part 1: %d\n", day5Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day5Part2(fileContents))
}

func day5Part1(input []string) int {
	rules := make(map[int][]int)
	var total int
	var startPrinting bool

	for _, l := range input {
		if l == "" {
			startPrinting = true
			continue
		}

		if !startPrinting {
			ss := strings.Split(l, "|")
			if len(ss) != 2 {
				fmt.Printf("incorrect rule format: %s", l)
				panic("incorrect rule format")
			}
			p1, err := strconv.Atoi(ss[0])
			if err != nil {
				panic(err)
			}
			p2, err := strconv.Atoi(ss[1])
			if err != nil {
				panic(err)
			}
			curRules, ok := rules[p1]
			if !ok {
				rules[p1] = []int{p2}
				continue
			}
			rules[p1] = append(curRules, p2)
			continue
		}

		splits := strings.Split(l, ",")
		isValid := true
	endline:
		for n, curPageStr := range splits {
			curPage, err := strconv.Atoi(curPageStr)
			if err != nil {
				panic(err)
			}
			for i := n; i < len(splits); i++ {
				page, err := strconv.Atoi(splits[i])
				if err != nil {
					panic(err)
				}
				if slices.Contains(rules[page], curPage) {
					isValid = false
					break endline
				}
			}
		}
		if isValid {
			middle := splits[(len(splits)-1)/2]
			middleInt, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}
			total += middleInt
		}
	}

	return total
}
func day5Part2(input []string) int {
	rules := make(map[int][]int)
	var startPrinting bool
	var invalidLines [][]string

	for _, l := range input {
		if l == "" {
			startPrinting = true
			continue
		}

		if !startPrinting {
			ss := strings.Split(l, "|")
			if len(ss) != 2 {
				fmt.Printf("incorrect rule format: %s", l)
				panic("incorrect rule format")
			}
			p1, err := strconv.Atoi(ss[0])
			if err != nil {
				panic(err)
			}
			p2, err := strconv.Atoi(ss[1])
			if err != nil {
				panic(err)
			}
			curRules, ok := rules[p1]
			if !ok {
				rules[p1] = []int{p2}
				continue
			}
			rules[p1] = append(curRules, p2)
			continue
		}

		splits := strings.Split(l, ",")
	endline:
		for n, curPageStr := range splits {
			curPage, err := strconv.Atoi(curPageStr)
			if err != nil {
				panic(err)
			}
			for i := n; i < len(splits); i++ {
				page, err := strconv.Atoi(splits[i])
				if err != nil {
					panic(err)
				}
				if slices.Contains(rules[page], curPage) {
					invalidLines = append(invalidLines, splits)
					break endline
				}
			}
		}

	}

	var total int
	for _, splits := range invalidLines {
		var curInts []int
		for _, s := range splits {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			curInts = append(curInts, i)
		}
		fmt.Printf("sorting %v\n", curInts)

		var sorted bool
		var n int
		for !sorted {
			curPage := curInts[n]
			var broken bool
			for i := n + 1; i < len(curInts); i++ {
				page := curInts[i]
				if slices.Contains(rules[page], curPage) {
					curInts[n] = page
					curInts[i] = curPage
					n = 0
					broken = true
					break
				}
			}
			if !broken {
				n++
			}
			if n == len(curInts) {
				sorted = true
			}
		}

		fmt.Println(curInts)
		middle := curInts[(len(curInts)-1)/2]
		fmt.Println(middle)
		total += middle
	}
	return total

}
