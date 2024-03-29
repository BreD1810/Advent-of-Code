package main

import (
	"aoc-2021/days"
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
	case "7":
		days.Day7()
	case "8":
		days.Day8()
	case "9":
		days.Day9()
	case "10":
		days.Day10()
	case "11":
		days.Day11()
	case "12":
		days.Day12()
	}
}
