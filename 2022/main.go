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
	case "3":
		days.Day3()
	case "4":
		days.Day4()
	case "5":
		days.Day5()
	case "6":
		days.Day6()
	default:
		log.Fatal("Day not found")
	}
}