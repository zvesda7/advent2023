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

type Periodic struct {
	period        int
	leadIn        []int
	periodOffsets []int
}

func Run() {
	var input, _ = utils.ReadLines("day21/input.txt")
	part1(input)
	part2(input)
}

func part1(input []string) {

	steps := 64
	walls, start, width, height := parse(input, 1)
	printGrid(walls, width, height, map[Point]int{}, steps)

	cnt, _, _ := countCells(start, walls, steps, width, height)

	fmt.Println("Part 1", cnt)
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func calc(step int, row int, periodics map[int]Periodic) int {
	period := periodics[row]
	total := 0
	if step < len(period.leadIn) {
		total += sum(period.leadIn[:step+1])
		step = 0
	} else {
		total += sum(period.leadIn)
		step -= len(period.leadIn) - 1
	}

	full := step / period.period
	partial := step % period.period
	total += full * sum(period.periodOffsets)

	if partial > 0 {
		total += sum(period.periodOffsets[:partial])
	}
	return total
}

func findPeriodic(nums []int, period int) Periodic {
	p := Periodic{}

	diffs := []int{}
	for i := 0; i < len(nums); i++ {
		pre := 0
		if i != 0 {
			pre = nums[i-1]
		}
		diffs = append(diffs, nums[i]-pre)
	}

	cnt := 0
	for i := 0; i < len(diffs)-period; i++ {
		if diffs[i] != diffs[i+period] {
			cnt = 0
		} else {
			cnt++
		}
		if cnt == period {
			p.period = period
			p.leadIn = diffs[:i-period+1]
			p.periodOffsets = diffs[i-period+1 : i+1]
			return p
		}
	}

	return p
}

func part2(input []string) {
	periodics := map[int]Periodic{}

	per := len(input)
	walls, start, w, h := parse(input, 5)
	seqs := map[int][]int{}
	for i := 0; i <= per*3; i += 1 {
		fmt.Println(i, per*3)
		_, _, distances := countCells(start, walls, i, w, h)

		plotCounts := map[int]int{}
		for point, d := range distances {
			if d <= i && d%2 == (i%2) {
				plotCounts[point.y-start.y]++
			}
		}

		for i := -per; i <= per; i++ {
			seqs[i] = append(seqs[i], plotCounts[i])
		}
	}

	for i := -per; i <= per; i++ {
		periodic := findPeriodic(seqs[i], per)
		periodics[i] = periodic
	}

	steps := 26501365
	sum := 0
	for row := -steps; row <= steps; row++ {
		segment := row / per
		modRow := row % per
		calcSteps := steps - per*abs(segment)
		sum += calc(calcSteps, modRow, periodics)
		if row%10000 == 0 {
			fmt.Println(row, sum)
		}
	}
	fmt.Printf("%v,%v\n", steps, sum)

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
