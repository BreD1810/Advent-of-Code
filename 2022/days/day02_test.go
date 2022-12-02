package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		inp := []string{"A Y", "B X", "C Z"}
		exp := 15
		res := day2Part1(inp)

		if res != exp {
			t.Fatalf("expected %d, got %d", exp, res)
		}
	})

	t.Run("Actual", func(t *testing.T) {
		inp := util.ReadFileLines("../inputs/day02-1.txt")
		exp := 10816
		res := day2Part1(inp)

		if res != exp {
			t.Fatalf("expected %d, got %d", exp, res)
		}
	})
}

func TestDay2Part2(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		inp := []string{"A Y", "B X", "C Z"}
		exp := 12
		res := day2Part2(inp)

		if res != exp {
			t.Fatalf("expected %d, got %d", exp, res)
		}
	})

	t.Run("Actual", func(t *testing.T) {
		inp := util.ReadFileLines("../inputs/day02-1.txt")
		exp := 11657
		res := day2Part2(inp)

		if res != exp {
			t.Fatalf("expected %d, got %d", exp, res)
		}
	})
}
