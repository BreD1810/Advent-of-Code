package days

import (
	"aoc-2022/util"
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func Day13() {
	fileContents := util.ReadFileLines("inputs/day13-1.txt")
	fmt.Printf("Part 1: %d\n", day13Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day13Part2(fileContents))
}

type distressPair struct {
	left  any
	right any
}

func day13Part1(fileContents []string) int {
	dps := parseDistressPairs(fileContents)
	for _, p := range dps {
		fmt.Println(p)
	}
	return sumRightOrder(dps)
}

func parseDistressPairs(fileContents []string) []distressPair {
	var dps []distressPair

	for i := 0; i < len(fileContents); i += 3 {
		dp := distressPair{
			left:  parseDitressSignal(fileContents[i]),
			right: parseDitressSignal(fileContents[i+1]),
		}
		dps = append(dps, dp)
	}

	return dps
}

func parseDitressSignal(s string) any {
	var ds any
	if err := json.Unmarshal([]byte(s), &ds); err != nil {
		log.Fatalf("Unable to parse signal %s\n", s)
	}
	return ds
}

func sumRightOrder(dps []distressPair) int {
	var sum int

	for i, dp := range dps {
		if checkRightOrder(dp.left, dp.right) <= 0 {
			sum += i + 1
		}
	}

	return sum
}

func checkRightOrder(left any, right any) int {
	lInt, okL := left.(float64)
	rInt, okR := right.(float64)
	if okL && okR {
		return int(lInt) - int(rInt)
	}

	var lList []any
	switch left.(type) {
	case []any, []float64:
		lList = left.([]any)
	case float64:
		lList = []any{left}
	}

	var rList []any
	switch right.(type) {
	case []any, []float64:
		rList = right.([]any)
	case float64:
		rList = []any{right}
	}

	for i := range lList {
		if len(rList) <= i {
			return 1
		}
		if res := checkRightOrder(lList[i], rList[i]); res != 0 {
			return res
		}
	}

	if len(lList) == len(rList) {
		return 0
	}
	return -1
}

func day13Part2(fileContents []string) int {
	var packets []any
	for _, l := range fileContents {
		if l == "" {
			continue
		}
		packets = append(packets, parseDitressSignal(l))
	}
	packets = append(packets, parseDitressSignal("[[2]]"))
	packets = append(packets, parseDitressSignal("[[6]]"))

	sort.Slice(packets, func(i, j int) bool {
		return checkRightOrder(packets[i], packets[j]) < 0
	})

	var decoderKey int
	for i, p := range packets {
		s, _ := json.Marshal(p)
		if string(s) == "[[2]]" || string(s) == "[[6]]" {
			if decoderKey != 0 {
				return decoderKey * (i + 1)
			}
			decoderKey = i + 1
		}
	}
	log.Fatalln("Didn't find both divider packets")
	return 0
}
