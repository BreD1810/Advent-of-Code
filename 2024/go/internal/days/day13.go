package days

import (
	"aoc-2024/internal/util"
	"fmt"
	"strconv"
)

func Day13() {
	fileContents := util.ReadFileLines("inputs/day8.txt")
	fmt.Printf("Part 1: %d\n", day13Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day13Part2(fileContents))
}

type buttonProblem struct {
	a     util.Coordinate
	b     util.Coordinate
	prize util.Coordinate
}

func day13Part1(input []string) int {
	var problems []buttonProblem
	for i := 0; i < len(input); i += 4 {
		bp := parseButtonProblem(input[i : i+3])
		problems = append(problems, bp)
	}

	fmt.Println(problems)

	var total int
	for i, p := range problems {
		fmt.Println("problem " + strconv.Itoa(i))

		det := p.a.X*p.b.Y - p.b.X*p.a.Y
		detX := p.prize.X*p.b.Y - p.b.X*p.prize.Y
		detY := p.a.X*p.prize.Y - p.prize.X*p.a.Y
		if det != 0 && detX == (detX/det)*det && detY == (detY/det)*det {
			total += (detX/det)*3 + (detY / det)
		}
	}

	return total
}

func parseButtonProblem(input []string) buttonProblem {
	var aX, aY, bX, bY, prizeX, prizeY int
	fmt.Sscanf(input[0], "Button A: X+%d, Y+%d", &aX, &aY)
	fmt.Sscanf(input[1], "Button B: X+%d, Y+%d", &bX, &bY)
	fmt.Sscanf(input[2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)

	return buttonProblem{
		a:     util.Coordinate{X: aX, Y: aY},
		b:     util.Coordinate{X: bX, Y: bY},
		prize: util.Coordinate{X: prizeX, Y: prizeY},
	}
}

func day13Part2(input []string) int {
	var problems []buttonProblem
	for i := 0; i < len(input); i += 4 {
		bp := parseButtonProblem(input[i : i+3])
		problems = append(problems, bp)
	}

	fmt.Println(problems)

	var total int
	for i, p := range problems {
		fmt.Println("problem " + strconv.Itoa(i))
		p.prize.X += 10000000000000
		p.prize.Y += 10000000000000

		det := p.a.X*p.b.Y - p.b.X*p.a.Y
		detX := p.prize.X*p.b.Y - p.b.X*p.prize.Y
		detY := p.a.X*p.prize.Y - p.prize.X*p.a.Y
		if det != 0 && detX == (detX/det)*det && detY == (detY/det)*det {
			total += (detX/det)*3 + (detY / det)
		}
	}

	return total
}
