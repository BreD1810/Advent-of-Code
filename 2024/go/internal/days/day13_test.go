package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day13Str = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestDay13Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day13Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  480,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day13.txt"),
			exp:  36758,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay13Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day13Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  875318608908,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day13.txt"),
			exp:  76358113886726,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
