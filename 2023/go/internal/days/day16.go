package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
)

func Day16() {
	fileContents := util.ReadFileLines("inputs/day16-example.txt")
	fmt.Printf("Part 1: %d\n", day16Part1(fileContents))
	//fmt.Printf("Part 2: %d\n", day16Part2(fileContents))
}

type direction util.Coordinate

var (
	north = direction{X: 0, Y: -1}
	east  = direction{X: 1, Y: 0}
	south = direction{X: 0, Y: 1}
	west  = direction{X: -1, Y: 0}
)

type energisedStart struct {
	dir      direction
	location util.Coordinate
}

func day16Part1(fileContents []string) int {
	energisedTiles := make(map[util.Coordinate][]direction)
	energiseGrid(fileContents, east, util.Coordinate{X: 0, Y: 0}, energisedTiles)
	return len(energisedTiles)
}

func day16Part2(fileContents []string) int {
	starts := getEnergisedTilesStarts(fileContents)
	maxEnergised := 0
	for _, startPos := range starts {
		energisedTiles := make(map[util.Coordinate][]direction)
		energiseGrid(fileContents, startPos.dir, startPos.location, energisedTiles)
		energisedScore := len(energisedTiles)
		if energisedScore > maxEnergised {
			maxEnergised = energisedScore
		}
	}
	return maxEnergised
}

func getEnergisedTilesStarts(inp []string) []energisedStart {
	var starts []energisedStart
	for y, row := range inp {
		if y == 0 {
			for x := 0; x < len(row); x++ {
				starts = append(starts, energisedStart{dir: south, location: util.Coordinate{X: x, Y: 0}})
			}
		}
		if y == len(inp)-1 {
			for x := 0; x < len(row); x++ {
				starts = append(starts, energisedStart{dir: north, location: util.Coordinate{X: x, Y: y}})
			}
		}
		starts = append(starts, energisedStart{dir: east, location: util.Coordinate{X: 0, Y: y}})
		starts = append(starts, energisedStart{dir: west, location: util.Coordinate{X: len(row) - 1, Y: y}})
	}
	return starts
}

func energiseGrid(grid []string, dir direction, curLoc util.Coordinate, energisedTiles map[util.Coordinate][]direction) {
	if curLoc.X < 0 || curLoc.X >= len(grid[0]) || curLoc.Y < 0 || curLoc.Y >= len(grid) {
		return
	}

	if visited, ok := energisedTiles[curLoc]; ok {
		for _, v := range visited {
			if v == dir {
				return
			}
		}
	}

	tileType := grid[curLoc.Y][curLoc.X]
	energisedTiles[curLoc] = append(energisedTiles[curLoc], dir)
	nextDirections := getNextDirections(dir, string(tileType))

	for _, nextDir := range nextDirections {
		energiseGrid(grid, nextDir, util.Coordinate{X: curLoc.X + nextDir.X, Y: curLoc.Y + nextDir.Y}, energisedTiles)
	}
}

func getNextDirections(dir direction, tileType string) []direction {
	switch dir {
	case north:
		switch tileType {
		case "|":
			return []direction{north}
		case "-":
			return []direction{east, west}
		case "\\":
			return []direction{west}
		case "/":
			return []direction{east}
		case ".":
			return []direction{north}
		}
	case east:
		switch tileType {
		case "|":
			return []direction{north, south}
		case "-":
			return []direction{east}
		case "\\":
			return []direction{south}
		case "/":
			return []direction{north}
		case ".":
			return []direction{east}
		}
	case south:
		switch tileType {
		case "|":
			return []direction{south}
		case "-":
			return []direction{east, west}
		case "\\":
			return []direction{east}
		case "/":
			return []direction{west}
		case ".":
			return []direction{south}
		}

	case west:
		switch tileType {
		case "|":
			return []direction{north, south}
		case "-":
			return []direction{west}
		case "\\":
			return []direction{north}
		case "/":
			return []direction{south}
		case ".":
			return []direction{west}
		}
	}

	log.Fatalf("couldn't get next direction for direction %d and type %s", dir, tileType)
	return nil
}
