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
	var input, _ = utils.ReadLines("day21/input.txt")
	//part1(input)
	part2(input)
}

func part1(input []string) {

	steps := 6
	walls, start, width, height := parse(input, 1)
	printGrid(walls, width, height, map[Point]int{})

	cnt := countCells(start, walls, steps)

	fmt.Println("Part 1", cnt)
}

func part2(input []string) {
	maxC := 500
	l := (len(input) - 1) / 2
	sc := maxC / l
	scale := sc*2 + 1
	walls, start, _, _ := parse(input, scale)

	nums := []int{6, 10, 50, 100, 500}
	for i := 0; i < len(nums); i++ {
		cnt := countCells(start, walls, nums[i])
		fmt.Println(nums[i], cnt)
	}
	//for i := 1; i <= 262; i++ {
	//	cnt := countCells(start, walls, i)
	//	fmt.Println(i, cnt)
	//}
	//fmt.Println()
	//for i := 1; i <= 10; i++ {
	//	cnt := countCells(start, walls, 65*i)
	//	fmt.Println(65*i, cnt)
	//}

}

func countCells(start Point, walls map[Point]bool, steps int) int {
	distances := map[Point]int{}
	distances[start] = 0

	set := []Node{{start, 0}}
	for len(set) > 0 {
		for _, n := range getNeighbors(walls, set[0]) {
			if _, found := distances[n.point]; !found && n.dist <= steps {
				distances[n.point] = n.dist
				set = append(set, n)
			}
		}

		set = set[1:]
		sort.Slice(set, func(i, j int) bool {
			return set[i].dist < set[i].dist
		})
	}

	cnt := 0
	for _, d := range distances {
		if d <= steps && d%2 == 0 {
			cnt++
		}
	}
	return cnt
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

func parse(input []string, scale int) (map[Point]bool, Point, int, int) {
	walls := map[Point]bool{}
	start := Point{}
	width := 0
	height := 0
	if scale == 1 {
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
		width = len(input[0]) + 2
		height = len(input) + 2
		for x := 0; x < width; x++ {
			walls[Point{x, 0}] = true
			walls[Point{x, height - 1}] = true
		}
		for y := 1; y < height-1; y++ {
			walls[Point{0, y}] = true
			walls[Point{width - 1, y}] = true
		}
	} else {

		width = len(input[0])
		height = len(input)
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[0]); x++ {
				char := input[y][x]
				if char == '#' {
					for a := 0; a < scale; a++ {
						for b := 0; b < scale; b++ {
							walls[Point{x + width*a, y + height*b}] = true
						}
					}
				} else if char == 'S' {
					start = Point{x + width*(scale/2), y + height*(scale/2)}
				}
			}
		}
	}

	return walls, start, width * scale, height * scale
}

func printGrid(walls map[Point]bool, width int, height int, distances map[Point]int) {

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, isWall := walls[Point{x, y}]
			_, isReach := distances[Point{x, y}]
			if isWall {
				fmt.Printf("#")
			} else if isReach {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
