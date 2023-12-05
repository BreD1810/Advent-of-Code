package main

import (
	"aoc-2023/internal/days"
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
	case "2":
		days.Day2()
	case "3":
		days.Day3()
	case "4":
		days.Day4()
	case "5":
		days.Day5()
	default:
		fmt.Fprintln(os.Stderr, "no day number supplied")
		os.Exit(1)
	}
}
