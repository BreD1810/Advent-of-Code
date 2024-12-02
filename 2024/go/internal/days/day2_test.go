package days

import (
	"aoc-2024/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day2Inp = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestDay2Part1(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{
			name: "example",
			inp:  day2Inp,
			exp:  2,
		},
		{
			name: "actual",
			inp:  util.ReadFileIntLines("../../inputs/day2.txt"),
			exp:  624,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day2Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay2Part2(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{
			name: "example",
			inp:  day2Inp,
			exp:  4,
		},
		{
			name: "actual",
			inp:  util.ReadFileIntLines("../../inputs/day2.txt"),
			exp:  658,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day2Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
