package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// const day10Str = `0123
// 1234
// 8765
// 9876`

const day10Str = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDay10Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day10Str, "\n")
	exampleInp := make([][]int, len(exampleLines))
	for i, l := range exampleLines {
		curA := make([]int, len(l))
		for j, d := range l {
			curA[j] = int(d) - '0'
		}
		exampleInp[i] = curA
	}

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			// exp:  1,
			exp: 36,
		},
		{
			name: "actual",
			inp:  util.ReadFileIntLines("../../inputs/day10.txt"),
			exp:  794,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day10Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay10Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day10Str, "\n")
	exampleInp := make([][]int, len(exampleLines))
	for i, l := range exampleLines {
		curA := make([]int, len(l))
		for j, d := range l {
			curA[j] = int(d) - '0'
		}
		exampleInp[i] = curA
	}

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			exp:  81,
		},
		{
			name: "actual",
			inp:  util.ReadFileIntLines("../../inputs/day10.txt"),
			exp:  1706,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day10Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
