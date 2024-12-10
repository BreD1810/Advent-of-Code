package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day4Str = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDay4Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day4Str, "\n")
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
			exp:  18,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day4.txt"),
			exp:  2530,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay4Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day4Str, "\n")
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
			exp:  9,
		},
		{
			name: "actual",
			inp:  util.ReadFile2DRune("../../inputs/day4.txt"),
			exp:  1921,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
