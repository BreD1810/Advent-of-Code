package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	exampleInp := []string{
		"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
		"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
		"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
		"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
		"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
		"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
		"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
		"Valve HH has flow rate=22; tunnel leads to valve GG",
		"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
		"Valve JJ has flow rate=21; tunnel leads to valve II",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 1651},
		{"Actual", util.ReadFileLines("../inputs/day16-1.txt"), 2029},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day16Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay16Part2(t *testing.T) {
	exampleInp := []string{
		"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
		"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
		"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
		"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
		"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
		"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
		"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
		"Valve HH has flow rate=22; tunnel leads to valve GG",
		"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
		"Valve JJ has flow rate=21; tunnel leads to valve II",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 1704},
		{"Actual", util.ReadFileLines("../inputs/day16-1.txt"), 2723},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day16Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
