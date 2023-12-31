package day22

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	var input, _ = utils.ReadLines("day22/input.txt")

	w := World{}
	w.bricks = parse(input)
	w.points = make(map[Point]int)
	for _, brick := range w.bricks {
		//fmt.Println(brick)
		for _, p := range brick.pts {
			w.points[p] = brick.index
		}
	}

	moved := true
	for moved {
		moved = false
		for _, brick := range w.bricks {
			for w.clearBelow(brick.index) {
				w.moveDown(brick.index)
				moved = true
			}
		}
	}

	cnt := 0
	for _, brick := range w.bricks {
		if w.safeToRemove(brick.index) {
			//fmt.Println("SAFE", brick.index)
			cnt++
		}
	}
	fmt.Println("Part1", cnt)

	cnt = 0
	for _, brick := range w.bricks {
		cnt += w.numThatFall(brick.index)
		//fmt.Println("n", cnt)
	}
	fmt.Println("Part2", cnt)
}

type World struct {
	points map[Point]int
	bricks []Brick
}

func (w *World) clone() World {
	w2 := World{}
	w2.points = make(map[Point]int)
	for k, v := range w.points {
		w2.points[k] = v
	}
	w2.bricks = make([]Brick, len(w.bricks))
	for i, x := range w.bricks {
		w2.bricks[i].index = x.index
		for _, p := range x.pts {
			w2.bricks[i].pts = append(w2.bricks[i].pts, p)
		}
	}
	return w2
}
func (w *World) clearBelow(bIndex int) bool {
	pBrick := w.bricks[bIndex]
	for _, p := range pBrick.pts {
		ptest := Point{p.x, p.y, p.z - 1}
		if bi, ok := w.points[ptest]; (ok && bi != bIndex) || p.z == 1 {
			return false
		}
	}
	return true
}

func (w *World) supported(bIndex int, bIndexExclude int) bool {
	pBrick := w.bricks[bIndex]
	for _, p := range pBrick.pts {
		ptest := Point{p.x, p.y, p.z - 1}
		if bi, ok := w.points[ptest]; ok && bi != bIndex && bi != bIndexExclude {
			return true
		}
	}
	return false
}

func (w *World) moveDown(bIndex int) {
	pBrick := w.bricks[bIndex]
	for i := 0; i < len(pBrick.pts); i++ {
		delete(w.points, pBrick.pts[i])
		pBrick.pts[i].z--
		w.points[pBrick.pts[i]] = pBrick.index
	}
}

func (w *World) safeToRemove(bIndex int) bool {
	pBrick := w.bricks[bIndex]
	for _, p := range pBrick.pts {
		ptest := Point{p.x, p.y, p.z + 1}
		if bi, ok := w.points[ptest]; ok && bi != bIndex {
			if !w.supported(bi, bIndex) {
				return false
			}
		}
	}
	return true
}

func (w *World) numThatFall(bIndex int) int {
	w2 := w.clone()
	pBrick := w2.bricks[bIndex]
	for i := 0; i < len(pBrick.pts); i++ {
		delete(w2.points, pBrick.pts[i])
	}

	moved := true
	bricksMoved := map[int]bool{}
	for moved {
		moved = false
		for _, brick := range w2.bricks {
			for w2.clearBelow(brick.index) {
				w2.moveDown(brick.index)
				moved = true
				bricksMoved[brick.index] = true
			}
		}
	}
	return len(bricksMoved)
}

type Point struct {
	x int
	y int
	z int
}
type Brick struct {
	index int
	pts   []Point
}

func parse(input []string) []Brick {
	bricks := []Brick{}
	for i, row := range input {
		split1 := strings.Split(row, "~")
		splita := strings.Split(split1[0], ",")
		splitb := strings.Split(split1[1], ",")

		a, b := Point{}, Point{}
		a.x, _ = strconv.Atoi(splita[0])
		a.y, _ = strconv.Atoi(splita[1])
		a.z, _ = strconv.Atoi(splita[2])
		b.x, _ = strconv.Atoi(splitb[0])
		b.y, _ = strconv.Atoi(splitb[1])
		b.z, _ = strconv.Atoi(splitb[2])
		bricks = append(bricks, Brick{i, fillBrick(a, b)})
	}
	return bricks
}

func fillBrick(a Point, b Point) []Point {
	brick := []Point{}
	if a.x != b.x {
		for x := a.x; x <= b.x; x++ {
			brick = append(brick, Point{x, a.y, a.z})
		}
	} else if a.y != b.y {
		for y := a.y; y <= b.y; y++ {
			brick = append(brick, Point{a.x, y, a.z})
		}
	} else if a.z != b.z {
		for z := a.z; z <= b.z; z++ {
			brick = append(brick, Point{a.x, a.y, z})
		}
	} else {
		//single point brick
		brick = append(brick, a)
	}
	return brick
}
