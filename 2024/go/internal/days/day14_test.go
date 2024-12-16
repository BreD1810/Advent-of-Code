package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day14Str = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

// const day14Str = `p=2,4 v=2,-3`

func TestDay14Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day14Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		xMax int
		yMax int
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			xMax: 11,
			yMax: 7,
			exp:  12,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day14.txt"),
			xMax: 101,
			yMax: 103,
			exp:  228690000,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day14Part1(tc.inp, tc.xMax, tc.yMax)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay14Part2(t *testing.T) {
	t.Parallel()
	exampleLines := util.ReadFileLines("../../inputs/day14.txt")
	res := day14Part2(exampleLines, 101, 103)
	assert.Equal(t, 7093, res)
}
