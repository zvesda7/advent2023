package day12

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {

	var input, _ = utils.ReadLines("day12/input.txt")

	sum := 0
	for _, row := range input {
		pat, counts := parseRow(row)
		combos := calcCombos(pat, counts)
		sum += combos
	}
	fmt.Println("Part 1", sum)

	sum = 0
	for _, row := range input {
		pat, counts := parseRow(row)

		newPat := pat
		newCounts := counts
		for i := 0; i < 4; i++ {
			newPat += "?" + pat
			newCounts = append(newCounts, counts...)
		}

		combos := calcCombos(newPat, newCounts)
		sum += combos

	}
	fmt.Println("Part 2", sum)
}

func calcCombos(pat string, counts []int) int {
	n := len(pat)
	for _, c := range counts {
		n -= c
	}
	n -= len(counts) - 1 //spaces that have to exist.
	k := len(counts) + 1 //gap before first

	memos := map[string]int{}
	cnt := perm([]byte(pat), counts, 1, n, k, &memos)
	return cnt
}

func perm(pat []byte, counts []int, first int, n int, k int, memos *map[string]int) int {
	if k == 1 {
		if _, ok := checkMatch(pat, counts, first, n); ok {
			return 1
		} else {
			return 0
		}
	}
	total := 0
	for x := 0; x <= n; x++ {
		if n-x >= 0 {

			if newPat, ok := checkMatch(pat, counts, first, x); ok {
				params := fmt.Sprintln(newPat, counts[1:], 0, n-x, k-1)
				if val, ok := (*memos)[params]; ok {
					total += val
				} else {
					val := perm(newPat, counts[1:], 0, n-x, k-1, memos)
					(*memos)[params] = val
					total += val
				}
			}
		}
	}
	return total
}

func checkMatch(pat []byte, counts []int, first int, extraSpace int) ([]byte, bool) {
	s := ""

	s += strings.Repeat(".", extraSpace)
	if first == 0 && len(counts) != 0 {
		s += "."
	}
	if len(counts) > 0 {
		s += strings.Repeat("#", counts[0])
	}

	for i := 0; i < len(s); i++ {
		if pat[i] != s[i] && pat[i] != '?' {
			return nil, false
		}
	}
	return pat[len(s):], true
}

func parseRow(row string) (string, []int) {
	s1 := strings.Split(row, " ")
	s2 := strings.Split(s1[1], ",")
	nums := []int{}
	for _, x := range s2 {
		n, _ := strconv.Atoi(x)
		nums = append(nums, n)
	}
	return s1[0], nums
}
