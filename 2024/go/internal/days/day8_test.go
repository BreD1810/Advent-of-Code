package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// const day8Str = `..........
// ..........
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ..........
// ..........
// ..........`
// const day8Str = `..........
// ..........
// ..........
// ....a.....
// ........a.
// .....a....
// ..........
// ..........
// ..........
// ..........`

// const day8Str = `T.........
// ...T......
// .T........
// ..........
// ..........
// ..........
// ..........
// ..........
// ..........
// ..........`

const day8Str = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestDay8Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day8Str, "\n")
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
			exp:  14,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day8.txt"),
			exp:  220,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay8Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day8Str, "\n")
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
			exp:  34,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day8.txt"),
			exp:  813,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
