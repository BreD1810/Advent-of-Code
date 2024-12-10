package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day7Str = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestDay7Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day7Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  3749,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day7.txt"),
			exp:  5837374519342,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay7Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day7Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  11387,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day7.txt"),
			exp:  492383931650959,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
