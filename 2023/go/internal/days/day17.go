package days

import (
	"aoc-2023/internal/util"
	"container/heap"
	"fmt"
	"log"
)

type movement int

const (
	moveUp movement = iota
	moveRight
	moveDown
	moveLeft
)

func Day17() {
	fileContents := util.ReadFileLines("inputs/day17-example.txt")
	fmt.Printf("Part 1: %d\n", day17Part1(fileContents))
}

func day17Part1(fileContents []string) int {
	grid := parseHeatLossGrid(fileContents)
	return getMinimalHeatLoss(grid, 0, 3)
}

func day17Part2(fileContents []string) int {
	grid := parseHeatLossGrid(fileContents)
	return getMinimalHeatLoss(grid, 4, 10)
}

func parseHeatLossGrid(inp []string) [][]int {
	out := make([][]int, len(inp))
	for y, row := range inp {
		out[y] = make([]int, len(row))
		for x, c := range row {
			out[y][x] = util.GetIntFromString(string(c))
		}
	}
	return out
}

type crucibleValue struct {
	location      util.Coordinate
	heatLoss      int
	straightCount int
	direction     movement
}

type cruciblePriorityQueue []*crucibleValue

func (pq cruciblePriorityQueue) Len() int { return len(pq) }

func (pq cruciblePriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq cruciblePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *cruciblePriorityQueue) Push(x any) {
	v := x.(*crucibleValue)
	*pq = append(*pq, v)
}

func (pq *cruciblePriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type crucibleCacheEntry struct {
	location      util.Coordinate
	straightCount int
	direction     movement
}

func getMinimalHeatLoss(grid [][]int, minBeforeTurning int, maxBeforeTurning int) int {
	straightMove := crucibleValue{
		location:      util.Coordinate{X: 1, Y: 0},
		heatLoss:      0,
		straightCount: 1,
		direction:     moveRight,
	}
	downMove := crucibleValue{
		location:      util.Coordinate{X: 0, Y: 1},
		heatLoss:      0,
		straightCount: 1,
		direction:     moveDown,
	}

	q := cruciblePriorityQueue{&straightMove, &downMove}
	heap.Init(&q)
	cache := make(map[crucibleCacheEntry]int)

	for q.Len() > 0 {
		curValue := heap.Pop(&q).(*crucibleValue)

		if curValue.location.X < 0 || curValue.location.X >= len(grid[0]) || curValue.location.Y < 0 || curValue.location.Y >= len(grid) {
			continue
		}

		curHeatLoss := grid[curValue.location.Y][curValue.location.X] + curValue.heatLoss
		if curValue.location.X == len(grid[0])-1 && curValue.location.Y == len(grid)-1 {
			return curHeatLoss
		}

		cacheVal := crucibleCacheEntry{location: curValue.location, straightCount: curValue.straightCount, direction: curValue.direction}
		if cachedLoss, ok := cache[cacheVal]; ok {
			if cachedLoss <= curHeatLoss {
				continue
			}
		}
		cache[cacheVal] = curHeatLoss

		if curValue.straightCount < maxBeforeTurning { // Max in a row is 3
			straightLoc, straightDirection := getNextCrucibleLocationAndDirection(curValue.location, curValue.direction, "straight")
			heap.Push(&q, &crucibleValue{
				location:      straightLoc,
				direction:     straightDirection,
				heatLoss:      curHeatLoss,
				straightCount: curValue.straightCount + 1,
			})
		}
		if curValue.straightCount >= minBeforeTurning {
			rightLoc, rightDirection := getNextCrucibleLocationAndDirection(curValue.location, curValue.direction, "right")
			heap.Push(&q, &crucibleValue{
				location:      rightLoc,
				direction:     rightDirection,
				heatLoss:      curHeatLoss,
				straightCount: 1,
			})
			leftLoc, leftDirection := getNextCrucibleLocationAndDirection(curValue.location, curValue.direction, "left")
			heap.Push(&q, &crucibleValue{
				location:      leftLoc,
				direction:     leftDirection,
				heatLoss:      curHeatLoss,
				straightCount: 1,
			})
		}
	}

	log.Fatalln("couldn't get to the end of the grid")
	return 0
}

func getNextCrucibleLocationAndDirection(loc util.Coordinate, curDirection movement, ins string) (util.Coordinate, movement) {
	switch curDirection {
	case moveUp:
		if ins == "right" {
			return util.Coordinate{X: loc.X + 1, Y: loc.Y}, moveRight
		}
		if ins == "left" {
			return util.Coordinate{X: loc.X - 1, Y: loc.Y}, moveLeft
		}
		if ins == "straight" {
			return util.Coordinate{X: loc.X, Y: loc.Y - 1}, moveUp
		}
	case moveRight:
		if ins == "right" {
			return util.Coordinate{X: loc.X, Y: loc.Y + 1}, moveDown
		}
		if ins == "left" {
			return util.Coordinate{X: loc.X, Y: loc.Y - 1}, moveUp
		}
		if ins == "straight" {
			return util.Coordinate{X: loc.X + 1, Y: loc.Y}, moveRight
		}
	case moveDown:
		if ins == "right" {
			return util.Coordinate{X: loc.X - 1, Y: loc.Y}, moveLeft
		}
		if ins == "left" {
			return util.Coordinate{X: loc.X + 1, Y: loc.Y}, moveRight
		}
		if ins == "straight" {
			return util.Coordinate{X: loc.X, Y: loc.Y + 1}, moveDown
		}
	case moveLeft:
		if ins == "right" {
			return util.Coordinate{X: loc.X, Y: loc.Y - 1}, moveUp
		}
		if ins == "left" {
			return util.Coordinate{X: loc.X, Y: loc.Y + 1}, moveDown
		}
		if ins == "straight" {
			return util.Coordinate{X: loc.X - 1, Y: loc.Y}, moveLeft
		}
	}

	log.Fatalf("couldn't figure out how to move from %v in direction %d with instruction %s\n", loc, curDirection, ins)
	return util.Coordinate{}, moveUp
}
