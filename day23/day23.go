package day23

import (
	"advent23/utils"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Path struct {
	prev   *Path
	points map[Point]bool
	curP   Point
	prevP  Point
	dist   int
}

func Run() {
	var input, _ = utils.ReadLines("day23/input.txt")
	grid, w, h := parse(input)

	start := Point{1, 0}
	end := Point{w - 2, h - 1}
	maxDistSoFar := 0
	maxByPoint := map[Point]int{}

	paths := []*Path{{nil, map[Point]bool{}, start, Point{0, 0}, 0}}
	cnt := 0
	for len(paths) > 0 {
		cnt++
		cPath := paths[len(paths)-1]
		neighbors := getOpenAdjPoints(grid, cPath)
		if len(neighbors) == 0 {
			if cPath.curP == end {
				if cPath.dist > maxDistSoFar {
					maxDistSoFar = cPath.dist
				}
			}
			//remove
			paths = paths[:len(paths)-1]
		} else if len(neighbors) == 1 {
			cPath.points[cPath.curP] = true
			cPath.prevP = cPath.curP
			cPath.curP = neighbors[0]
			cPath.dist++
		} else if len(neighbors) > 1 {
			maxByPoint[cPath.curP] = cPath.dist
			cPath.points[cPath.curP] = true
			paths = paths[:len(paths)-1]
			for i := 0; i < len(neighbors); i++ {
				newPath := &Path{cPath, map[Point]bool{}, neighbors[i], cPath.curP, cPath.dist + 1}
				paths = append(paths, newPath)
			}
		}
		//printPath(grid, w, h, cPath)
		if cnt%100000 == 0 {
			fmt.Println("Max", cnt, maxDistSoFar, len(paths))
		}
	}

	fmt.Println("Answer", maxDistSoFar, cnt)
}

var dirs = [4]Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func getOpenAdjPoints(grid map[Point]byte, p *Path) []Point {
	newPoints := []Point{}
	for _, d := range dirs {
		test := Point{p.curP.x + d.x, p.curP.y + d.y}
		if cell, ok := grid[test]; ok {
			if moveAllowed(cell, d, true) && test != p.prevP && !isAlreadyOnPath(p, test) {
				newPoints = append(newPoints, test)
			}
		}
	}
	return newPoints
}

func moveAllowed(cell byte, dir Point, part2 bool) bool {
	if cell == '.' {
		return true
	} else if cell == '>' && (part2 || dir.x == 1) {
		return true
	} else if cell == 'v' && (part2 || dir.y == 1) {
		return true
	} else if cell == '<' && (part2 || dir.x == -1) {
		return true
	} else if cell == '^' && (part2 || dir.y == -1) {
		return true
	}
	return false
}

func isAlreadyOnPath(p *Path, point Point) bool {
	for p != nil {
		if _, ok := p.points[point]; ok {
			return true
		}
		p = p.prev
	}
	return false
}

func parse(input []string) (map[Point]byte, int, int) {
	grid := map[Point]byte{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			grid[Point{x, y}] = input[y][x]
		}
	}
	return grid, len(input[0]), len(input)
}

func printPath(grid map[Point]byte, w int, h int, p *Path) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if isAlreadyOnPath(p, Point{x, y}) {
				fmt.Printf("O")
			} else if cell, ok := grid[Point{x, y}]; ok {
				fmt.Printf(string(cell))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
