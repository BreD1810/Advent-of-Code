package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"
)

func Day18() {
	fileContents := util.ReadFileLines("inputs/day18-1.txt")
	fmt.Printf("Part 1: %d\n", day18Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day18Part2(fileContents))
}

func day18Part1(fileContents []string) int {
	cubeMap := parseCubes(fileContents)
	var totalSA int

	for cube := range cubeMap {
		var cubesSA int
		for _, n := range cubeNeighbours(cube) {
			if _, found := cubeMap[n]; !found {
				cubesSA++
			}
		}
		totalSA += cubesSA
	}

	return totalSA
}

func parseCubes(fileContents []string) map[util.Coordinate3D]struct{} {
	cubeMap := make(map[util.Coordinate3D]struct{})

	for _, l := range fileContents {
		splits := strings.Split(l, ",")
		x := util.GetIntFromString(splits[0])
		y := util.GetIntFromString(splits[1])
		z := util.GetIntFromString(splits[2])
		cubeMap[util.Coordinate3D{X: x, Y: y, Z: z}] = struct{}{}
	}

	return cubeMap
}

func cubeNeighbours(c util.Coordinate3D) []util.Coordinate3D {
	return []util.Coordinate3D{
		{X: c.X - 1, Y: c.Y, Z: c.Z},
		{X: c.X + 1, Y: c.Y, Z: c.Z},
		{X: c.X, Y: c.Y - 1, Z: c.Z},
		{X: c.X, Y: c.Y + 1, Z: c.Z},
		{X: c.X, Y: c.Y, Z: c.Z - 1},
		{X: c.X, Y: c.Y, Z: c.Z + 1},
	}
}

func day18Part2(fileContents []string) int {
	cubes := parseCubes(fileContents)

	lavaBounds := util.Coordinate3D{X: 0, Y: 0, Z: 0}

	for cube := range cubes {
		if lavaBounds.X < cube.X {
			lavaBounds.X = cube.X
		}
		if lavaBounds.Y < cube.Y {
			lavaBounds.Y = cube.Y
		}
		if lavaBounds.Z < cube.Z {
			lavaBounds.Z = cube.Z
		}
	}

	startingCube := util.Coordinate3D{X: 0, Y: 0, Z: 0}
	exterior := make(map[util.Coordinate3D]struct{})
	return lavaExteriorSA(startingCube, cubes, exterior, lavaBounds)
}

// Traverse around the outside of the cube
func lavaExteriorSA(curCube util.Coordinate3D, cubes map[util.Coordinate3D]struct{}, exterior map[util.Coordinate3D]struct{}, lavaBounds util.Coordinate3D) int {
	if _, isExterior := exterior[curCube]; isExterior {
		return 0
	}
	if curCube.X < -1 || curCube.X > lavaBounds.X+1 ||
		curCube.Y < -1 || curCube.Y > lavaBounds.Y+1 ||
		curCube.Z < -1 || curCube.Z > lavaBounds.Z+1 {
		return 0
	}
	if _, isInScan := cubes[curCube]; isInScan {
		return 1
	}

	exterior[curCube] = struct{}{}
	var count int
	for _, n := range cubeNeighbours(curCube) {
		count += lavaExteriorSA(n, cubes, exterior, lavaBounds)
	}
	return count
}
