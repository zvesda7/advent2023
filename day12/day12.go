package day12

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {

	var input, _ = utils.ReadLines("day12/test.txt")
	//fmt.Println(combos(7, []int{2, 1}))

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

		combos1 := calcCombos(pat, counts)
		combos2 := calcCombos(pat+"?"+pat, append(counts, counts...))
		combos3 := calcCombos(pat+"?"+pat+"?"+pat, append(append(counts, counts...), counts...))
		combos4 := calcCombos(pat+"?"+pat+"?"+pat+"?"+pat, append(append(append(counts, counts...), counts...), counts...))
		ratio := combos2 / combos1
		ratio2 := combos3 / combos2
		ratio3 := combos4 / combos3
		if ratio != ratio2 {
			fmt.Println(pat, counts, combos1, combos2, combos3, combos4, ratio, ratio2, ratio3, "bad")
		} else {
			fmt.Println(pat, counts, combos1, ratio, ratio2, "good")
		}
		total := combos1
		for i := 0; i < 4; i++ {
			total *= ratio
		}

		sum += total
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

	return perm([]byte(pat), counts, n, k, []int{})
}

func perm(pat []byte, counts []int, n int, k int, prefix []int) int {

	if k == 1 {
		newPre := append(prefix, n)

		if checkMatch(pat, counts, newPre) {
			return 1
		} else {
			return 0
		}
	}
	total := 0
	for x := 0; x <= n; x++ {
		if n-x >= 0 {
			newPre := append(prefix, x)
			if checkMatch(pat, counts, newPre) {
				total += perm(pat, counts, n-x, k-1, newPre)
			}
		}
	}
	return total
}

func checkMatch(pat []byte, counts []int, prefix []int) bool {
	s := ""
	for i, c := range prefix {
		s += strings.Repeat(".", c)
		if i > 0 && i < len(counts) {
			s += "."
		}
		if i < len(counts) {
			s += strings.Repeat("#", counts[i])
		}
	}
	for i := 0; i < len(s); i++ {
		if pat[i] != s[i] && pat[i] != '?' {
			return false
		}
	}
	return true
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
