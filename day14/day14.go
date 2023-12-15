package day14

import (
	"advent23/utils"
	"fmt"
)

type World struct {
	rocks  map[int]bool
	walls  map[int]bool
	width  int
	height int
}

func Run() {

	var input, _ = utils.ReadLines("day14/input.txt")

	world := parse(input)

	tilt(0, -1, &world)
	fmt.Println("part1", calcLoad(&world))

	sampleSize := 300
	runs := []int{}
	for i := 0; i < sampleSize; i++ {
		tilt(0, -1, &world)
		tilt(-1, 0, &world)
		tilt(0, 1, &world)
		tilt(1, 0, &world)
		load := calcLoad(&world)
		runs = append(runs, load)
	}
	period := period(runs)
	fmt.Println("period", period)

	calc := 1000000000 - sampleSize
	calc = calc % period
	fmt.Println(runs[len(runs)-period+calc-1])
}

func period(runs []int) int {
	for i := len(runs) - 5; i >= 0; i-- {
		fail := false
		for j := 0; j < 5; j++ {
			if runs[i-j] != runs[len(runs)-1-j] {
				fail = true
			}
		}
		if fail == false {
			return len(runs) - i - 1
		}
	}
	return 0
}

func calcLoad(world *World) int {
	sum := 0
	for p := range world.rocks {
		sum += world.height - y(p)
	}
	return sum
}
func tilt(dx int, dy int, world *World) {
	settled := false
	for !settled {
		settled = true
		for p, _ := range world.rocks {
			xTry := x(p) + dx
			yTry := y(p) + dy
			pTry := pos(xTry, yTry)
			_, hasRock := world.rocks[pTry]
			_, hasWall := world.walls[pTry]
			if xTry >= 0 && xTry < world.width && yTry >= 0 && yTry < world.height && !hasRock && !hasWall {
				delete(world.rocks, p)
				world.rocks[pTry] = true
				settled = false
			}
		}
	}
}

func parse(input []string) World {
	w := World{}
	w.rocks = make(map[int]bool)
	w.walls = make(map[int]bool)
	w.height = len(input)
	w.width = len(input[0])
	for y, r := range input {
		for x, c := range r {
			if c == 'O' {
				w.rocks[pos(x, y)] = true
			} else if c == '#' {
				w.walls[pos(x, y)] = true
			}
		}
	}
	return w
}

func pos(x, y int) int {
	return y*10000 + x
}

func x(pos int) int {
	return pos % 10000
}

func y(pos int) int {
	return pos / 10000
}
