package days

import (
	"aoc-2021/util"
	"errors"
	"fmt"
	"log"
	"strings"
)

type bingoBoard [5][5]*bingoEntry

type bingoEntry struct {
	val    int
	marked bool
}

const (
	noRows = 5
	noCols = 5
)

func Day4() {
	fileContents := util.ReadFileLines("inputs/day04-1.txt")
	// fileContents := util.ReadFileLines("inputs/test.txt")
	fmt.Printf("Part 1: %d\n", day4Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day4Part2(fileContents))
}

func day4Part1(lines []string) int {
	bingoCallStrings := strings.Split(lines[0], ",")
	bingoCalls := make([]int, len(bingoCallStrings))

	for i, call := range bingoCallStrings {
		bingoCalls[i] = util.GetIntFromString(call)
	}

	boards := getBingoBoards(lines[2:])
	winner, winningVal, err := findWinner(boards, bingoCalls)
	if err != nil {
		log.Fatal(err)
	}

	return winner.getScore(winningVal)
}

func day4Part2(lines []string) int {
	bingoCallStrings := strings.Split(lines[0], ",")
	bingoCalls := make([]int, len(bingoCallStrings))

	for i, call := range bingoCallStrings {
		bingoCalls[i] = util.GetIntFromString(call)
	}

	boards := getBingoBoards(lines[2:])
	lastWinner, winningVal, err := findLastWinner(boards, bingoCalls)
	if err != nil {
		log.Fatal(err)
	}

	return lastWinner.getScore(winningVal)
}

func getBingoBoards(lines []string) []bingoBoard {
	boards := make([]bingoBoard, 0)

	for i := 0; i < len(lines); i += 6 {
		var board bingoBoard

		for j := 0; j < noRows; j++ {
			row := strings.Fields(lines[i+j])
			for k := 0; k < noCols; k++ {
				val := util.GetIntFromString(row[k])
				board[j][k] = &bingoEntry{val, false}
			}
		}

		boards = append(boards, board)
	}

	return boards
}

// Returns the winner and the winning call
func findWinner(boards []bingoBoard, bingoCalls []int) (bingoBoard, int, error) {
	for i := 0; i < len(bingoCalls); i++ {
		currentValue := bingoCalls[i]

		for _, board := range boards {
			for j := 0; j < noRows; j++ {
				for k := 0; k < noCols; k++ {
					if board[j][k].val == currentValue {
						board[j][k].marked = true
					}
				}
			}

			if board.checkIfWinner() {
				return board, currentValue, nil
			}
		}
	}

	var defBoard bingoBoard
	return defBoard, 0, errors.New("didn't find a winning board")
}

func findLastWinner(boards []bingoBoard, bingoCalls []int) (bingoBoard, int, error) {
	var lastWinner bingoBoard

	for i := 0; i < len(bingoCalls); i++ {
		currentValue := bingoCalls[i]

		noBoards := len(boards)
		for n := 0; n < noBoards; n++ {
			board := boards[n]
			for j := 0; j < noRows; j++ {
				for k := 0; k < noCols; k++ {
					if board[j][k].val == currentValue {
						board[j][k].marked = true
					}
				}
			}

			if board.checkIfWinner() {
				lastWinner = boards[n]
				boards[n] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
				noBoards--
				n = 0
			}
		}

		if len(boards) == 0 {
			return lastWinner, currentValue, nil
		}
	}

	return lastWinner, 0, errors.New("more than 1 winning board")
}

func (board bingoBoard) checkIfWinner() bool {
	for i := 0; i < noRows; i++ {
		isWinner := true
		for j := 0; j < noCols; j++ {
			isWinner = isWinner && board[i][j].marked
		}
		if isWinner {
			return true
		}
	}

	for i := 0; i < noCols; i++ {
		isWinner := true
		for j := 0; j < noRows; j++ {
			isWinner = isWinner && board[j][i].marked
		}
		if isWinner {
			return true
		}
	}

	return false
}

func (board bingoBoard) getScore(winningVal int) int {
	sum := 0

	for i := 0; i < noRows; i++ {
		for j := 0; j < noCols; j++ {
			if !board[i][j].marked {
				sum += board[i][j].val
			}
		}
	}

	return sum * winningVal
}

func printBoards(boards []bingoBoard) {
	for _, board := range boards {
		for i := 0; i < noRows; i++ {
			for j := 0; j < noCols; j++ {
				if board[i][j].marked {
					fmt.Printf("%d.\t", board[i][j].val)
				} else {
					fmt.Printf("%d\t", board[i][j].val)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
