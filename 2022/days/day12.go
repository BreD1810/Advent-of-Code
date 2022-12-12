package days

import (
	"aoc-2022/util"
	"fmt"
	"runtime"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

func Day12() {
	fileContents := util.ReadFileLines("inputs/day12-1.txt")

	fmt.Printf("Part 1: %d\n", day12Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day12Part2Pooling(fileContents))
}

type heightmap struct {
	data  [][]rune
	start util.Coordinate
	dest  util.Coordinate
}

func day12Part1(fileContents []string) int {
	hm := parseHeightmap(fileContents)
	// printHeightmap(hm)
	return runDijkstras(hm)
}

func parseHeightmap(fileContents []string) heightmap {
	hmData := make([][]rune, len(fileContents))
	start := util.Coordinate{}
	dest := util.Coordinate{}

	for i, l := range fileContents {
		hmData[i] = []rune(l)
		if n := strings.IndexRune(l, 'S'); n >= 0 {
			start.X = n
			start.Y = i
			hmData[i][n] = 'a'
		}

		if n := strings.IndexRune(l, 'E'); n >= 0 {
			dest.X = n
			dest.Y = i
			hmData[i][n] = 'z'
		}
	}

	return heightmap{data: hmData, start: start, dest: dest}
}

func coordinateToInt(c util.Coordinate) int {
	return c.X*1000 + c.Y
}

func runDijkstras(hm heightmap) int {
	g := dijkstra.NewGraph()

	for y, row := range hm.data {
		for x := range row {
			g.AddVertex(coordinateToInt(util.Coordinate{X: x, Y: y}))
		}
	}

	for y, row := range hm.data {
		for x, h := range row {
			c := util.Coordinate{X: x, Y: y}
			cInt := coordinateToInt(c)

			ns := findNeighbours(c, len(row)-1, len(hm.data)-1)
			for _, n := range ns {
				if hm.data[n.Y][n.X] <= h+1 {
					g.AddArc(cInt, coordinateToInt(n), 1)
				}
			}
		}
	}

	// fmt.Printf("Starting at %d, ending at %d\n", coordinateToInt(hm.start), coordinateToInt(hm.dest))

	res, err := g.Shortest(coordinateToInt(hm.start), coordinateToInt(hm.dest))
	if err != nil {
		// fmt.Println("No path found")
		return util.MaxInt
	}
	return int(res.Distance)
}

func findNeighbours(c util.Coordinate, maxX int, maxY int) []util.Coordinate {
	ns := make([]util.Coordinate, 0)

	// Up
	if c.Y > 0 {
		ns = append(ns, util.Coordinate{X: c.X, Y: c.Y - 1})
	}
	// Down
	if c.Y < maxY {
		ns = append(ns, util.Coordinate{X: c.X, Y: c.Y + 1})
	}
	// Left
	if c.X > 0 {
		ns = append(ns, util.Coordinate{X: c.X - 1, Y: c.Y})
	}
	// Right
	if c.X < maxX {
		ns = append(ns, util.Coordinate{X: c.X + 1, Y: c.Y})
	}

	return ns
}

func printHeightmap(hm heightmap) {
	for i, row := range hm.data {
		for j := range row {
			fmt.Printf("%c", hm.data[i][j])
		}
		fmt.Println()
	}
}
func day12Part2(fileContents []string) int {
	return day12Part2FanOut(fileContents)
}

func day12Part2Pooling(fileContents []string) int {
	hm := parseHeightmap(fileContents)

	aLocs := getALocations(hm)
	workCh := make(chan util.Coordinate, len(aLocs))
	resCh := make(chan int, len(aLocs))
	minSteps := util.MaxInt

	noCpus := runtime.NumCPU()
	for t := 0; t < noCpus; t++ {
		go func() {
			for l := range workCh {
				newHm := heightmap{data: hm.data, start: l, dest: hm.dest}
				resCh <- runDijkstras(newHm)
			}
		}()
	}

	for _, loc := range aLocs {
		workCh <- loc
	}

	for i := 0; i < len(aLocs); i++ {
		d := <-resCh
		if d < minSteps {
			minSteps = d
		}
	}

	return minSteps
}

func day12Part2FanOut(fileContents []string) int {
	hm := parseHeightmap(fileContents)

	aLocs := getALocations(hm)
	ch := make(chan int, len(aLocs))
	minSteps := util.MaxInt

	for _, loc := range aLocs {
		go func(loc util.Coordinate) {
			newHm := heightmap{data: hm.data, start: loc, dest: hm.dest}
			ch <- runDijkstras(newHm)
		}(loc)
	}

	for i := 0; i < len(aLocs); i++ {
		d := <-ch
		if d < minSteps {
			minSteps = d
		}
	}

	return minSteps
}

func day12Part2Single(fileContents []string) int {
	hm := parseHeightmap(fileContents)

	aLocs := getALocations(hm)
	// fmt.Printf("Number of a: %d\n", len(aLocs))
	minSteps := util.MaxInt

	for _, loc := range aLocs {
		newHm := heightmap{data: hm.data, start: loc, dest: hm.dest}
		d := runDijkstras(newHm)
		if d < minSteps {
			minSteps = d
		}
	}

	return minSteps
}

func getALocations(hm heightmap) []util.Coordinate {
	aLocs := make([]util.Coordinate, 0)

	for y, row := range hm.data {
		for x, h := range row {
			if h == 'a' {
				ns := findNeighbours(util.Coordinate{X: x, Y: y}, len(row)-1, len(hm.data)-1)
				for _, n := range ns {
					if hm.data[n.Y][n.X] == 'b' {
						aLocs = append(aLocs, util.Coordinate{X: x, Y: y})
						continue
					}
				}
			}
		}
	}

	return aLocs
}
