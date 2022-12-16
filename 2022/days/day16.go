package days

import (
	"aoc-2022/util"
	"fmt"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/ernestosuarez/itertools"
)

func Day16() {
	fileContents := util.ReadFileLines("inputs/day03-1.txt")
	fmt.Printf("Part 1: %d\n", day16Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day16Part2(fileContents))
}

func day16Part1(fileContents []string) int {
	valveNameMapping, flowRates, connections := parseValveTunnels(fileContents)
	fastestRouteMatrix := calcFastestRouteMatrix(connections, valveNameMapping)

	var valvesWithFlow []string
	for valve, rate := range flowRates {
		if rate != 0 {
			valvesWithFlow = append(valvesWithFlow, valve)
		}
	}

	return calcMaxReleased(fastestRouteMatrix, flowRates, 0, 30, 0, 0, "AA", valvesWithFlow)
}

func parseValveTunnels(fileContents []string) (map[string]int, map[string]int, map[string][]string) {
	// var valves []valve
	valveNameMapping := make(map[string]int)
	flowRates := make(map[string]int)
	connections := make(map[string][]string)

	for i, l := range fileContents {
		l = strings.TrimPrefix(l, "Valve ")
		firstSplit := strings.Split(l, " has flow rate=")
		valveName := firstSplit[0]
		secondSplit := strings.Split(firstSplit[1], ";")
		flowRate := util.GetIntFromString(secondSplit[0])
		valveList := strings.TrimPrefix(secondSplit[1], " tunnels lead to valves ")
		valveList = strings.TrimPrefix(valveList, " tunnel leads to valve ")
		tunnelsTo := strings.Split(valveList, ", ")

		valveNameMapping[valveName] = i
		flowRates[valveName] = flowRate
		connections[valveName] = append(connections[valveName], tunnelsTo...)
	}

	return valveNameMapping, flowRates, connections
}

func calcFastestRouteMatrix(connections map[string][]string, valveNameMapping map[string]int) map[string]map[string]int {
	g := dijkstra.NewGraph()

	for _, v := range valveNameMapping {
		g.AddVertex(v)
	}

	for s, ds := range connections {
		sourceName := valveNameMapping[s]
		for _, d := range ds {
			destName := valveNameMapping[d]
			g.AddArc(sourceName, destName, 1)
		}
	}

	matrix := make(map[string]map[string]int)
	for name1, value1 := range valveNameMapping {
		matrix[name1] = make(map[string]int)
		for name2, value2 := range valveNameMapping {
			best, _ := g.Shortest(value1, value2)
			matrix[name1][name2] = int(best.Distance)
		}
	}

	return matrix
}

func calcMaxReleased(fastestRouteMatrix map[string]map[string]int, flowRates map[string]int, currentTime int, timeLimit int, currentReleased int, currentFlow int, currentLoc string, remaining []string) int {
	releasedWithoutMoving := currentReleased + (timeLimit-currentTime)*currentFlow
	maxReleased := releasedWithoutMoving

	for _, v := range remaining {
		timeToTravelAndOpen := fastestRouteMatrix[currentLoc][v] + 1
		if currentTime+timeToTravelAndOpen < timeLimit {
			newTime := currentTime + timeToTravelAndOpen
			newReleased := currentReleased + timeToTravelAndOpen*currentFlow
			newFlow := currentFlow + flowRates[v]
			newRemaining := util.RemoveFromStringSlice(v, remaining)
			finalPressure := calcMaxReleased(fastestRouteMatrix, flowRates, newTime, timeLimit, newReleased, newFlow, v, newRemaining)
			if maxReleased < finalPressure {
				maxReleased = finalPressure
			}
		}
	}

	return maxReleased
}

func day16Part2(fileContents []string) int {
	valveNameMapping, flowRates, connections := parseValveTunnels(fileContents)
	fastestRouteMatrix := calcFastestRouteMatrix(connections, valveNameMapping)

	var valvesWithFlow []string
	for valve, rate := range flowRates {
		if rate != 0 {
			valvesWithFlow = append(valvesWithFlow, valve)
		}
	}

	var toSend [][]string
	for i := 1; i < len(valvesWithFlow)/2; i++ {
		for myValves := range itertools.CombinationsStr(valvesWithFlow, i) {
			toSend = append(toSend, myValves)
		}
	}

	workCh := make(chan []string, len(toSend))
	releasedCh := make(chan int, len(toSend))

	for t := 0; t < len(toSend); t++ {
		go func() {
			myValves := <-workCh
			maxReleasedMe := calcMaxReleased(fastestRouteMatrix, flowRates, 0, 26, 0, 0, "AA", myValves)

			elephantValves := getElephantValves(valvesWithFlow, myValves)
			maxReleasedElephants := calcMaxReleased(fastestRouteMatrix, flowRates, 0, 26, 0, 0, "AA", elephantValves)

			releasedCh <- maxReleasedMe + maxReleasedElephants
		}()
	}

	for i := 0; i < len(toSend); i++ {
		workCh <- toSend[i]
	}

	var maxReleased int
	for i := 0; i < len(toSend); i++ {
		released := <-releasedCh
		if maxReleased < released {
			maxReleased = released
		}
	}

	return maxReleased
}

func getElephantValves(allValves, myValves []string) (elephantValves []string) {
	for _, v := range allValves {
		found := false
		for _, mv := range myValves {
			if v == mv {
				found = true
				break
			}
		}
		if !found {
			elephantValves = append(elephantValves, v)
		}
	}

	return
}
