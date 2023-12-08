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

var cardValues = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var strengths = map[int]int{
	50: 6,
	41: 5,
	32: 4,
	31: 3,
	22: 2,
	21: 1,
	11: 0,
}

func Run() {
	var input, _ = utils.ReadLines("day7/input.txt")
	hands := parse(input)

	fmt.Println("Part 1", calcTotal(hands))

	//convert jacks to jokers
	for k := 0; k < len(hands); k++ {
		for i := 0; i < 5; i++ {
			if hands[k].cards[i] == 11 {
				hands[k].cards[i] = 1
			}
		}
	}

	fmt.Println("Part 2", calcTotal(hands))
}

func calcTotal(hands []Hand) int {
	sort.Slice(hands, func(i, j int) bool {
		return calcStrength(&hands[i]) < calcStrength(&hands[j])
	})

	total := 0
	for i, h := range hands {
		total += h.bid * (i + 1)
	}
	return total
}

func calcStrength(a *Hand) int {
	cardCounts := [15]int{}
	for _, c := range a.cards {
		cardCounts[c]++
	}
	joker := cardCounts[1]
	cardCounts[1] = 0
	sort.Ints(cardCounts[:])
	cardCounts[14] += joker

	strength := strengths[cardCounts[14]*10+cardCounts[13]]

	for i := 0; i < 5; i++ {
		strength += strength*0x10 + a.cards[i]
	}

	return strength
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
