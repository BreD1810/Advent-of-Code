package days

import (
	"aoc-2022/util"
	"fmt"
	"log"
	"strings"
)

func Day15() {
	fileContents := util.ReadFileLines("inputs/day015-1.txt")
	fmt.Printf("Part 1: %d\n", day15Part1(fileContents, 2000000))
	fmt.Printf("Part 2: %d\n", day15Part2(fileContents, 4000000))
}

type sensorBeaconPair struct {
	sensorLoc util.Coordinate
	beaconLoc util.Coordinate
}

func day15Part1(fileContents []string, rowToCheck int) int {
	pairs := parseSensorBeacons(fileContents)
	return countBeaconlessLocationsAtRow(pairs, rowToCheck)
}

func parseSensorBeacons(fileContents []string) (pairs []sensorBeaconPair) {
	for _, l := range fileContents {
		equalsSplits := strings.Split(l, "=")

		sensorX := util.GetIntFromString(strings.Split(equalsSplits[1], ",")[0])
		sensorY := util.GetIntFromString(strings.Split(equalsSplits[2], ":")[0])
		beaconX := util.GetIntFromString(strings.Split(equalsSplits[3], ",")[0])
		beaconY := util.GetIntFromString(equalsSplits[4])

		pairs = append(pairs, sensorBeaconPair{
			sensorLoc: util.Coordinate{X: sensorX, Y: sensorY},
			beaconLoc: util.Coordinate{X: beaconX, Y: beaconY},
		})
	}
	return
}

func countBeaconlessLocationsAtRow(pairs []sensorBeaconPair, row int) int {
	sensorsAndBeacons := make(map[int]struct{})
	coveredSpaces := make(map[int]struct{})

	for _, p := range pairs {
		if p.sensorLoc.Y == row {
			sensorsAndBeacons[p.sensorLoc.X] = struct{}{}
		}
		if p.beaconLoc.Y == row {
			sensorsAndBeacons[p.beaconLoc.X] = struct{}{}
		}

		manhatten := getManhatten(p.sensorLoc, p.beaconLoc)
		yDistanceToRow := util.IntAbs(row - p.sensorLoc.Y)
		radiusAtRow := manhatten - yDistanceToRow
		if radiusAtRow >= 0 {
			minX := p.sensorLoc.X - int(radiusAtRow)
			maxX := p.sensorLoc.X + int(radiusAtRow)
			for x := minX; x <= maxX; x++ {
				coveredSpaces[x] = struct{}{}
			}
		}
	}

	return len(coveredSpaces) - len(sensorsAndBeacons)
}

func day15Part2(fileContents []string, largestXY int) int {
	pairs := parseSensorBeacons(fileContents)
	fmt.Println(pairs)
	unseenBeaconLoc := getUnseenBeacon(pairs, largestXY)
	return unseenBeaconLoc.X*4000000 + unseenBeaconLoc.Y
}

func getUnseenBeacon(pairs []sensorBeaconPair, largestXY int) util.Coordinate {
	for y := 0; y <= largestXY; y++ {
		for x := 0; x <= largestXY; x++ {
			inside := false
			for _, p := range pairs {
				sBManhatten := getManhatten(p.sensorLoc, p.beaconLoc)
				sPManhatten := getManhatten(p.sensorLoc, util.Coordinate{X: x, Y: y})
				if sPManhatten <= sBManhatten {
					inside = true
					x = p.sensorLoc.X + sBManhatten - util.IntAbs(p.sensorLoc.Y-y)
					break
				}
			}
			if !inside {
				return util.Coordinate{X: x, Y: y}
			}
		}
	}
	log.Fatalln("Unable to find hidden beacon")
	return util.Coordinate{X: 0, Y: 0}
}

func getManhatten(a, b util.Coordinate) int {
	return util.IntAbs(a.X-b.X) + util.IntAbs(a.Y-b.Y)
}
