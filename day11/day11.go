package day11

import (
	"advent23/utils"
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

func Run() {

	var input, _ = utils.ReadLines("day11/input.txt")

	fmt.Println("Part 1", calculate(input, 2))
	fmt.Println("Part 2", calculate(input, 1000000))
}

func calculate(input []string, space int) int {
	emptyRows := map[int]bool{}
	for y := 0; y < len(input); y++ {
		if strings.Index(input[y], "#") == -1 {
			emptyRows[y] = true
		}
	}
	emptyCols := map[int]bool{}
	for x := 0; x < len(input[0]); x++ {
		isEmpty := true
		for y := 0; y < len(input); y++ {
			if input[y][x] == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			emptyCols[x] = true
		}
	}

	points := []Point{}
	y_exp := 0
	for y, row := range input {
		if emptyRows[y] {
			y_exp += space - 1
		}
		x_exp := 0
		for x, cell := range row {
			if emptyCols[x] {
				x_exp += space - 1
			}
			if cell == '#' {
				points = append(points, Point{x + x_exp, y + y_exp})
			}
		}
	}

	sum := 0
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			sum += dist(p1, points[j])
		}
	}
	return sum
}

func dist(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
