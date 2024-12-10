package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day5Str = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDay5Part1(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day5Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  143,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day5.txt"),
			exp:  7074,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day5Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay5Part2(t *testing.T) {
	t.Parallel()
	exampleLines := strings.Split(day5Str, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example",
			inp:  exampleLines,
			exp:  123,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day5.txt"),
			exp:  4828,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day5Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
