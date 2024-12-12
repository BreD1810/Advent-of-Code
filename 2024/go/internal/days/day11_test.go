package days

import (
	"aoc-2024/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day11Str = `125 17`

func TestDay11Part1(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{
			name: "example",
			inp:  day11Str,
			exp:  55312,
		},
		{
			name: "actual",
			inp:  util.ReadFileLine("../../inputs/day11.txt"),
			exp:  175006,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day11Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay11Part2(t *testing.T) {
	t.Parallel()

	input := util.ReadFileLine("../../inputs/day11.txt")
	res := day11Part2(input)
	assert.Equal(t, 207961583799296, res)
}
