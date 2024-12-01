package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day1Str = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1(t *testing.T) {
	t.Parallel()
	exampleInp := strings.Split(day1Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{name: "example",
			inp: exampleInp,
			exp: 11},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day1.txt"),
			exp:  1879048,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay1Part2(t *testing.T) {
	t.Parallel()
	exampleInp := strings.Split(day1Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{name: "example",
			inp: exampleInp,
			exp: 31},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day1.txt"),
			exp:  21024792,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
