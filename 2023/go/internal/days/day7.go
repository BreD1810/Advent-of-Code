package days

import (
	"aoc-2023/internal/util"
	"fmt"
	"log"
	"sort"
	"strings"
)

func Day7() {
	fileContents := util.ReadFileLines("inputs/day7-actual.txt")
	fmt.Printf("Part 1: %d\n", day7Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day7Part2(fileContents))
}

type camelCard int

const (
	Two camelCard = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	J
	Q
	K
	A
)

type handStrength int

const (
	HighCard handStrength = iota
	OnePair
	TwoPair
	ThreeOAK
	FullHouse
	FourOAK
	FiveOAK
)

type handInfo struct {
	cardCounts map[camelCard]int
	cards      []camelCard
	bid        int
}

type handInfos []handInfo

func (his handInfos) Len() int {
	return len(his)
}

func (his handInfos) Swap(i, j int) {
	his[i], his[j] = his[j], his[i]
}

func (his handInfos) Less(i, j int) bool {
	strengthI := determineHandStrength(his[i])
	strengthJ := determineHandStrength(his[j])

	if strengthI == strengthJ {
		iCards := his[i].cards
		jCards := his[j].cards

		for n, ci := range iCards {
			cj := jCards[n]
			if ci != cj {
				return ci < cj
			}
		}
	}

	return strengthI < strengthJ
}

func day7Part1(fileContents []string) int {
	hands := make(handInfos, len(fileContents))
	for i, l := range fileContents {
		ss := strings.Split(l, " ")
		cardCounts := make(map[camelCard]int)
		cards := make([]camelCard, 5)

		runes := []rune(ss[0])
		for j := 0; j < 5; j++ {
			card := parseCamelCard(runes[j])
			cards[j] = card
			cardCounts[card] += 1
		}

		hands[i] = handInfo{cards: cards, cardCounts: cardCounts, bid: util.GetIntFromString(ss[1])}
	}

	sort.Sort(hands)

	total := 0
	for rank, h := range hands {
		total += (rank + 1) * h.bid
	}
	return total
}

func parseCamelCard(c rune) camelCard {
	switch c {
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return T
	case 'J':
		return J
	case 'Q':
		return Q
	case 'K':
		return K
	case 'A':
		return A
	}
	log.Fatalf("couldn't parse camel card %c", c)
	return Two
}

func determineHandStrength(hi handInfo) handStrength {
	counts := hi.cardCounts
	switch len(counts) {
	case 1:
		return FiveOAK
	case 2:
		for _, v := range counts {
			if v == 4 || v == 1 {
				return FourOAK
			}
			if v == 3 || v == 2 {
				return FullHouse
			}
		}
	case 3:
		for _, v := range counts {
			if v == 3 {
				return ThreeOAK
			}
			if v == 2 {
				return TwoPair
			}
		}
	case 4:
		return OnePair
	default:
		return HighCard
	}
	log.Fatalf("unable to determine strength of %+v", hi)
	return HighCard
}

type camelCard2 int

const (
	J2   camelCard2 = iota
	Two2            = iota
	Three2
	Four2
	Five2
	Six2
	Seven2
	Eight2
	Nine2
	T2
	Q2
	K2
	A2
)

type handInfos2 []handInfo2

func (his handInfos2) Len() int {
	return len(his)
}

func (his handInfos2) Swap(i, j int) {
	his[i], his[j] = his[j], his[i]
}

func (his handInfos2) Less(i, j int) bool {
	strengthI := determineHandStrength2(his[i])
	strengthJ := determineHandStrength2(his[j])

	if strengthI == strengthJ {
		iCards := his[i].cards
		jCards := his[j].cards

		for n, ci := range iCards {
			cj := jCards[n]
			if ci != cj {
				return ci < cj
			}
		}
	}

	return strengthI < strengthJ
}

type handInfo2 struct {
	cardCounts map[camelCard2]int
	cards      []camelCard2
	bid        int
	jokers     int
}

func day7Part2(fileContents []string) int {
	hands := make(handInfos2, len(fileContents))
	for i, l := range fileContents {
		ss := strings.Split(l, " ")
		cardCounts := make(map[camelCard2]int)
		cards := make([]camelCard2, 5)
		jokers := 0

		runes := []rune(ss[0])
		for j := 0; j < 5; j++ {
			card := parseCamelCard2(runes[j])
			if card == J2 {
				jokers++
				continue
			}
			cards[j] = card
			cardCounts[card] += 1
		}

		hands[i] = handInfo2{cards: cards, cardCounts: cardCounts, bid: util.GetIntFromString(ss[1]), jokers: jokers}
	}

	sort.Sort(hands)

	total := 0
	for rank, h := range hands {
		total += (rank + 1) * h.bid
		fmt.Printf("hand: %+v, strength: %d\n", h.cards, determineHandStrength2(h))
	}
	return total
}

func parseCamelCard2(c rune) camelCard2 {
	switch c {
	case '2':
		return Two2
	case '3':
		return Three2
	case '4':
		return Four2
	case '5':
		return Five2
	case '6':
		return Six2
	case '7':
		return Seven2
	case '8':
		return Eight2
	case '9':
		return Nine2
	case 'T':
		return T2
	case 'J':
		return J2
	case 'Q':
		return Q2
	case 'K':
		return K2
	case 'A':
		return A2
	}
	log.Fatalf("couldn't parse camel card %c", c)
	return J2
}

func determineHandStrength2(hi handInfo2) handStrength {
	counts := hi.cardCounts

	mostCard, mostCardCount := getMostCards(counts, nil)
	_, secondMostCardCount := getMostCards(counts, &mostCard)

	switch mostCardCount + hi.jokers {
	case 5:
		return FiveOAK
	case 4:
		return FourOAK
	case 3:
		if secondMostCardCount == 2 {
			return FullHouse
		}
		return ThreeOAK
	case 2:
		if secondMostCardCount == 2 {
			return TwoPair
		}
		return OnePair
	case 1:
		return HighCard
	}

	return HighCard
}

func getMostCards(cardCounts map[camelCard2]int, exclude *camelCard2) (camelCard2, int) {
	largestCount := 0
	var largestCard camelCard2

	for c, v := range cardCounts {
		if exclude != nil && c == *exclude {
			continue
		}

		if v > largestCount {
			largestCount = v
			largestCard = c
		}
	}
	return largestCard, largestCount
}
