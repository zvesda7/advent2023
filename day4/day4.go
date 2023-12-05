package day4

import (
	"advent23/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Run() {
	var input, _ = utils.ReadLines("day4/input.txt")
	getNums := regexp.MustCompile("[0-9]+")

	totalPoints := 0
	cardCounts := make(map[int]int)
	allCards := 0
	for cardI, card := range input {
		split1 := strings.Split(card, ":")
		split2 := strings.Split(split1[1], "|")
		winNums := make(map[string]bool)
		for _, num := range getNums.FindAllString(split2[1], -1) {
			winNums[num] = true
		}
		cardCount := 0
		for _, num := range getNums.FindAllString(split2[0], -1) {
			if winNums[num] {
				cardCount++
			}
		}

		//part1
		if cardCount > 0 {
			totalPoints += int(math.Pow(2, float64(cardCount)-1))
		}

		//part2
		cardCounts[cardI]++ //original
		//copies
		for i := 0; i < cardCounts[cardI]; i++ {
			for j := 0; j < cardCount; j++ {
				cardCounts[cardI+j+1]++
			}
		}
		allCards += cardCounts[cardI]
	}
	fmt.Println("Part1", totalPoints)
	fmt.Println("Part2", allCards)
}
