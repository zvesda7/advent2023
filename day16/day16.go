package day16

import (
	"advent23/utils"
	"fmt"
	"strings"
)

type Beam struct {
	x  int
	y  int
	dx int
	dy int
}

func (b *Beam) hash() int {
	return b.y*100000 + b.x*100 + b.dx*10 + b.dy
}

func Run() {

	var grid, _ = utils.ReadLines("day16/input.txt")

	fmt.Println("Part 1", calcEnergy(Beam{0, 0, 1, 0}, grid))

	m := 0
	for y := 0; y < len(grid); y++ {
		m = max(m, calcEnergy(Beam{0, y, 1, 0}, grid))
		m = max(m, calcEnergy(Beam{len(grid[y]) - 1, y, -1, 0}, grid))
	}
	for x := 0; x < len(grid[0]); x++ {
		m = max(m, calcEnergy(Beam{x, 0, 0, 1}, grid))
		m = max(m, calcEnergy(Beam{x, len(grid) - 1, 0, -1}, grid))
	}
	fmt.Println("Part 2", m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcEnergy(initial Beam, grid []string) int {

	beams := []Beam{}
	beams = append(beams, initial) //initial, top left moving right

	width := len(grid[0])
	height := len(grid)

	energy := make([][]byte, height)
	for y := 0; y < height; y++ {
		energy[y] = []byte(strings.Repeat(".", width))
	}

	enerCount := 0
	visited := map[int]bool{}
	for len(beams) > 0 {
		for i := len(beams) - 1; i >= 0; i-- {
			//process current
			if energy[beams[i].y][beams[i].x] != '#' {
				enerCount++
			}
			energy[beams[i].y][beams[i].x] = '#'
			newBeam := update(&beams[i], grid[beams[i].y][beams[i].x])
			if newBeam != nil {
				beams = append(beams, *newBeam)
			}
			//move next
			px := beams[i].x + beams[i].dx
			py := beams[i].y + beams[i].dy
			if px < 0 || px >= width || py < 0 || py >= height || visited[beams[i].hash()] {
				beams = append(beams[:i], beams[i+1:]...)
			} else {
				visited[beams[i].hash()] = true
				beams[i].x = px
				beams[i].y = py
			}
		}
	}
	return enerCount
}

var rotations = map[int]map[byte]int{
	1:     {'/': -1000, '\\': 1000, '-': 1, '|': -1000},
	1000:  {'/': -1, '\\': 1, '-': -1, '|': 1000},
	-1:    {'/': 1000, '\\': -1000, '-': -1, '|': -1000},
	-1000: {'/': 1, '\\': -1, '-': -1, '|': -1000},
}

var newBeams = map[int]map[byte]int{
	1:     {'/': 0, '\\': 0, '-': 0, '|': 1000},
	1000:  {'/': 0, '\\': 0, '-': 1, '|': 0},
	-1:    {'/': 0, '\\': 0, '-': 0, '|': 1000},
	-1000: {'/': 0, '\\': 0, '-': 1, '|': 0},
}

func update(beam *Beam, item byte) *Beam {
	if item == '.' {
		return nil
	}
	dir := beam.dy*1000 + beam.dx
	newDir := rotations[dir][item]
	//fmt.Println(*beam, newDir)
	beam.dx = newDir % 1000
	beam.dy = newDir / 1000
	newBeamDir := newBeams[dir][item]
	if newBeamDir > 0 {
		//fmt.Println(newBeamDir)
		newBeam := Beam{beam.x, beam.y, newBeamDir % 1000, newBeamDir / 1000}
		return &newBeam
	}
	return nil
}
