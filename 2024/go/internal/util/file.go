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
