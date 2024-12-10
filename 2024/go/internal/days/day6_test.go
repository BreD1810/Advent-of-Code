package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day6Str = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day6Str, "\n")
	exampleInp := make([][]rune, len(exampleLines))
	for i, l := range exampleLines {
		exampleInp[i] = []rune(l)
	}

	tcs := []struct {
		name string
		inp  [][]rune
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			exp:  42,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day6.txt"),
			exp:  4454,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day6Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay6Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day6Str, "\n")
	exampleInp := make([][]rune, len(exampleLines))
	for i, l := range exampleLines {
		exampleInp[i] = []rune(l)
	}

	tcs := []struct {
		name string
		inp  [][]rune
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			exp:  6,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day6.txt"),
			exp:  1503,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day6Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
