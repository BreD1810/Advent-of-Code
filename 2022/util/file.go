package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

// ReadFileLine reads a single line file
func ReadFileLine(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	return string(file)
}

// ReadFileInts reads a file to an int array
func ReadFileInts(path string) []int {
	fileContents := ReadFileLines(path)
	fileInts := make([]int, len(fileContents))
	for i, s := range fileContents {
		fileInts[i] = GetIntFromString(s)
	}
	return fileInts
}

// ReadFileIntLines reads a file of int lines to a 2D array
func ReadFileIntLines(path string) [][]int {
	fileContents := ReadFileLines(path)
	a := make([][]int, len(fileContents))
	for i, l := range fileContents {
		curA := make([]int, len(l))
		for j, d := range l {
			curA[j] = int(d) - '0'
			// curA = append(curA, int(d)-'0')
		}
		a[i] = curA
		// a = append(a, curA)
	}

	return a
}
