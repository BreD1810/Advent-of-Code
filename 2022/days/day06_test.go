package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay6Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example1", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"Example2", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"Example3", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"Example4", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"Example4", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
		{"Actual", util.ReadFileLine("../inputs/day06-1.txt"), 1833},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day6Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay6Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example1", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"Example2", "bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"Example3", "nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"Example4", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"Example4", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
		{"Actual", util.ReadFileLine("../inputs/day06-1.txt"), 3425},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day6Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
