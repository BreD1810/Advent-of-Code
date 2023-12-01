package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	exampleInp := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 142},
		{"Actual", util.ReadFileLines("../../inputs/day1.txt"), 55108},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	exampleInp := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 281},
		{"Actual", util.ReadFileLines("../../inputs/day1.txt"), 56324},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
