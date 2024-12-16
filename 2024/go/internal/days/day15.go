package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day15() {
	fileContents := util.ReadFileLines("inputs/day15.txt")
	fmt.Printf("Part 1: %d\n", day15Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day15Part2(fileContents))
}

func day15Part1(input []string) int {
	walls := make(map[util.Coordinate]struct{})
	boxes := make(map[util.Coordinate]struct{})
	var moves []util.Movement

	var maxX, maxY int
	var startLoc util.Coordinate
	for y, l := range input {
		if y == 1 {
			maxX = len(l)
		}
		if l == "" {
			maxY = y - 1
			continue
		}

		for x, r := range l {
			switch r {
			case '#':
				walls[util.Coordinate{X: x, Y: y}] = struct{}{}
			case '.':
				continue
			case 'O':
				boxes[util.Coordinate{X: x, Y: y}] = struct{}{}
			case '@':
				startLoc = util.Coordinate{X: x, Y: y}
			case '<':
				moves = append(moves, util.MoveLeft)
			case '^':
				moves = append(moves, util.MoveUp)
			case '>':
				moves = append(moves, util.MoveRight)
			case 'v':
				moves = append(moves, util.MoveDown)
			default:
				panic("couldn't parse line")
			}
		}
	}

	for _, m := range moves {
		moveLanternFish(&startLoc, walls, boxes, m, maxX, maxY)
		// printLanternFish(&startLoc, walls, boxes, m, maxX, maxY)
	}

	var total int
	for b := range boxes {
		total += (100 * b.Y) + b.X
	}

	return total
}

