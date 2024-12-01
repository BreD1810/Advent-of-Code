package main

import (
	"aoc-2024/internal/days"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "no day number supplied")
		os.Exit(1)
	}
	day := os.Args[1]
	switch day {
	case "1":
		days.Day1()
	default:
		fmt.Fprintln(os.Stderr, "no day number supplied")
		os.Exit(1)
	}
}
