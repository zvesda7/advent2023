package day9

import (
	"advent23/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	var input, _ = utils.ReadLines("day9/input.txt")
	pats := parse(input)

	sum := 0
	for _, pat := range pats {
		sum += nextInPat(pat)
	}
	fmt.Println(sum)

	sum = 0
	for _, pat := range pats {
		slices.Reverse(pat)
		sum += nextInPat(pat)
	}
	fmt.Println(sum)
}

func nextInPat(pat []int) int {
	pats := [][]int{}
	pats = append(pats, pat)
	for i := 0; true; i++ {
		var pat = make([]int, len(pats[i])-1)
		allZero := true
		for j := 0; j < len(pats[i])-1; j++ {
			pat[j] = pats[i][j+1] - pats[i][j]
			if pat[j] != 0 {
				allZero = false
			}
		}
		pats = append(pats, pat)
		if allZero {
			break
		}
	}

	tot := 0
	for i := len(pats) - 2; i >= 0; i-- {
		tot += pats[i][len(pats[i])-1]
	}
	return tot
}

func parse(input []string) [][]int {
	pats := [][]int{}
	for _, row := range input {
		pat := []int{}
		for _, n := range strings.Split(row, " ") {
			x, _ := strconv.Atoi(n)
			pat = append(pat, x)
		}
		pats = append(pats, pat)
	}
	return pats
}
