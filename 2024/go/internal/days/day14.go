package days

import (
	"aoc-2024/internal/util"
	"fmt"
)

func Day14() {
	fileContents := util.ReadFileLines("inputs/day14.txt")
	fmt.Printf("Part 1: %d\n", day14Part1(fileContents, 101, 103))
	fmt.Printf("Part 2: %d\n", day14Part2(fileContents, 101, 103))
}

type robotInfo struct {
	pos util.Coordinate
	vel util.Coordinate
}

func day14Part1(input []string, xMax, yMax int) int {
	robots := make([]*robotInfo, len(input))
	for i, l := range input {
		var xPos, yPos, xVel, yVel int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &xPos, &yPos, &xVel, &yVel)
		robots[i] = &robotInfo{
			pos: util.Coordinate{X: xPos, Y: yPos},
			vel: util.Coordinate{X: xVel, Y: yVel},
		}
	}

	for i := 0; i < 100; i++ {
		for _, r := range robots {
			moveRobot(r, xMax, yMax)
		}
	}

	xMid := (xMax - 1) / 2
	yMid := (yMax - 1) / 2
	var topLeft, topRight, lowerLeft, lowerRight int
	for _, r := range robots {
		if r.pos.X < xMid && r.pos.Y < yMid {
			topLeft += 1
		} else if r.pos.X > xMid && r.pos.Y < yMid {
			topRight += 1
		} else if r.pos.X < xMid && r.pos.Y > yMid {
			lowerLeft += 1
		} else if r.pos.X > xMid && r.pos.Y > yMid {
			lowerRight += 1
		}
	}

	fmt.Printf("topright: %d\n", topRight)
	fmt.Printf("topleft: %d\n", topLeft)
	fmt.Printf("lowerLeft: %d\n", lowerRight)
	fmt.Printf("lowerRight: %d\n", lowerRight)

	return topLeft * topRight * lowerLeft * lowerRight
}

func moveRobot(robot *robotInfo, xMax, yMax int) {
	robot.pos.X += robot.vel.X
	if robot.pos.X < 0 {
		robot.pos.X = xMax - util.IntAbs(robot.pos.X)
	}
	if robot.pos.X >= xMax {
		robot.pos.X -= xMax
	}

	robot.pos.Y += robot.vel.Y
	if robot.pos.Y < 0 {
		robot.pos.Y = yMax - util.IntAbs(robot.pos.Y)
	}
	if robot.pos.Y >= yMax {
		robot.pos.Y -= yMax
	}
}

func day14Part2(input []string, xMax, yMax int) int {
	robots := make([]*robotInfo, len(input))
	for i, l := range input {
		var xPos, yPos, xVel, yVel int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &xPos, &yPos, &xVel, &yVel)
		robots[i] = &robotInfo{
			pos: util.Coordinate{X: xPos, Y: yPos},
			vel: util.Coordinate{X: xVel, Y: yVel},
		}
	}

	maxIters := xMax * yMax // Cannot be more moves available than this
	var iterCount int
	for i := 0; i < maxIters; i++ {
		for _, r := range robots {
			moveRobot(r, xMax, yMax)
		}
		if found := checkRobotTree(robots, 10); found { // Check there's 10 robots in a row
			fmt.Println("found!")
			iterCount = i + 1
			break
		}
	}

	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			var found bool
			for _, r := range robots {
				if r.pos.X == x && r.pos.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return iterCount
}

func checkRobotTree(robots []*robotInfo, len int) bool {
	for _, r := range robots {
		for l := 1; l <= len; l++ {
			var found bool
			for _, r2 := range robots {
				if r2.pos.X == r.pos.X+l && r2.pos.Y == r.pos.Y+l {
					found = true
					break
				}
			}
			if !found {
				break
			}

			if l == len {
				return true
			}
		}

	}

	return false
}
