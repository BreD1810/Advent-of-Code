package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"strings"
)

type sourceDestinationLength struct {
	destination int
	source      int
	length      int
}

func Day5() {
	fileContents := util.ReadFileLines("inputs/day5-actual.txt")
	fmt.Printf("Part 1: %d\n", day5Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day5Part2(fileContents))
}

func day5Part1(fileContents []string) int {
	seeds := parseSeeds(fileContents[0])
	listOfMaps := getListOfMaps(fileContents)

	lowest := util.MaxInt
	for _, s := range seeds {
		s = mapSeed(listOfMaps, s)
		if s < lowest {
			lowest = s
		}
	}
	return lowest
}

func parseSeeds(s string) []int {
	var seeds []int
	s = strings.TrimPrefix(s, "seeds: ")
	ss := strings.Split(s, " ")
	for _, v := range ss {
		seeds = append(seeds, util.GetIntFromString(v))
	}
	return seeds
}

func getListOfMaps(fileContents []string) [][]sourceDestinationLength {
	listOfMaps := make([][]sourceDestinationLength, 7)
	for i := 0; i < len(listOfMaps); i++ {
		listOfMaps[i] = []sourceDestinationLength{}
	}

	mapCounter := 0

	for i, l := range fileContents[3:] {
		if strings.ContainsRune(l, '-') || i == len(fileContents)-1 {
			continue
		}
		if l == "" {
			mapCounter++
			continue
		}

		ss := strings.Split(l, " ")
		destinationStart := util.GetIntFromString(ss[0])
		sourceStart := util.GetIntFromString(ss[1])
		length := util.GetIntFromString(ss[2])
		curMap := sourceDestinationLength{
			destination: destinationStart,
			source:      sourceStart,
			length:      length,
		}

		listOfMaps[mapCounter] = append(listOfMaps[mapCounter], curMap)
	}

	return listOfMaps
}

func mapSeed(listOfMaps [][]sourceDestinationLength, s int) int {
	for i := 0; i < len(listOfMaps); i++ {
		curMaps := listOfMaps[i]
		for _, m := range curMaps {
			if s >= m.source && s < m.source+m.length {
				s = m.destination + (s - m.source)
				break
			}
		}
	}
	return s
}

func day5Part2(fileContents []string) int {
	seeds := parseSeeds(fileContents[0])
	listOfMaps := getListOfMaps(fileContents)
	resCh := make(chan int, len(seeds)/2)

	for i := 0; i < len(seeds); i += 2 {
		go func(i int) {
			lowest := util.MaxInt
			for j := 0; j < seeds[i+1]; j++ {
				loc := mapSeed(listOfMaps, seeds[i]+j)
				if loc < lowest {
					lowest = loc
				}
			}
			resCh <- lowest
		}(i)
	}

	lowest := util.MaxInt
	for i := 0; i < len(seeds)/2; i++ {
		res := <-resCh
		if res < lowest {
			lowest = res
		}
	}

	return lowest
}
