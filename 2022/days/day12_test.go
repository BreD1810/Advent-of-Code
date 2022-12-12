package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	exampleInp := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 31},
		{"Actual", util.ReadFileLines("../inputs/day12-1.txt"), 462},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay12Part2(t *testing.T) {
	exampleInp := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 29},
		{"Actual", util.ReadFileLines("../inputs/day12-1.txt"), 451},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	b.Run("Single threaded", benchSingleThread)
	b.Run("Fan out", benchFanOut)
	b.Run("Pooling", benchPool)
}

var outSave int

func benchSingleThread(b *testing.B) {
	var res int
	inp := util.ReadFileLines("../inputs/day12-1.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = day12Part2Single(inp)
	}
	outSave = res
}

func benchFanOut(b *testing.B) {
	var res int
	inp := util.ReadFileLines("../inputs/day12-1.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = day12Part2FanOut(inp)
	}
	outSave = res
}

func benchPool(b *testing.B) {
	var res int
	inp := util.ReadFileLines("../inputs/day12-1.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = day12Part2Pooling(inp)
	}
	outSave = res
}
