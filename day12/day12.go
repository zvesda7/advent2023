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
	//for _, row := range input {
	//	pat, counts := parseRow(row)
	//	combos := calcCombos(pat, counts)
	//	sum += combos
	//}
	//fmt.Println("Part 1", sum)

	sum = 0
	for _, row := range input[3:4] {
		pat, counts := parseRow(row)
		newPat := pat
		newCounts := counts
		for i := 0; i < 4; i++ {
			newPat += "?" + pat
			newCounts = append(newCounts, counts...)
		}
		combos := calcCombos(newPat, newCounts)
		fmt.Println(newPat, newCounts, combos)
		sum += combos
	}
	fmt.Println("Part 2", sum)
}

func match(pat string, actual string) bool {
	for i := 0; i < len(pat); i++ {
		if pat[i] != actual[i] && pat[i] != '?' {
			return false
		}
	}
	return true
}

func calcCombos(pat string, counts []int) int {
	comboS := combos(len(pat), counts)
	count := 0
	for _, combo := range comboS {
		if match(pat, combo) {
			count++
		}
	}
	return count
}

func combos(n int, counts []int) []string {
	for _, c := range counts {
		n -= c
	}
	n -= len(counts) - 1 //spaces that have to exist.
	k := len(counts) + 1 //gap before first
	perms := [][]int{}
	fmt.Println("complexity", n, k)
	perm(n, k, []int{}, &perms)

	final := []string{}
	for _, p := range perms {
		s := ""
		for i, c := range p {
			s += strings.Repeat(".", c)
			if i > 0 && i < len(p)-1 {
				s += "."
			}
			if i < len(counts) {
				s += strings.Repeat("#", counts[i])
			}
		}
		final = append(final, s)
		//fmt.Println(s, p)
	}

	return final
}

func perm(n int, k int, prefix []int, running *[][]int) {
	if k == 1 {
		prefix := append(prefix, n)
		temp := make([]int, len(prefix))
		copy(temp, prefix)
		(*running) = append((*running), temp)
		//fmt.Println(n, k, temp)
	} else {
		for x := 0; x <= n; x++ {
			if n-x >= 0 {
				newPre := append(prefix, x)
				perm(n-x, k-1, newPre, running)
			}
		}
	}
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
