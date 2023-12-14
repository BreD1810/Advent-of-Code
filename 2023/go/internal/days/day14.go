package days

import (
	"aoc-2023/internal/util"
	"fmt"
)

type rollingRockInfo struct {
	stationaryRocks map[util.Coordinate]struct{}
	rollingRocks    map[util.Coordinate]struct{}
	maxX            int
	maxY            int
}

func Day14() {
	fileContents := util.ReadFileLines("inputs/day14-example.txt")
	fmt.Printf("Part 1: %d\n", day14Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day14Part2(fileContents))
}

func day14Part1(fileContents []string) int {
	rockInfo := parseRollingRocks(fileContents)
	rollRocksNorth(&rockInfo)
	return getTotalLoad(rockInfo)
}

func day14Part2(fileContents []string) int {
	rockInfo := parseRollingRocks(fileContents)
	rollCycleCache := map[string]int{}
	billion := 1_000_000_000

	for cycle := 1; cycle <= billion; cycle++ {
		rollRocksNorth(&rockInfo)
		rollRocksWest(&rockInfo)
		rollRocksSouth(&rockInfo)
		rollRocksEast(&rockInfo)

		s := fmt.Sprintf("%+v", rockInfo.rollingRocks)

		if cycleStart, ok := rollCycleCache[s]; ok {
			cycleLength := cycle - cycleStart
			cycle = billion - ((billion - cycle) % cycleLength)
			continue
		}
		rollCycleCache[s] = cycle
	}
	return getTotalLoad(rockInfo)
}

func parseRollingRocks(input []string) rollingRockInfo {
	stationaryRocks := make(map[util.Coordinate]struct{})
	rollingRocks := make(map[util.Coordinate]struct{})

	for y, l := range input {
		for x, c := range l {
			if c == '#' {
				stationaryRocks[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
			if c == 'O' {
				rollingRocks[util.Coordinate{X: x, Y: y}] = struct{}{}
			}
		}
	}

	return rollingRockInfo{stationaryRocks: stationaryRocks, rollingRocks: rollingRocks, maxX: len(input[0]), maxY: len(input)}
}

func rollRocksNorth(rockInfo *rollingRockInfo) {
	newRollingRocks := make(map[util.Coordinate]struct{})
	for x := 0; x < rockInfo.maxX; x++ {
		nextRockY := 0
		for y := 0; y < rockInfo.maxY; y++ {
			if _, isStationaryRock := rockInfo.stationaryRocks[util.Coordinate{X: x, Y: y}]; isStationaryRock {
				nextRockY = y + 1
			}
			if _, isRollingRock := rockInfo.rollingRocks[util.Coordinate{X: x, Y: y}]; isRollingRock {
				newRollingRocks[util.Coordinate{X: x, Y: nextRockY}] = struct{}{}
				nextRockY += 1
			}
		}
	}
	rockInfo.rollingRocks = newRollingRocks
}

func rollRocksSouth(rockInfo *rollingRockInfo) {
	newRollingRocks := make(map[util.Coordinate]struct{})
	for x := 0; x < rockInfo.maxX; x++ {
		nextRockY := rockInfo.maxY - 1
		for y := rockInfo.maxY - 1; y >= 0; y-- {
			if _, isStationaryRock := rockInfo.stationaryRocks[util.Coordinate{X: x, Y: y}]; isStationaryRock {
				nextRockY = y - 1
			}
			if _, isRollingRock := rockInfo.rollingRocks[util.Coordinate{X: x, Y: y}]; isRollingRock {
				newRollingRocks[util.Coordinate{X: x, Y: nextRockY}] = struct{}{}
				nextRockY -= 1
			}
		}
	}
	rockInfo.rollingRocks = newRollingRocks
}

func rollRocksWest(rockInfo *rollingRockInfo) {
	newRollingRocks := make(map[util.Coordinate]struct{})
	for y := 0; y < rockInfo.maxY; y++ {
		nextRockX := 0
		for x := 0; x < rockInfo.maxX; x++ {
			if _, isStationaryRock := rockInfo.stationaryRocks[util.Coordinate{X: x, Y: y}]; isStationaryRock {
				nextRockX = x + 1
			}
			if _, isRollingRock := rockInfo.rollingRocks[util.Coordinate{X: x, Y: y}]; isRollingRock {
				newRollingRocks[util.Coordinate{X: nextRockX, Y: y}] = struct{}{}
				nextRockX += 1
			}
		}
	}
	rockInfo.rollingRocks = newRollingRocks
}

func rollRocksEast(rockInfo *rollingRockInfo) {
	newRollingRocks := make(map[util.Coordinate]struct{})
	for y := 0; y < rockInfo.maxY; y++ {
		nextRockX := rockInfo.maxX - 1
		for x := rockInfo.maxX - 1; x >= 0; x-- {
			if _, isStationaryRock := rockInfo.stationaryRocks[util.Coordinate{X: x, Y: y}]; isStationaryRock {
				nextRockX = x - 1
			}
			if _, isRollingRock := rockInfo.rollingRocks[util.Coordinate{X: x, Y: y}]; isRollingRock {
				newRollingRocks[util.Coordinate{X: nextRockX, Y: y}] = struct{}{}
				nextRockX -= 1
			}
		}
	}
	rockInfo.rollingRocks = newRollingRocks
}

func getTotalLoad(rockInfo rollingRockInfo) int {
	total := 0
	for rLoc := range rockInfo.rollingRocks {
		total += rockInfo.maxY - rLoc.Y
	}
	return total
}
