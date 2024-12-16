package days

import (
	"aoc-2024/internal/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day15Str1 = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`

const day15Str2 = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const day15Str3 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func TestDay15Part1(t *testing.T) {
	t.Parallel()
	exampleLines1 := strings.Split(day15Str1, "\n")
	exampleLines2 := strings.Split(day15Str2, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{
			name: "example1",
			inp:  exampleLines1,
			exp:  2028,
		},
		{
			name: "example2",
			inp:  exampleLines2,
			exp:  10092,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day15.txt"),
			exp:  1577255,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day15Part1(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestDay15Part2(t *testing.T) {
	t.Parallel()
	exampleLines2 := strings.Split(day15Str2, "\n")
	// exampleLines3 := strings.Split(day15Str3, "\n")

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		// {
		// 	name: "example1",
		// 	inp:  exampleLines3,
		// 	exp:  -1,
		// },
		{
			name: "example2",
			inp:  exampleLines2,
			exp:  9021,
		},
		{
			name: "actual",
			inp:  util.ReadFileLines("../../inputs/day15.txt"),
			exp:  1597035,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day15Part2(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
