package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	exampleInp := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 13},
		{"Actual", util.ReadFileLines("../inputs/day13-1.txt"), 6272},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay13Part2(t *testing.T) {
	exampleInp := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 140},
		{"Actual", util.ReadFileLines("../inputs/day13-1.txt"), 22288},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
