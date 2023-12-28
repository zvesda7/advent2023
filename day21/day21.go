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
	var input, _ = utils.ReadLines("day21/test2.txt")
	//part1(input)
	part2n(input)
}

func part1(input []string) {

	steps := 6
	walls, start, width, height := parse(input, 1)
	printGrid(walls, width, height, map[Point]int{}, steps)

	cnt, _, _ := countCells(start, walls, steps, width, height)

	fmt.Println("Part 1", cnt)
}

type PlotCount struct {
	rowNum int
	width  int
	left   bool
	odd    bool
}

func part2n(input []string) {
	maxStep := len(input) * 2
	walls, start, w, h := parse(input, 19)
	plotCounts := map[PlotCount]int{}
	for i := 0; i <= maxStep*3; i += 1 {
		_, _, distances := countCells(start, walls, i, w, h)
		for point, d := range distances {
			if d <= i && d%2 == (i%2) {
				width := i - abs(point.y-start.y)
				if point.x != start.x {
					rowNum := point.y - start.y
					left := point.x < start.x
					odd := width%2 == 1
					plotCounts[PlotCount{rowNum, width, left, odd}]++
					if width == maxStep+1 {
						if abs(point.x-start.x) != maxStep+1 {
							plotCounts[PlotCount{rowNum, width - 1, left, odd}]++
						}
					}
				}
			}
		}
	}

	//fmt.Println(plotCounts)
	//steps := 63

	for steps := 5; steps <= 41; steps++ {
		sum := 0
		for i := -steps; i <= steps; i++ {
			rowNum := i % maxStep
			width := steps - abs(i)
			blockWidth := maxStep
			widthMod := width % blockWidth
			widthMult := width / blockWidth

			odd := width%2 == 1

			cntRemLeft := plotCounts[PlotCount{rowNum, widthMod, true, odd}]
			cntMultLeft := widthMult * plotCounts[PlotCount{rowNum, blockWidth, true, odd}]
			cntRemRight := plotCounts[PlotCount{rowNum, widthMod, false, odd}]
			cntMultRight := widthMult * plotCounts[PlotCount{rowNum, blockWidth, false, odd}]
			x0 := 0
			if abs(i)%2 == steps%2 {
				x0 = 1
			}
			s := cntRemLeft + cntMultLeft + cntRemRight + cntMultRight + x0
			fmt.Println(i, cntRemLeft, cntMultLeft, cntRemRight, cntMultRight, x0, s)
			sum += s
		}
		fmt.Println("Part2", steps, sum)
	}
}

func part2(input []string) {
	r := 5
	maxC := 54 * r
	l := (len(input) - 1) / 2
	sc := maxC / l
	scale := sc*2 + 1
	fmt.Println("scale", scale)
	walls, start, w, h := parse(input, 7)

	//nums := []int{6, 10, 50, 100, 500, 1000}
	//for i := 0; i < len(nums); i++ {
	//	cnt := countCells(start, walls, nums[i])
	//	fmt.Println(nums[i], cnt)
	//}

	//nums := []int{2 * r, 6 * r, 18 * r}

	//	for i := r + 1; i < 54*r; i += 2 {
	//		cnt, furth := countCells(start, walls, i, w, h)
	//		fmt.Printf("%v,%v,%v\n", i, cnt, furth)
	//	}

	for i := 5; i <= 62; i += 1 {
		cnt, furth, distances := countCells(start, walls, i, w, h)
		fmt.Printf("%v,%v,%v\n", i, cnt, furth)

		plotCounts := map[int]int{}
		for point, d := range distances {
			if d <= i && d%2 == (i%2) {
				plotCounts[point.y-start.y]++
			}
		}
		for j := -i; j <= i; j++ {
			//fmt.Println(plotCounts[j])
		}

		//fmt.Scanln()
		//fmt.Println(nums[i], cnt)
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

func countCells(start Point, walls map[Point]bool, steps int, w int, h int) (int, int, map[Point]int) {
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
	max := 0
	for px, d := range distances {
		if d <= steps && d%2 == (steps%2) {
			cnt++
			if abs(px.x-start.x) > max {
				max = abs(px.x - start.x)
			}
			if abs(px.y-start.y) > max {
				max = abs(px.y - start.y)
			}
		}
	}
	//printGrid(walls, w, h, distances, steps)
	return cnt, max, distances
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func printGrid(walls map[Point]bool, width int, height int, distances map[Point]int, steps int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, isWall := walls[Point{x, y}]
			d, isReach := distances[Point{x, y}]
			if isWall {
				fmt.Printf("\033[0m")
				fmt.Printf("#")
			} else if isReach && d <= steps && d%2 == (steps%2) {
				fmt.Printf("\033[32m")
				fmt.Printf("x")
			} else {
				fmt.Printf("\033[0m")
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
