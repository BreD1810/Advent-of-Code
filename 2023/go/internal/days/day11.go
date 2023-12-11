package days

import (
	"aoc-2023/internal/util"
	"fmt"
)

func Day11() {
	fileContents := util.ReadFileLines("inputs/day11-actual.txt")
	fmt.Printf("Part 1: %d\n", day11Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day11Part2(fileContents))
}

func day11Part1(fileContents []string) int {
	galaxies := getGalaxies(fileContents)
	expandedGalaxies := expandGalaxies(fileContents, galaxies, 1)

	total := 0
	for i, curG := range expandedGalaxies {
		for j, g := range expandedGalaxies {
			if i == j {
				continue
			}
			total += util.IntAbs(curG.X-g.X) + util.IntAbs(curG.Y-g.Y)
		}
	}

	return total / 2
}

func day11Part2(fileContents []string) int {
	galaxies := getGalaxies(fileContents)
	expandedGalaxies := expandGalaxies(fileContents, galaxies, 1_000_000-1)

	total := 0
	for i, curG := range expandedGalaxies {
		for j, g := range expandedGalaxies {
			if i == j {
				continue
			}
			total += util.IntAbs(curG.X-g.X) + util.IntAbs(curG.Y-g.Y)
		}
	}

	return total / 2
}

func getGalaxies(input []string) []util.Coordinate {
	var galaxies []util.Coordinate

	for y, l := range input {
		for x, c := range l {
			if c == '#' {
				galaxies = append(galaxies, util.Coordinate{X: x, Y: y})
			}
		}
	}

	return galaxies
}

func expandGalaxies(space []string, galaxies []util.Coordinate, expansionFactor int) []util.Coordinate {
	newGalaxies := make([]util.Coordinate, len(galaxies))
	copy(newGalaxies, galaxies)

	// Horizontal
	for y, l := range space {
		containsGalaxy := false
		for _, c := range l {
			if c == '#' {
				containsGalaxy = true
				break
			}
		}
		if !containsGalaxy {
			for n, g := range galaxies {
				if g.Y > y {
					newGalaxies[n].Y += expansionFactor
				}
			}
		}
	}

	// Vertical
	for x := 0; x < len(space[0]); x++ {
		containsGalaxy := false
		for _, l := range space {
			if l[x] == '#' {
				containsGalaxy = true
				break
			}
		}
		if !containsGalaxy {
			for n, g := range galaxies {
				if g.X > x {
					newGalaxies[n].X += expansionFactor
				}
			}
		}
	}

	return newGalaxies
}
