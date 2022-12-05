package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type stackMove struct {
	number      int
	source      int
	destination int
}

func Day5() {
	fileContents := util.ReadFileLines("inputs/day05-1.txt")
	fmt.Printf("Part 1: %q\n", day5Part1(fileContents))
	fmt.Printf("Part 2: %q\n", day5Part2(fileContents))
}

func day5Part1(fileContents []string) string {
	linesProcessed, stacks := readStacks(fileContents)
	stacks = reverseStacks(stacks)
	moves := readMoves(fileContents, linesProcessed)

	final := playMoves(stacks, moves)
	var out string
	for _, s := range final {
		v := s.Peek()
		if v != nil {
			out = fmt.Sprintf("%s%c", out, v)
		}
	}

	return out
}

func readStacks(lines []string) (int, []stack.Stack) {
	var linesProcessed int
	noStacks := ((len(lines[0]) - 3) / 4) + 1
	stacks := make([]stack.Stack, noStacks)

	for _, l := range lines {
		if strings.HasPrefix(l, " 1 ") {
			linesProcessed++
			break
		}

		var stackCounter int
		for j := 0; j < len(l); j += 4 {
			c := l[j+1]
			if c != ' ' {
				stacks[stackCounter].Push(c)
			}
			stackCounter++
		}
		linesProcessed++
	}

	return linesProcessed, stacks
}

func readMoves(lines []string, linesProcessed int) []stackMove {
	moves := make([]stackMove, 0)
	for i := linesProcessed + 1; i < len(lines); i++ {
		line := lines[i]
		splits := strings.Split(line, " ")
		num := util.GetIntFromString(splits[1])
		src := util.GetIntFromString(splits[3])
		dst := util.GetIntFromString(splits[5])

		m := stackMove{number: num, source: src, destination: dst}
		moves = append(moves, m)
	}
	return moves

}

func reverseStacks(stacks []stack.Stack) []stack.Stack {
	for i, s := range stacks {
		newS := stack.New()
		for s.Peek() != nil {
			newS.Push(s.Pop())
		}
		stacks[i] = *newS
	}
	return stacks
}

func playMoves(stacks []stack.Stack, moves []stackMove) []stack.Stack {
	for _, m := range moves {
		for i := 0; i < m.number; i++ {
			v := stacks[m.source-1].Pop()
			stacks[m.destination-1].Push(v)
		}
	}
	return stacks
}

func day5Part2(fileContents []string) string {
	linesProcessed, stacks := readStacks(fileContents)
	stacks = reverseStacks(stacks)
	moves := readMoves(fileContents, linesProcessed)

	final := playMoves2(stacks, moves)
	var out string
	for _, s := range final {
		v := s.Peek()
		if v != nil {
			out = fmt.Sprintf("%s%c", out, v)
		}
	}

	return out
}

func playMoves2(stacks []stack.Stack, moves []stackMove) []stack.Stack {
	for _, m := range moves {
		curStack := stack.New()
		for i := 0; i < m.number; i++ {
			v := stacks[m.source-1].Pop()
			curStack.Push(v)
		}
		for curStack.Peek() != nil {
			stacks[m.destination-1].Push(curStack.Pop())
		}
	}
	return stacks
}
