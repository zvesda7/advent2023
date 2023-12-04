package day3

import (
	"advent23/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func Run() {
	part1()
	part2()
}

func part1() {
	var instr, _ = utils.ReadLines("day3/input.txt")
	adjPoints := make(map[int]bool)
	for y, row := range instr {
		for x, cell := range row {
			if cell != '.' && !unicode.IsDigit(cell) {
				adjPoints[(y-1)*10000+(x-1)] = true
				adjPoints[(y-1)*10000+(x-0)] = true
				adjPoints[(y-1)*10000+(x+1)] = true
				adjPoints[(y-0)*10000+(x-1)] = true
				adjPoints[(y-0)*10000+(x+1)] = true
				adjPoints[(y+1)*10000+(x-1)] = true
				adjPoints[(y+1)*10000+(x-0)] = true
				adjPoints[(y+1)*10000+(x+1)] = true
			}
		}
	}

	sumParts := 0
	re := regexp.MustCompile("[0-9]+")
	for y, row := range instr {
		for _, match := range re.FindAllStringIndex(row, -1) {
			//fmt.Println(row[match[0]:match[1]])
			for x := match[0]; x < match[1]; x++ {
				if adjPoints[y*10000+x] {
					partNum, _ := strconv.Atoi(row[match[0]:match[1]])
					sumParts += partNum
					break
				}
			}
		}
	}

	fmt.Println("Part1", sumParts)
}

func part2() {
	var instr, _ = utils.ReadLines("day3/input.txt")
	adjPoints := make(map[int]int)
	gearNum := 0
	for y, row := range instr {
		for x, cell := range row {
			if cell == '*' {
				adjPoints[(y-1)*10000+(x-1)] = gearNum
				adjPoints[(y-1)*10000+(x-0)] = gearNum
				adjPoints[(y-1)*10000+(x+1)] = gearNum
				adjPoints[(y-0)*10000+(x-1)] = gearNum
				adjPoints[(y-0)*10000+(x+1)] = gearNum
				adjPoints[(y+1)*10000+(x-1)] = gearNum
				adjPoints[(y+1)*10000+(x-0)] = gearNum
				adjPoints[(y+1)*10000+(x+1)] = gearNum
				gearNum++
			}
		}
	}

	partsByGear := make(map[int][]int)
	re := regexp.MustCompile("[0-9]+")
	for y, row := range instr {
		for _, match := range re.FindAllStringIndex(row, -1) {
			for x := match[0]; x < match[1]; x++ {
				if gearId, ok := adjPoints[y*10000+x]; ok {
					partNum, _ := strconv.Atoi(row[match[0]:match[1]])
					partsByGear[gearId] = append(partsByGear[gearId], partNum)
					break
				}
			}
		}
	}

	sumProduct := 0
	for _, parts := range partsByGear {
		if len(parts) == 2 {
			sumProduct += (parts[0] * parts[1])
		}
	}

	fmt.Println("Part2", sumProduct)
}
