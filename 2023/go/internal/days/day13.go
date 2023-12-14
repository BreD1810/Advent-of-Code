package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
)

func Day13() {
	fileContents := util.ReadFileLines("inputs/day13-actual.txt")
	fmt.Printf("Part 1: %d\n", day13Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day13Part2(fileContents))
}

type mirrorPatternInfo struct {
	rocks map[util.Coordinate]struct{}
	maxX  int
	maxY  int
}

func day13Part1(fileContents []string) int {
	patterns := parseMirrorPatterns(fileContents)
	score := 0
	for _, p := range patterns {
		verticalScore, horizontalScore := findMirrorScore(p)
		score += verticalScore + (100 * horizontalScore)
	}
	return score
}

func day13Part2(fileContents []string) int {
	patterns := parseMirrorPatterns(fileContents)
	score := 0
	for _, p := range patterns {
		verticalScore, horizontalScore := findMirrorScoreWithSmudges(p)
		score += verticalScore + (100 * horizontalScore)
	}
	return score
}

func parseMirrorPatterns(input []string) []mirrorPatternInfo {
	curRocks := make(map[util.Coordinate]struct{})
	i, maxX := 0, 0
	var patterns []mirrorPatternInfo

	for _, l := range input {
		if l == "" {
			patterns = append(patterns, mirrorPatternInfo{rocks: curRocks, maxX: maxX, maxY: i - 1})
			curRocks = make(map[util.Coordinate]struct{})
			i, maxX = 0, 0
			continue
		}
		for x, c := range l {
			if c == '#' {
				curRocks[util.Coordinate{X: x, Y: i}] = struct{}{}
			}
		}
		i++
		maxX = len(l) - 1
	}

	patterns = append(patterns, mirrorPatternInfo{rocks: curRocks, maxX: maxX, maxY: i - 1})

	return patterns
}

func findMirrorScore(pattern mirrorPatternInfo) (verticalScore int, horizontalScore int) {
	for x := 1; x <= pattern.maxX; x++ {
		columnsToCheck := min(x, pattern.maxX-x+1)
		isMirror := true
	mirrorCheck:
		for i := 0; i <= columnsToCheck; i++ {
			for r, _ := range pattern.rocks {
				if r.X < x && r.X >= x-columnsToCheck {
					toCheck := util.Coordinate{X: x + util.IntAbs(x-r.X) - 1, Y: r.Y}
					if _, ok := pattern.rocks[toCheck]; !ok {
						isMirror = false
						break mirrorCheck
					}
				}
				if r.X >= x && r.X < x+columnsToCheck {
					toCheck := util.Coordinate{X: x - util.IntAbs(x-r.X) - 1, Y: r.Y}
					if _, ok := pattern.rocks[toCheck]; !ok {
						isMirror = false
						break mirrorCheck
					}
				}
			}
		}
		if isMirror {
			return x, 0
		}
	}

	for y := 1; y <= pattern.maxY; y++ {
		rowsToCheck := min(y, pattern.maxY-y+1)
		isMirror := true
	mirrorCheck2:
		for i := 0; i <= rowsToCheck; i++ {
			for r, _ := range pattern.rocks {
				if r.Y < y && r.Y >= y-rowsToCheck {
					toCheck := util.Coordinate{X: r.X, Y: y + util.IntAbs(y-r.Y) - 1}
					if _, ok := pattern.rocks[toCheck]; !ok {
						isMirror = false
						break mirrorCheck2
					}
				}
				if r.Y >= y && r.Y < y+rowsToCheck {
					toCheck := util.Coordinate{X: r.X, Y: y - util.IntAbs(y-r.Y) - 1}
					if _, ok := pattern.rocks[toCheck]; !ok {
						isMirror = false
						break mirrorCheck2
					}
				}
			}
		}
		if isMirror {
			return 0, y
		}
	}

	log.Fatalf("couldn't find mirror in pattern %v\n", pattern)
	return 0, 0
}

func findMirrorScoreWithSmudges(pattern mirrorPatternInfo) (verticalScore int, horizontalScore int) {
	for x := 1; x <= pattern.maxX; x++ {
		columnsToCheck := min(x, pattern.maxX-x+1)
		isMirror := true
		var missingRock *util.Coordinate
	mirrorCheck:
		for i := 0; i <= columnsToCheck; i++ {
			for r, _ := range pattern.rocks {
				if r.X < x && r.X >= x-columnsToCheck {
					toCheck := util.Coordinate{X: x + util.IntAbs(x-r.X) - 1, Y: r.Y}
					if _, ok := pattern.rocks[toCheck]; !ok {
						if missingRock != nil && (missingRock.X != toCheck.X || missingRock.Y != toCheck.Y) {
							isMirror = false
							break mirrorCheck
						}
						missingRock = &toCheck
					}
				}
				if r.X >= x && r.X < x+columnsToCheck {
					toCheck := util.Coordinate{X: x - util.IntAbs(x-r.X) - 1, Y: r.Y}
					if _, ok := pattern.rocks[toCheck]; !ok {
						if missingRock != nil && (missingRock.X != toCheck.X || missingRock.Y != toCheck.Y) {
							isMirror = false
							break mirrorCheck
						}
						missingRock = &toCheck
					}
				}
			}
		}
		if isMirror && missingRock != nil {
			return x, 0
		}
	}

	for y := 1; y <= pattern.maxY; y++ {
		rowsToCheck := min(y, pattern.maxY-y+1)
		isMirror := true
		var missingRock *util.Coordinate
	mirrorCheck2:
		for i := 0; i <= rowsToCheck; i++ {
			for r, _ := range pattern.rocks {
				if r.Y < y && r.Y >= y-rowsToCheck {
					toCheck := util.Coordinate{X: r.X, Y: y + util.IntAbs(y-r.Y) - 1}
					if _, ok := pattern.rocks[toCheck]; !ok {
						if missingRock != nil && (missingRock.X != toCheck.X || missingRock.Y != toCheck.Y) {
							isMirror = false
							break mirrorCheck2
						}
						missingRock = &toCheck
					}
				}
				if r.Y >= y && r.Y < y+rowsToCheck {
					toCheck := util.Coordinate{X: r.X, Y: y - util.IntAbs(y-r.Y) - 1}
					if _, ok := pattern.rocks[toCheck]; !ok {
						if missingRock != nil && (missingRock.X != toCheck.X || missingRock.Y != toCheck.Y) {
							isMirror = false
							break mirrorCheck2
						}
						missingRock = &toCheck
					}
				}
			}
		}
		if isMirror && missingRock != nil {
			return 0, y
		}
	}

	log.Fatalf("couldn't find mirror in pattern %v\n", pattern)
	return 0, 0
}