func printLanternFish(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			cur := util.Coordinate{X: x, Y: y}
			if _, ok := walls[cur]; ok {
				fmt.Printf("#")
				continue
			}
			if _, ok := boxes[cur]; ok {
				fmt.Printf("O")
				continue
			}
			if *pos == cur {
				fmt.Printf("@")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
}

func moveLanternFish(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, move util.Movement, maxX, maxY int) {
	var checkPos util.Coordinate
	// fmt.Println(move)
	switch move {
	case util.MoveUp:
		checkPos = util.Coordinate{X: pos.X, Y: pos.Y - 1}
	case util.MoveDown:
		checkPos = util.Coordinate{X: pos.X, Y: pos.Y + 1}
	case util.MoveLeft:
		checkPos = util.Coordinate{X: pos.X - 1, Y: pos.Y}
	case util.MoveRight:
		checkPos = util.Coordinate{X: pos.X + 1, Y: pos.Y}
	}

	// Cannot move if it's a wall next
	if _, ok := walls[checkPos]; ok {
		return
	}
	// If it's not a box next, move into the position
	if _, ok := boxes[checkPos]; !ok {
		*pos = checkPos
		return
	}
	// Box chain
	switch move {
	case util.MoveRight:
		for i := pos.X + 2; i < maxX; i++ {
			// If there's a wall at the end of the chain, no move
			next := util.Coordinate{X: i, Y: pos.Y}
			if _, ok := walls[next]; ok {
				return
			}
			// If gap, move
			if _, ok := boxes[next]; !ok {
				delete(boxes, checkPos)
				*pos = checkPos
				boxes[next] = struct{}{}
				return
			}
		}
	case util.MoveLeft:
		for i := pos.X - 2; i >= 0; i-- {
			// If there's a wall at the end of the chain, no move
			next := util.Coordinate{X: i, Y: pos.Y}
			if _, ok := walls[next]; ok {
				return
			}
			// If gap, move
			if _, ok := boxes[next]; !ok {
				delete(boxes, checkPos)
				*pos = checkPos
				boxes[next] = struct{}{}
				return
			}
		}
	case util.MoveUp:
		for i := pos.Y - 2; i >= 0; i-- {
			// If there's a wall at the end of the chain, no move
			next := util.Coordinate{X: pos.X, Y: i}
			if _, ok := walls[next]; ok {
				return
			}
			// If gap, move
			if _, ok := boxes[next]; !ok {
				delete(boxes, checkPos)
				*pos = checkPos
				boxes[next] = struct{}{}
				return
			}
		}
	case util.MoveDown:
		for i := pos.Y + 2; i < maxY; i++ {
			// If there's a wall at the end of the chain, no move
			next := util.Coordinate{X: pos.X, Y: i}
			if _, ok := walls[next]; ok {
				return
			}
			// If gap, move
			if _, ok := boxes[next]; !ok {
				delete(boxes, checkPos)
				*pos = checkPos
				boxes[next] = struct{}{}
				return
			}
		}
	}
}

func day15Part2(input []string) int {
	walls := make(map[util.Coordinate]struct{})
	boxes := make(map[util.Coordinate]struct{})
	var moves []util.Movement

	var maxX, maxY int
	var startLoc util.Coordinate
	for y, l := range input {
		if y == 1 {
			maxX = len(l) * 2
		}
		if l == "" {
			maxY = y
			continue
		}

		for x, r := range l {
			switch r {
			case '#':
				walls[util.Coordinate{X: (x * 2), Y: y}] = struct{}{}
				walls[util.Coordinate{X: (x * 2) + 1, Y: y}] = struct{}{}
			case '.':
				continue
			case 'O':
				boxes[util.Coordinate{X: (x * 2), Y: y}] = struct{}{}
			case '@':
				startLoc = util.Coordinate{X: x * 2, Y: y}
				fmt.Println(startLoc)
			case '<':
				moves = append(moves, util.MoveLeft)
			case '^':
				moves = append(moves, util.MoveUp)
			case '>':
				moves = append(moves, util.MoveRight)
			case 'v':
				moves = append(moves, util.MoveDown)
			default:
				panic("couldn't parse line")
			}
		}
	}
	printLanternFish2(&startLoc, walls, boxes, maxX, maxY)

	for _, m := range moves {
		moveLanternFish2(&startLoc, walls, boxes, m, maxX, maxY)
		// printLanternFish2(&startLoc, walls, boxes, maxX, maxY)
	}

	var total int
	for b := range boxes {
		total += (100 * b.Y) + b.X
	}

	return total
}

func printLanternFish2(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			cur := util.Coordinate{X: x, Y: y}
			if _, ok := walls[cur]; ok {
				fmt.Printf("#")
				continue
			}
			if _, ok := boxes[cur]; ok {
				fmt.Printf("[]")
				x++
				continue
			}
			if *pos == cur {
				fmt.Printf("@")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
}

func moveLanternFish2(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, move util.Movement, maxX, maxY int) {
	switch move {
	case util.MoveUp:
		moveLanternFishUp(pos, walls, boxes, maxY)
	case util.MoveDown:
		moveLanternFishDown(pos, walls, boxes, maxY)
	case util.MoveLeft:
		moveLanternFishLeft(pos, walls, boxes, maxX)
	case util.MoveRight:
		moveLanternFishRight(pos, walls, boxes, maxX)
	}
}

func moveLanternFishRight(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxX int) {
	checkPos := util.Coordinate{X: pos.X + 1, Y: pos.Y}
	// Cannot move if it's a wall
	if _, ok := walls[checkPos]; ok {
		return
	}

	// move into empty
	if _, ok := boxes[checkPos]; !ok {
		*pos = checkPos
		return
	}

	toMove := []util.Coordinate{checkPos}
	for i := pos.X + 3; i < maxX; i += 2 {
		// If there's a wall at the end of the chain, no move
		next := util.Coordinate{X: i, Y: pos.Y}
		if _, ok := walls[next]; ok {
			return
		}
		// If gap, move
		if _, ok := boxes[next]; !ok {
			*pos = checkPos
			for _, b := range toMove {
				delete(boxes, b)
			}
			for _, b := range toMove {
				boxes[util.Coordinate{X: b.X + 1, Y: b.Y}] = struct{}{}
			}
			return
		}
		toMove = append(toMove, next)
	}
}

func moveLanternFishLeft(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxX int) {
	checkPos := util.Coordinate{X: pos.X - 1, Y: pos.Y}
	// Cannot move if it's a wall
	if _, ok := walls[checkPos]; ok {
		return
	}
	// move into empty
	checkPos2 := util.Coordinate{X: pos.X - 2, Y: pos.Y}
	if _, ok := boxes[checkPos2]; !ok {
		*pos = checkPos
		return
	}

	toMove := []util.Coordinate{checkPos2}
	for i := pos.X - 4; i < maxX; i -= 2 {
		// If there's a wall at the end of the chain, no move
		next := util.Coordinate{X: i, Y: pos.Y}
		nextRight := util.Coordinate{X: i + 1, Y: pos.Y}
		if _, ok := walls[nextRight]; ok {
			return
		}
		// If gap, move
		if _, ok := boxes[next]; !ok {
			*pos = checkPos
			for _, b := range toMove {
				delete(boxes, b)
			}
			for _, b := range toMove {
				boxes[util.Coordinate{X: b.X - 1, Y: b.Y}] = struct{}{}
			}
			return
		}
		toMove = append(toMove, next)
	}
}

func moveLanternFishUp(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxY int) {
	checkPos := util.Coordinate{X: pos.X, Y: pos.Y - 1}
	// Cannot move if it's a wall
	if _, ok := walls[checkPos]; ok {
		return
	}
	// move into empty, checking directly above and up left
	_, okBox1 := boxes[checkPos]
	checkPos2 := util.Coordinate{X: pos.X - 1, Y: pos.Y - 1}
	_, okBox2 := boxes[checkPos2]
	if !okBox1 && !okBox2 {
		*pos = util.Coordinate{X: pos.X, Y: pos.Y - 1}
		return
	}

	var toMove []util.Coordinate
	var lastBoxes []util.Coordinate
	if okBox1 {
		lastBoxes = append(lastBoxes, checkPos)
		toMove = append(toMove, checkPos)
	}
	if okBox2 {
		lastBoxes = append(lastBoxes, checkPos2)
		toMove = append(toMove, checkPos2)
	}

	for i := pos.Y - 2; i < maxY; i-- {
		var nextLastBoxes []util.Coordinate
		for _, b := range lastBoxes {
			// Skip if there's a wall in the way
			if _, ok := walls[util.Coordinate{X: b.X, Y: i}]; ok {
				return
			}
			if _, ok := walls[util.Coordinate{X: b.X + 1, Y: i}]; ok {
				return
			}

			box1Loc := util.Coordinate{X: b.X, Y: i}
			if _, ok := boxes[box1Loc]; ok {
				toMove = append(toMove, box1Loc)
				nextLastBoxes = append(nextLastBoxes, box1Loc)
			}

			box2Loc := util.Coordinate{X: b.X - 1, Y: i}
			if _, ok := boxes[box2Loc]; ok {
				toMove = append(toMove, box2Loc)
				nextLastBoxes = append(nextLastBoxes, box2Loc)
			}

			box3Loc := util.Coordinate{X: b.X + 1, Y: i}
			if _, ok := boxes[box3Loc]; ok {
				toMove = append(toMove, box3Loc)
				nextLastBoxes = append(nextLastBoxes, box3Loc)
			}
		}
		// If there's all spaces
		if len(nextLastBoxes) == 0 {
			*pos = util.Coordinate{X: pos.X, Y: pos.Y - 1}
			for _, b := range toMove {
				delete(boxes, b)
			}
			for _, b := range toMove {
				boxes[util.Coordinate{X: b.X, Y: b.Y - 1}] = struct{}{}
			}

			return
		}
		lastBoxes = nextLastBoxes
	}
}

func moveLanternFishDown(pos *util.Coordinate, walls map[util.Coordinate]struct{}, boxes map[util.Coordinate]struct{}, maxY int) {
	checkPos := util.Coordinate{X: pos.X, Y: pos.Y + 1}
	// Cannot move if it's a wall
	if _, ok := walls[checkPos]; ok {
		return
	}
	// move into empty, checking directly above and up left
	_, okBox1 := boxes[checkPos]
	checkPos2 := util.Coordinate{X: pos.X - 1, Y: pos.Y + 1}
	_, okBox2 := boxes[checkPos2]
	if !okBox1 && !okBox2 {
		*pos = util.Coordinate{X: pos.X, Y: pos.Y + 1}
		return
	}

	var toMove []util.Coordinate
	var lastBoxes []util.Coordinate
	if okBox1 {
		lastBoxes = append(lastBoxes, checkPos)
		toMove = append(toMove, checkPos)
	}
	if okBox2 {
		lastBoxes = append(lastBoxes, checkPos2)
		toMove = append(toMove, checkPos2)
	}

	for i := pos.Y + 2; i < maxY; i++ {
		var nextLastBoxes []util.Coordinate
		for _, b := range lastBoxes {
			// Skip if there's a wall in the way
			if _, ok := walls[util.Coordinate{X: b.X, Y: i}]; ok {
				return
			}
			if _, ok := walls[util.Coordinate{X: b.X + 1, Y: i}]; ok {
				return
			}

			box1Loc := util.Coordinate{X: b.X, Y: i}
			if _, ok := boxes[box1Loc]; ok {
				toMove = append(toMove, box1Loc)
				nextLastBoxes = append(nextLastBoxes, box1Loc)
			}

			box2Loc := util.Coordinate{X: b.X - 1, Y: i}
			if _, ok := boxes[box2Loc]; ok {
				toMove = append(toMove, box2Loc)
				nextLastBoxes = append(nextLastBoxes, box2Loc)
			}

			box3Loc := util.Coordinate{X: b.X + 1, Y: i}
			if _, ok := boxes[box3Loc]; ok {
				toMove = append(toMove, box3Loc)
				nextLastBoxes = append(nextLastBoxes, box3Loc)
			}
		}
		// If there's all spaces
		if len(nextLastBoxes) == 0 {
			*pos = util.Coordinate{X: pos.X, Y: pos.Y + 1}
			for _, b := range toMove {
				delete(boxes, b)
			}
			for _, b := range toMove {
				boxes[util.Coordinate{X: b.X, Y: b.Y + 1}] = struct{}{}
			}
			return
		}
		lastBoxes = nextLastBoxes
	}
}
