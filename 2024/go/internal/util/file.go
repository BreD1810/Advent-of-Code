package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFileLines reads file to a string array
func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if e := scanner.Err(); e != nil {
		log.Fatalln("ReadFileLines error:", e)
	}

	return lines
}

// ReadFileIntLines reads a file of int lines to a 2D array
func ReadFileIntLines(path string) [][]int {
	fileContents := ReadFileLines(path)
	a := make([][]int, len(fileContents))
	for i, l := range fileContents {
		curA := []int{}
		lineVals := strings.Fields(l)
		for _, d := range lineVals {
			intVal, err := strconv.Atoi(d)
			if err != nil {
				panic(fmt.Sprintf("error converting string to int: %s", err))
			}

			curA = append(curA, intVal)
		}
		a[i] = curA
	}

	return a
}
