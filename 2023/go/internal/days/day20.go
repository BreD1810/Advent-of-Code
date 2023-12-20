package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"regexp"
	"strings"
)

func Day20() {
	fileContents := util.ReadFileLines("inputs/day20-example-1.txt")
	fmt.Printf("Part 1: %d\n", day20Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day20Part2(fileContents))
}

type flipFlop struct {
	state   bool
	outputs []string
}

type conjunction struct {
	memories map[string]bool
	outputs  []string
}

func day20Part1(fileContents []string) int {
	startingModules, flipFlops, conjunctions := parsePulseInput(fileContents)

	lowCount, highCount := 0, 0
	for i := 0; i < 1000; i++ {
		l, h := pressButton(startingModules, flipFlops, conjunctions)
		lowCount += l
		highCount += h
	}
	return lowCount * highCount
}

func day20Part2(fileContents []string) int {
	startingModules, flipFlops, conjunctions := parsePulseInput(fileContents)
	return getCountToActivateRx(startingModules, flipFlops, conjunctions)
}

var pulseRegex = regexp.MustCompile(`(.+) -> (.+)`)

func parsePulseInput(inputs []string) ([]string, map[string]flipFlop, map[string]conjunction) {
	var initialTargets []string
	flipFlops := make(map[string]flipFlop)
	conjunctions := make(map[string]conjunction)

	for _, l := range inputs {
		matches := pulseRegex.FindStringSubmatch(l)
		module := matches[1]
		outputs := strings.Split(matches[2], ", ")

		if module == "broadcaster" {
			initialTargets = outputs
			continue
		}

		if strings.HasPrefix(module, "%") {
			flipFlops[module[1:]] = flipFlop{state: false, outputs: outputs}
			continue
		}

		if strings.HasPrefix(module, "&") {
			conjunctions[module[1:]] = conjunction{memories: make(map[string]bool), outputs: outputs}
			continue
		}
	}

	for name, ff := range flipFlops {
		for _, out := range ff.outputs {
			if _, ok := conjunctions[out]; ok {
				conjunctions[out].memories[name] = false
			}
		}
	}

	for name, conj := range conjunctions {
		for _, out := range conj.outputs {
			if _, ok := conjunctions[out]; ok {
				conjunctions[out].memories[name] = false
			}
		}
	}

	return initialTargets, flipFlops, conjunctions
}

type pulseInstruction struct {
	isHigh bool
	source string
	target string
}

func pressButton(startingModules []string, flipFlops map[string]flipFlop, conjunctions map[string]conjunction) (int, int) {
	lowCount, highCount := 1, 0
	var instructions []pulseInstruction
	for _, m := range startingModules {
		instructions = append(instructions, pulseInstruction{isHigh: false, source: "broadcast", target: m})
		lowCount++
	}

	for len(instructions) != 0 {
		curInstruction := instructions[0]
		instructions = instructions[1:]

		if ff, ok := flipFlops[curInstruction.target]; ok {
			if !curInstruction.isHigh {
				sendState := false
				if !ff.state {
					sendState = true
				}

				for _, out := range ff.outputs {
					instructions = append(instructions, pulseInstruction{isHigh: sendState, source: curInstruction.target, target: out})
				}

				ff.state = !ff.state
				flipFlops[curInstruction.target] = ff

				if sendState {
					highCount += len(ff.outputs)
				} else {
					lowCount += len(ff.outputs)
				}
			}
			continue
		}

		if conj, ok := conjunctions[curInstruction.target]; ok {
			conj.memories[curInstruction.source] = curInstruction.isHigh
			conjunctions[curInstruction.target] = conj
			sendState := false
			for _, state := range conj.memories {
				if !state {
					sendState = true
					break
				}
			}
			for _, out := range conj.outputs {
				instructions = append(instructions, pulseInstruction{isHigh: sendState, source: curInstruction.target, target: out})
			}

			if sendState {
				highCount += len(conj.outputs)
			} else {
				lowCount += len(conj.outputs)
			}
		}
	}

	return lowCount, highCount
}

func getCountToActivateRx(startingModules []string, flipFlops map[string]flipFlop, conjunctions map[string]conjunction) int {
	pressCount := 0

	var finalConj string
	for name, conj := range conjunctions {
		for _, out := range conj.outputs {
			if out == "rx" {
				finalConj = name
			}
		}
	}

	finalConjInputs := make(map[string]int)
	totalFinalConjInputs := 0
	for name, ff := range flipFlops {
		for _, out := range ff.outputs {
			if out == finalConj {
				finalConjInputs[name] = 0
				totalFinalConjInputs++
			}
		}
	}
	for name, conj := range conjunctions {
		for _, out := range conj.outputs {
			if out == finalConj {
				finalConjInputs[name] = 0
				totalFinalConjInputs++
			}
		}
	}

	for {
		pressCount += 1
		var instructions []pulseInstruction
		for _, m := range startingModules {
			instructions = append(instructions, pulseInstruction{isHigh: false, source: "broadcast", target: m})
		}
		for len(instructions) != 0 {
			curInstruction := instructions[0]
			instructions = instructions[1:]

			if ff, ok := flipFlops[curInstruction.target]; ok {
				if !curInstruction.isHigh {
					sendState := false
					if !ff.state {
						sendState = true
					}

					for _, out := range ff.outputs {
						if out == "rx" && !sendState {
							return pressCount
						}
						instructions = append(instructions, pulseInstruction{isHigh: sendState, source: curInstruction.target, target: out})
					}

					ff.state = !ff.state
					flipFlops[curInstruction.target] = ff
				}
				continue
			}

			if conj, ok := conjunctions[curInstruction.target]; ok {
				if curInstruction.target == finalConj {
					if curInstruction.isHigh && finalConjInputs[curInstruction.source] == 0 {
						finalConjInputs[curInstruction.source] = pressCount
					}

					isDone := true
					var vals []int
					for _, v := range finalConjInputs {
						if v == 0 {
							isDone = false
							break
						}
						vals = append(vals, v)
					}
					if isDone {
						return util.GetLeastCommonMultiple(vals)
					}
				}

				conj.memories[curInstruction.source] = curInstruction.isHigh
				conjunctions[curInstruction.target] = conj
				sendState := false
				for _, state := range conj.memories {
					if !state {
						sendState = true
						break
					}
				}
				for _, out := range conj.outputs {
					if out == "rx" && !sendState {
						return pressCount
					}
					instructions = append(instructions, pulseInstruction{isHigh: sendState, source: curInstruction.target, target: out})
				}
			}
		}
	}
}
