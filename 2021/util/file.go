package util

import (
	"bufio"
	"fmt"
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
	scanner_err := scanner.Err()
	if scanner_err != nil {
		fmt.Println(scanner_err)
	}

	return lines
}
