package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay7Part1(t *testing.T) {
	exampleInp := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	exampleInp2 := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"dir b",
		"10 a.txt",
		"20 b.txt",
		"$ cd a",
		"$ ls",
		"dir b",
		"dir c",
		"12 c.txt",
		"$ cd b",
		"$ ls",
		"5 d.txt",
		"$ cd ..",
		"$ cd ..",
		"$ cd b",
		"$ ls",
		"dir d",
		"$ cd d",
		"$ ls",
		"22 n.txt",
		"$ cd ..",
		"$ cd ..",
		"$ cd a",
		"$ cd c",
		"$ ls",
		"7 f.txt",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 95437},
		{"Example2", exampleInp2, 156},
		{"Actual", util.ReadFileLines("../inputs/day07-1.txt"), 1447046},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay7Part2(t *testing.T) {
	exampleInp := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 24933642},
		{"Actual", util.ReadFileLines("../inputs/day07-1.txt"), 578710},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
