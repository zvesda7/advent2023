package day1

import (
	"advent23/utils"
	"fmt"
	"strings"
)

func Day1() {
	part1()
	part2()
}

func part1() {
	var instr, _ = utils.ReadLines("day1/input.txt")

	sum := 0
	for _, row := range instr {
		firstNum := -1
		lastNum := -1
		for _, cell := range row {
			if cell >= '0' && cell <= '9' {
				if firstNum == -1 {
					firstNum = int(cell) - '0'
					lastNum = firstNum
				} else {
					lastNum = int(cell) - '0'
				}
			}
		}
		fmt.Println(firstNum, lastNum)
		if firstNum != -1 {
			sum += firstNum*10 + lastNum
		}
	}
	fmt.Println("part 1 answer", sum)
}

func part2() {
	var instr, _ = utils.ReadLines("day1/input.txt")

	tokens := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	sum := 0
	for _, row := range instr {
		firstNum := -1
		lastNum := -1
		for i, _ := range row {
			for k, token := range tokens {
				if strings.Index(row[i:], token) == 0 {
					if firstNum == -1 {
						firstNum = k % 10
						lastNum = firstNum
					} else {
						lastNum = k % 10
					}
				}
			}
		}
		fmt.Println(firstNum, lastNum)
		if firstNum != -1 {
			sum += firstNum*10 + lastNum
		}
	}
	fmt.Println("part 2 answer", sum)
}
