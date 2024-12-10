package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	t.Parallel()
	exampleInp := strings.Split(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			exp:  161,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day3.txt"),
			exp:  174336360,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day3Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay3Part2(t *testing.T) {
	t.Parallel()
	exampleInp := strings.Split(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleInp,
			exp:  48,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day3.txt"),
			exp:  88802350,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day3Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
