package day1

import (
	"advent23/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	fmt.Println("hello")
	var instr, err = utils.ReadLines("day1/day1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var elves []int
	var curElfCals int
	for _, x := range instr {
		if x == "" {
			elves = append(elves, curElfCals)
			curElfCals = 0
		} else {
			var pint, err = strconv.Atoi(x)
			if err != nil {
				os.Exit(0)
			}
			curElfCals = curElfCals + pint
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	var total int
	for i := 0; i < 3; i++ {
		fmt.Println("elf", i, elves[i])
		total += elves[i]
	}
	fmt.Println("total", total)
}
