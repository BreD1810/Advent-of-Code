package main

import (
	"aoc-2022/days"
	"log"
	"os"
)

func main() {
	day := os.Args[1]
	switch day {
	case "1":
		days.Day1()
	case "2":
		days.Day2()
	default:
		log.Fatal("Day not found")
	}
}
