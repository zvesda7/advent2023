package day18

import (
	"advent23/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Step struct {
	dir Point
	cnt int
}

type Point struct {
	x int
	y int
}

func Run() {
	var input, _ = utils.ReadLines("day18/test.txt")

	part1(input)
	part2(input)

}

func part2(input []string) {
	steps := parse2(input)

	verts := []Point{}
	curr := Point{0, 0}
	verts = append(verts, curr)
	perim := 0
	for _, step := range steps {
		dx := step.dir.x * step.cnt
		dy := step.dir.y * step.cnt
		perim += abs(dx) + abs(dy)
		curr.x += dx
		curr.y += dy
		verts = append(verts, curr)
	}
	fmt.Println(curr)

	slices.Reverse(verts) //counter clockwise
	fmt.Println(perim)
	fmt.Println("Part 2", shoelace(verts)+perim/2+1)
}

func shoelace(verts []Point) int {
	sum1, sum2 := 0, 0
	for i := 0; i < len(verts)-1; i++ {
		sum1 += verts[i].x * verts[i+1].y
		sum2 += verts[i+1].x * verts[i].y
	}
	sum1 += verts[len(verts)-1].x * verts[0].y
	sum2 += verts[0].x * verts[len(verts)-1].y
	return abs(sum1-sum2) / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1(input []string) {
	steps := parse(input)

	points := map[Point]bool{}

	curr := Point{0, 0}
	points[curr] = true
	for _, step := range steps {
		for i := 0; i < step.cnt; i++ {
			curr.x += step.dir.x
			curr.y += step.dir.y
			points[curr] = true
		}
	}

	n := []Point{{1, 1}}
	for len(n) > 0 {
		for _, dir := range DIR_LU {
			test := Point{n[0].x + dir.x, n[0].y + dir.y}
			if _, ok := points[test]; !ok {
				points[test] = true
				n = append(n, test)
			}
		}

		n = n[1:]
	}

	fmt.Println("Part 1", len(points))
}

var DIR_LU = map[byte]Point{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

var DIR2_LU = map[byte]Point{
	'3': {0, -1},
	'0': {1, 0},
	'1': {0, 1},
	'2': {-1, 0},
}

func parse(input []string) []Step {
	steps := []Step{}
	for _, r := range input {
		split := strings.Split(r, " ")
		s := Step{}
		s.dir = DIR_LU[r[0]]
		s.cnt, _ = strconv.Atoi(split[1])
		steps = append(steps, s)
	}
	return steps
}

func parse2(input []string) []Step {
	steps := []Step{}
	for _, r := range input {
		split := strings.Split(r, " ")
		s := Step{}
		s.dir = DIR2_LU[split[2][7]]
		cnt64, _ := strconv.ParseInt(split[2][2:7], 16, 32)
		s.cnt = int(cnt64)
		steps = append(steps, s)
	}
	return steps
}
