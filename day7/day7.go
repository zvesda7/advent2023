package day7

import (
	"advent23/utils"
	"fmt"
	"sort"
	"strconv"
)

type Hand struct {
	cards [5]int
	bid   int
}

type HandInfo struct {
	fiveOfAKind  int
	fourOfAKind  int
	threeOfAKind int
	pairs        [2]int //max 2
	cards        [5]int //max 5
}

var cardValues = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func Run() {
	var input, _ = utils.ReadLines("day7/test.txt")
	hands := parse(input)
	fmt.Println(hands)
}

func HandCompare(a, b Hand) bool {

}

func calcStrength(a *Hand) int {
	info := HandInfo{}
	cardCounts := [13]int{}
	for _, c := range a.cards {
		cardCounts[c]++
	}

	pairI, cardI := 0, 0
	for i, cnt := range cardCounts {
		if cnt == 5 {
			info.fiveOfAKind = i
		} else if cnt == 4 {
			info.fourOfAKind = i
		} else if cnt == 3 {
			info.threeOfAKind = i
		} else if cnt == 2 {
			info.pairs[pairI] = i
			pairI++
		} else if cnt == 1 {
			info.cards[cardI] = i
			cardI++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(info.pairs[:])))
	sort.Sort(sort.Reverse(sort.IntSlice(info.cards[:])))
	
	placeMultipliers := []int{0x10000, 0x1000, 0x100, 0x10, 0x1, 0}

	strength := 0
	if info.fiveOfAKind > 0 {
		return info.fiveOfAKind * 
	}
}

func parse(input []string) []Hand {
	hands := []Hand{}
	for _, hs := range input {
		h := Hand{}
		h.bid, _ = strconv.Atoi(hs[6:])
		for i := 0; i < 5; i++ {
			h.cards[i] = cardValues[hs[i]]
		}
		hands = append(hands, h)
	}
	return hands
}
