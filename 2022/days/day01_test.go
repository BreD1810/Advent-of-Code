package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		inp := []string{
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
		}
		exp := 24000

		res := day1Part1(inp)

		if res != exp {
			t.Fatalf("Expected %d, got %d", exp, res)
		}
	})

	t.Run("Input", func(t *testing.T) {
		inp := util.ReadFileLines("../inputs/day01-1.txt")
		exp := 72478

		res := day1Part1(inp)

		if res != exp {
			t.Fatalf("Expected %d, got %d", exp, res)
		}
	})
}

func TestDay1Part2(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		inp := []string{
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
		}
		exp := 45000

		res := day1Part2(inp)

		if res != exp {
			t.Fatalf("Expected %d, got %d", exp, res)
		}
	})

	t.Run("Input", func(t *testing.T) {
		inp := util.ReadFileLines("../inputs/day01-1.txt")
		exp := 210367

		res := day1Part2(inp)

		if res != exp {
			t.Fatalf("Expected %d, got %d", exp, res)
		}
	})
}
