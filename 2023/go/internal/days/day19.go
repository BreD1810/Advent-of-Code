package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"maps"
	"regexp"
	"strings"
)

type workflows map[string]ruleChain

type rule struct {
	field       rune
	greaterThan bool
	value       int
	target      string
}

type ruleChain struct {
	rules  []rule
	target string
}

type machinePart struct {
	x int
	m int
	a int
	s int
}

func Day19() {
	fileContents := util.ReadFileLines("inputs/day19-example.txt")
	fmt.Printf("Part 1: %d\n", day19Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day19Part1(fileContents))
}

func day19Part1(fileContents []string) int {
	wkflws, parts := parseWorkflowsAndParts(fileContents)

	total := 0
	for _, p := range parts {
		if accepted := checkPartAccepted(p, wkflws); accepted {
			total += p.x + p.m + p.a + p.s
		}
	}
	return total
}

func day19Part2(fileContents []string) int {
	wkflws := make(workflows)
	for _, l := range fileContents {
		if l == "" {
			break
		}
		name, rules := parseWorkflow(l)
		wkflws[name] = rules
	}

	ranges := map[rune]util.Coordinate{
		'x': {X: 1, Y: 4000},
		'm': {X: 1, Y: 4000},
		'a': {X: 1, Y: 4000},
		's': {X: 1, Y: 4000},
	}

	return checkPartsDivideAndConquer(wkflws, "in", ranges)
}

func parseWorkflowsAndParts(inp []string) (workflows, []machinePart) {
	wkflws := make(workflows)
	var parts []machinePart
	partsStart := 0

	for i, l := range inp {
		if l == "" {
			partsStart = i + 1
			break
		}
		name, rules := parseWorkflow(l)
		wkflws[name] = rules
	}

	for i := partsStart; i < len(inp); i++ {
		parts = append(parts, parseMachinePart(inp[i]))
	}

	return wkflws, parts
}

var workflowRegex = regexp.MustCompile(`(.*){(.*)}`)

func parseWorkflow(s string) (string, ruleChain) {
	matches := workflowRegex.FindStringSubmatch(s)
	rulesStrings := strings.Split(matches[2], ",")

	rules := make([]rule, len(rulesStrings)-1)
	for i, r := range rulesStrings {
		if !strings.ContainsRune(r, ':') {
			break
		}

		field := r[0]

		symbol := r[1]
		greaterThan := false
		if symbol == '>' {
			greaterThan = true
		}

		splits := strings.Split(r[2:], ":")
		value := util.GetIntFromString(splits[0])
		workflow := splits[1]

		rules[i] = rule{field: rune(field), greaterThan: greaterThan, value: value, target: workflow}
	}

	return matches[1], ruleChain{rules: rules, target: rulesStrings[len(rulesStrings)-1]}
}

var machinePartsRegex = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func parseMachinePart(s string) machinePart {
	matches := machinePartsRegex.FindStringSubmatch(s)
	return machinePart{
		x: util.GetIntFromString(matches[1]),
		m: util.GetIntFromString(matches[2]),
		a: util.GetIntFromString(matches[3]),
		s: util.GetIntFromString(matches[4]),
	}
}

func checkPartAccepted(p machinePart, wkflws workflows) bool {
	curChain := wkflws["in"]
	for {
		endOfChain := true
		for _, r := range curChain.rules {
			var value int
			switch r.field {
			case 'x':
				value = p.x
			case 'm':
				value = p.m
			case 'a':
				value = p.a
			case 's':
				value = p.s
			}
			if r.greaterThan {
				if value > r.value {
					if r.target == "A" {
						return true
					}
					if r.target == "R" {
						return false
					}
					curChain = wkflws[r.target]
					endOfChain = false
					break
				}
			} else {
				if value < r.value {
					if r.target == "A" {
						return true
					}
					if r.target == "R" {
						return false
					}
					curChain = wkflws[r.target]
					endOfChain = false
					break
				}
			}
		}
		if endOfChain {
			if curChain.target == "A" {
				return true
			}
			if curChain.target == "R" {
				return false
			}
			curChain = wkflws[curChain.target]
		}
	}
}

func checkPartsDivideAndConquer(wkflws workflows, curWorkflow string, ranges map[rune]util.Coordinate) int {
	if curWorkflow == "R" {
		return 0
	}
	if curWorkflow == "A" {
		res := 1
		for _, v := range ranges {
			res *= v.Y - v.X + 1
		}
		return res
	}

	total := 0
	curRules := wkflws[curWorkflow].rules
	for i := 0; i <= len(curRules); i++ {
		if i == len(curRules) {
			total += checkPartsDivideAndConquer(wkflws, wkflws[curWorkflow].target, ranges)
			continue
		}

		r := curRules[i]
		curRange := ranges[r.field]
		var trueRange util.Coordinate
		var falseRange util.Coordinate

		if r.greaterThan {
			trueRange = util.Coordinate{X: r.value + 1, Y: curRange.Y}
			falseRange = util.Coordinate{X: curRange.X, Y: r.value}
		} else {
			trueRange = util.Coordinate{X: curRange.X, Y: r.value - 1}
			falseRange = util.Coordinate{X: r.value, Y: curRange.Y}
		}

		if trueRange.X <= trueRange.Y {
			clonedRanges := maps.Clone(ranges)
			clonedRanges[r.field] = trueRange
			total += checkPartsDivideAndConquer(wkflws, r.target, clonedRanges)
		}

		if falseRange.X > falseRange.Y {
			break
		}

		ranges[r.field] = falseRange
	}
	return total
}
