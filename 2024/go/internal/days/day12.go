package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day12() {
	fileContents := util.ReadFile2DRune("inputs/day12.txt")
	fmt.Printf("Part 1: %d\n", day12Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day12Part2(fileContents))
}

func day12Part1(input [][]rune) int {
	var total int
	visited := make(map[util.Coordinate]struct{})

	for y, row := range input {
		for x, val := range row {
			coord := util.Coordinate{X: x, Y: y}
			if _, ok := visited[coord]; ok {
				continue
			}

			area, perimeter := calcPlantAreaPerimeter(val, input, coord, visited)
			total += area * perimeter
		}
	}

	return total
}

// calcPlantAreaPerimeter returns the area and perimeter
func calcPlantAreaPerimeter(plant rune, input [][]rune, loc util.Coordinate, visited map[util.Coordinate]struct{}) (int, int) {
	if loc.X < 0 || loc.Y < 0 || loc.X >= len(input[0]) || loc.Y >= len(input) || input[loc.Y][loc.X] != plant {
		return 0, 1
	}

	if _, ok := visited[loc]; ok {
		if input[loc.Y][loc.X] == plant {
			return 0, 0
		}
		return 0, 1
	}

	visited[loc] = struct{}{}

	area := 1
	var perimeter int
	upArea, upPeri := calcPlantAreaPerimeter(plant, input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, visited)
	downArea, downPeri := calcPlantAreaPerimeter(plant, input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, visited)
	leftArea, leftPeri := calcPlantAreaPerimeter(plant, input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, visited)
	rightArea, rightPeri := calcPlantAreaPerimeter(plant, input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, visited)
	area += upArea + downArea + leftArea + rightArea
	perimeter += upPeri + downPeri + leftPeri + rightPeri

	return area, perimeter
}

func day12Part2(input [][]rune) int {
	var total int
	visited := make(map[util.Coordinate]struct{})

	for y, row := range input {
		for x, val := range row {
			coord := util.Coordinate{X: x, Y: y}
			if _, ok := visited[coord]; ok {
				continue
			}

			area, sides := calcPlantAreaSides(val, input, coord, visited)

			total += area * sides
		}
	}

	return total
}

// calcPlantAreaSides returns the area and number of sides
func calcPlantAreaSides(plant rune, input [][]rune, loc util.Coordinate, visited map[util.Coordinate]struct{}) (int, int) {
	newVisited := make(map[util.Coordinate]struct{})
	calcPlantLocations(plant, input, loc, visited, newVisited)
	corners := getCorners(len(input[0])-1, len(input), newVisited)
	return len(newVisited), corners
}

func calcPlantLocations(plant rune, input [][]rune, loc util.Coordinate, visited map[util.Coordinate]struct{}, newVisited map[util.Coordinate]struct{}) {
	if loc.X < 0 || loc.Y < 0 || loc.X >= len(input[0]) || loc.Y >= len(input) || input[loc.Y][loc.X] != plant {
		return
	}

	if _, ok := visited[loc]; ok {
		return

	}

	visited[loc] = struct{}{}
	newVisited[loc] = struct{}{}

	calcPlantLocations(plant, input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, visited, newVisited)
	calcPlantLocations(plant, input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, visited, newVisited)
	calcPlantLocations(plant, input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, visited, newVisited)
	calcPlantLocations(plant, input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, visited, newVisited)
}

func getCorners(maxX, maxY int, locs map[util.Coordinate]struct{}) int {
	var total int
	for loc := range locs {
		_, up := locs[util.Coordinate{X: loc.X, Y: loc.Y - 1}]
		_, down := locs[util.Coordinate{X: loc.X, Y: loc.Y + 1}]
		_, left := locs[util.Coordinate{X: loc.X - 1, Y: loc.Y}]
		_, right := locs[util.Coordinate{X: loc.X + 1, Y: loc.Y}]
		_, upLeft := locs[util.Coordinate{X: loc.X - 1, Y: loc.Y - 1}]
		_, upRight := locs[util.Coordinate{X: loc.X + 1, Y: loc.Y - 1}]
		_, downLeft := locs[util.Coordinate{X: loc.X - 1, Y: loc.Y + 1}]
		_, downRight := locs[util.Coordinate{X: loc.X + 1, Y: loc.Y + 1}]

		// Top left outer corner
		if !up && !left {
			total++
		}
		// Top right outer corner
		if !up && !right {
			total++
		}
		// Bottom left outer corner
		if !down && !left {
			total++
		}
		// Bottom right outer corner
		if !down && !right {
			total++
		}

		// Top left inner corner (J)
		if up && left && !upLeft {
			total++
		}
		// Top Right inner corner (L)
		if up && right && !upRight {
			total++
		}
		// Bottom Left inner corner
		if down && left && !downLeft {
			total++
		}
		// Bottom Right inner corner
		if down && right && !downRight {
			total++
		}
	}
	return total
}

// calcPlantAreaSides returns the area and number of sides
// func calcPlantAreaSides(plant rune, input [][]rune, loc util.Coordinate, visited map[util.Coordinate]struct{}) (int, int) {
// 	area := 1
// 	sides := 4
// 	visited[loc] = struct{}{}
// 	upArea, upSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, visited, util.MoveUp)
// 	downArea, downSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, visited, util.MoveDown)
// 	leftArea, leftSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, visited, util.MoveLeft)
// 	rightArea, rightSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, visited, util.MoveRight)
// 	return area + upArea + downArea + leftArea + rightArea, sides + upSides + downSides + leftSides + rightSides
// }

// calcPlantAreaSidesHelper returns the area and number of sides
// func calcPlantAreaSidesHelper(plant rune, input [][]rune, loc util.Coordinate, visited map[util.Coordinate]struct{}, dir util.Movement) (int, int) {
// 	if loc.X < 0 || loc.Y < 0 || loc.X >= len(input[0]) || loc.Y >= len(input) || input[loc.Y][loc.X] != plant {
// 		return 0, 0
// 	}

// 	if _, ok := visited[loc]; ok {
// 		return 0, 0
// 	}

// 	visited[loc] = struct{}{}

// 	area := 1
// 	var sides int
// 	upArea, upSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X, Y: loc.Y - 1}, visited, util.MoveUp)
// 	downArea, downSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X, Y: loc.Y + 1}, visited, util.MoveDown)
// 	leftArea, leftSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X - 1, Y: loc.Y}, visited, util.MoveLeft)
// 	rightArea, rightSides := calcPlantAreaSidesHelper(plant, input, util.Coordinate{X: loc.X + 1, Y: loc.Y}, visited, util.MoveRight)
// 	area += upArea + downArea + leftArea + rightArea
// 	sides += upSides + downSides + leftSides + rightSides
// 	if upSides == sides || downSides == sides || leftSides == sides || rightSides == sides {
// 		return area, sides + 1
// 	}

// 	return area, sides
// }
