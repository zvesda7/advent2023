package day21

import (
	"advent23/utils"
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
}

type Node struct {
	point Point
	dist  int
}

func Run() {
	var input, _ = utils.ReadLines("day21/test.txt")
	steps := 6
	walls, start, width, height := parse(input)
	printGrid(walls, width, height, map[Point]int{})

	distances := map[Point]int{}
	distances[start] = 0

	set := []Node{{start, 0}}
	for len(set) > 0 {
		for _, n := range getNeighbors(walls, set[0]) {
			if _, found := distances[n.point]; !found {
				distances[n.point] = n.dist
				set = append(set, n)
			}
		}

		set = set[1:]
		sort.Slice(set, func(i, j int) bool {
			return set[i].dist < set[i].dist
		})
	}
	//printGrid(walls, width, height, distances)

	cnt := 0
	for _, d := range distances {
		if d <= steps && d%2 == 0 {
			cnt++
		}
	}
	fmt.Println("Part 1", cnt)
}

var dirs = [4]Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func getNeighbors(walls map[Point]bool, n Node) []Node {
	neighbors := []Node{}
	for _, d := range dirs {
		test := Node{Point{n.point.x + d.x, n.point.y + d.y}, n.dist + 1}
		if _, hit := walls[test.point]; !hit {
			neighbors = append(neighbors, test)
		}
	}
	return neighbors
}

func parse(input []string) (map[Point]bool, Point, int, int) {
	walls := map[Point]bool{}
	start := Point{}

	for y := 1; y <= len(input); y++ {
		for x := 1; x <= len(input[0]); x++ {
			char := input[y-1][x-1]
			if char == '#' {
				walls[Point{x, y}] = true
			} else if char == 'S' {
				start = Point{x, y}
			}
		}
	}
	width := len(input[0]) + 2
	height := len(input) + 2
	for x := 0; x < width; x++ {
		walls[Point{x, 0}] = true
		walls[Point{x, height - 1}] = true
	}
	for y := 1; y < height-1; y++ {
		walls[Point{0, y}] = true
		walls[Point{width - 1, y}] = true
	}

	return walls, start, width, height
}

func printGrid(walls map[Point]bool, width int, height int, distances map[Point]int) {
	return
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, isWall := walls[Point{x, y}]
			dist, isReach := distances[Point{x, y}]
			if isWall {
				fmt.Printf("##")
			} else if isReach {
				fmt.Printf("%2d", dist)
			} else {
				fmt.Printf("..")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
