package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"strconv"
	"strings"
)

func Day11() {
	fileContents := util.ReadFileLine("inputs/day11.txt")
	fmt.Printf("Part 1: %d\n", day11Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day11Part2(fileContents))
}

func day11Part1(input string) int {
	stones := strings.Split(input, " ")
	for i := 0; i < 25; i++ {
		var newStones []string
		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
				continue
			}
			stoneLen := len(stone)
			if stoneLen%2 == 0 {
				stone1 := stone[0 : stoneLen/2]
				stone1 = strings.TrimLeft(stone1, "0")
				if stone1 == "" {
					stone1 = "0"
				}
				newStones = append(newStones, stone1)

				stone2 := stone[stoneLen/2:]
				stone2 = strings.TrimLeft(stone2, "0")
				if stone2 == "" {
					stone2 = "0"
				}
				newStones = append(newStones, stone2)
				continue
			}
			stoneval, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			stoneval *= 2024
			newStones = append(newStones, strconv.Itoa(stoneval))
		}
		stones = newStones
	}
	return len(stones)
}

// stoneCache stores stone value to value produced
type stoneIter struct {
	stoneVal  int
	iteration int
}

var stoneCache = make(map[stoneIter]int)

func day11Part2(input string) int {
	stones := strings.Split(input, " ")
	var res int
	for _, s := range stones {
		stone, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res += stoneBlink(stone, 75)
	}

	return res
}

func stoneBlink(val int, iteration int) int {
	if iteration == 0 {
		return 1
	}

	if cval, ok := stoneCache[stoneIter{stoneVal: val, iteration: iteration}]; ok {
		return cval
	}

	if val == 0 {
		res := stoneBlink(1, iteration-1)
		stoneCache[stoneIter{stoneVal: val, iteration: iteration}] = res
		return res
	}

	strVal := strconv.Itoa(val)
	if len(strVal)%2 == 0 {
		stone1, err := strconv.Atoi(strVal[:len(strVal)/2])
		if err != nil {
			panic(err)
		}
		stone2, err := strconv.Atoi(strVal[len(strVal)/2:])
		if err != nil {
			panic(err)
		}
		res := stoneBlink(stone1, iteration-1) + stoneBlink(stone2, iteration-1)
		stoneCache[stoneIter{stoneVal: val, iteration: iteration}] = res
		return res
	}

	res := stoneBlink(val*2024, iteration-1)
	stoneCache[stoneIter{stoneVal: val, iteration: iteration}] = res
	return res
}
