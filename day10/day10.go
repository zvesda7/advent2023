package day10

import (
	"advent23/utils"
	"fmt"
)

func Run() {
	var input, _ = utils.ReadLines("day10/input.txt")
	startDir := 10000 //adjust depending where start is
	insideDir := -1   //adjust depending where start is

	points := map[int]byte{}
	startPt := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			points[pos(x, y)] = input[y][x]
			if points[pos(x, y)] == 'S' {
				startPt = pos(x, y)
			}
		}
	}

	fence := map[int]bool{}
	fence[startPt] = true

	cur := startPt + startDir
	prev := startPt
	for cur != startPt {
		fence[cur] = true
		cur, prev, _ = walk(&points, cur, prev, insideDir)
	}
	fmt.Println("Part 1", len(fence)/2)

	//loop again
	enclosed := map[int]bool{}
	cur = startPt + startDir
	prev = startPt
	for cur != startPt {
		floodFill(cur+insideDir, &fence, &points, &enclosed)
		cur, prev, insideDir = walk(&points, cur, prev, insideDir)
		floodFill(prev+insideDir, &fence, &points, &enclosed)
	}
	fmt.Println("Part 2", len(enclosed))
}

func floodFill(point int, fence *map[int]bool, points *map[int]byte, enclosed *map[int]bool) {
	stack := []int{}
	stack = append(stack, point)

	for len(stack) > 0 {
		point = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		_, isFence := (*fence)[point]
		_, isCounted := (*enclosed)[point]
		_, isPoint := (*points)[point]
		if !isFence && !isCounted && isPoint {
			(*enclosed)[point] = true
			for _, d := range dirs {
				stack = append(stack, point+d)
			}
		}
	}

}

var dirs = []int{1, 10000, -1, -10000}

var rules = map[int]map[byte]int{
	-1:     {'F': 10000, '-': -1, 'L': -10000, '|': 0, 'J': 0, '7': 0},
	-10000: {'F': 1, '-': 0, 'L': 0, '|': -10000, 'J': 0, '7': -1},
	1:      {'F': 0, '-': 1, 'L': 0, '|': 0, 'J': -10000, '7': 10000},
	10000:  {'F': 0, '-': 0, 'L': 1, '|': 10000, 'J': -1, '7': 0},
}

var rotations = map[int]map[byte]int{
	-1:     {'F': -1, '-': 0, 'L': 1, '|': 0, 'J': 0, '7': 0},
	-10000: {'F': 1, '-': 0, 'L': 0, '|': 0, 'J': 0, '7': -1},
	1:      {'F': 0, '-': 0, 'L': 0, '|': 0, 'J': -1, '7': 1},
	10000:  {'F': 0, '-': 0, 'L': -1, '|': 0, 'J': 1, '7': 0},
}

var rtransform = map[int]map[int]int{
	-1:     {-1: 10000, 0: -1, 1: -10000},
	-10000: {-1: -1, 0: -10000, 1: 1},
	1:      {-1: -10000, 0: 1, 1: 10000},
	10000:  {-1: 1, 0: 10000, 1: -1},
}

func walk(points *map[int]byte, cur int, prev int, insideDir int) (int, int, int) {
	newPos := cur + rules[cur-prev][(*points)[cur]]
	newRot := rotations[cur-prev][(*points)[cur]]
	insideDir = rtransform[insideDir][newRot]

	return newPos, cur, insideDir
}

func pos(x, y int) int {
	return y*10000 + x
}
