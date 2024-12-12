package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day12Str1 = `AAAA
BBCD
BBCC
EEEC`

const day12Str2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const day12Str3 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

const day12Str4 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

const day12Str = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestDay12Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day12Str, "\n")
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
			// exp:  140,
			// exp: 772,
			exp: 1930,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day12.txt"),
			exp:  1473276,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay12Part2(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		inp  [][]rune
		exp  int
	}{
		{
			name: "example1",
			inp:  get2dRuneInput(day12Str1),
			exp:  80,
		},
		{
			name: "example2",
			inp:  get2dRuneInput(day12Str2),
			exp:  436,
		},
		{
			name: "example3",
			inp:  get2dRuneInput(day12Str3),
			exp:  236,
		},
		{
			name: "example4",
			inp:  get2dRuneInput(day12Str4),
			exp:  368,
		},
		{
			name: "example",
			inp:  get2dRuneInput(day12Str),
			exp:  1206,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day12.txt"),
			exp:  901100,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func get2dRuneInput(s string) [][]rune {
	exampleLines := strings.Split(s, "\n")
	exampleInp := make([][]rune, len(exampleLines))
	for i, l := range exampleLines {
		exampleInp[i] = []rune(l)
	}
	return exampleInp
}
