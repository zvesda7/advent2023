package day17

import (
	"advent23/utils"
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
}

type PHash int

func (p *Point) hash() PHash {
	return PHash(p.y*1000 + p.x)
}

func Run() {

	var input, _ = utils.ReadLines("day17/test.txt")

	width := len(input[0])
	height := len(input)
	efforts := map[PHash]int{}
	dists := map[PHash]int{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			p := Point{x, y}
			efforts[p.hash()] = int(input[y][x] - '0')
		}
	}

	points := []Point{}
	p0 := Point{0, 0}
	dists[p0.hash()] = 0
	points = append(points, p0)
	for len(points) > 0 {
		newPoints := []Point{}
		for i := 0; i < len(points); i++ {
			adj := neighbors(points[i], width, height)
			for j := 0; j < len(adj); j++ {
				if _, found := dists[adj[j].hash()]; !found {
					dists[adj[j].hash()] = dists[points[i].hash()] + efforts[adj[j].hash()]
					newPoints = append(newPoints, adj[j])
				}
			}
		}
		points = newPoints
		sort.Slice(points, func(i, j int) bool {
			return dists[points[i].hash()] < dists[points[j].hash()]
		})
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			fmt.Printf("[%3d]", dists[p.hash()])
		}
		fmt.Printf("\n")
	}
}

func neighbors(p Point, width int, height int) []Point {
	ps := []Point{}
	if p.x > 0 {
		ps = append(ps, Point{p.x - 1, p.y})
	}
	if p.x < width-1 {
		ps = append(ps, Point{p.x + 1, p.y})
	}
	if p.y > 0 {
		ps = append(ps, Point{p.x, p.y - 1})
	}
	if p.y < height-1 {
		ps = append(ps, Point{p.x, p.y + 1})
	}
	return ps
}
