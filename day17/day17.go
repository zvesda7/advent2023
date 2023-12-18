package day17

import (
	"advent23/utils"
	"bufio"
	"fmt"
	"os"
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
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			p := Point{x, y}
			efforts[p.hash()] = int(input[y][x] - '0')
		}
	}

	points := []Point{}
	p0 := Point{0, 0}
	points = append(points, p0)

	dists := map[PHash]int{}
	dists[p0.hash()] = 0

	prevP := map[PHash]Point{}
	prevP[p0.hash()] = Point{-1, -1}

	for len(points) > 0 {

		adj := neighbors(points[0], width, height)
		//p1 := prevP[points[0].hash()]
		//p2 := prevP[p1.hash()]
		for j := 0; j < len(adj); j++ {
			if _, found := dists[adj[j].hash()]; !found {
				if allowed(points[0], adj[j], dists, efforts, width, height, 0) {
					dists[adj[j].hash()] = dists[points[0].hash()] + efforts[adj[j].hash()]
					points = append(points, adj[j])
					//prevP[adj[j].hash()] = points[0]
				}
			}
		}
		points = points[1:]
		sort.Slice(points, func(i, j int) bool {
			return dists[points[i].hash()] < dists[points[j].hash()]
		})
		print(width, height, dists, efforts)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	path := map[PHash]bool{}
	pp := Point{width - 1, height - 1}
	for !(pp.x == 0 && pp.y == 0) {
		path[pp.hash()] = true
		pp = prevP[pp.hash()]
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//p := Point{x, y}
			//if _, ok := path[p.hash()]; ok {
			//	fmt.Printf("\033[32m")
			//}
			//fmt.Printf("[%3d]", dists[p.hash()])
			//fmt.Printf("\033[0m")
		}
		//fmt.Printf("\n")
	}
}

func allowed(s Point, d Point, dists map[PHash]int, efforts map[PHash]int, width int, height int, depth int) bool {
	adj := neighbors(s, width, height)
	tadj := []Point{}
	sDist := dists[s.hash()]
	sEff := efforts[s.hash()]
	for i := 0; i < len(adj); i++ {
		tDist := dists[adj[i].hash()]
		if tDist+sEff == sDist {
			tadj = append(tadj, adj[i])
		}
	}
	if len(tadj) == 0 {
		return true
	}
	for i := 0; i < len(tadj); i++ {
		if !inRow([]Point{tadj[i], s, d}) {
			return true
		}
	}
	if depth == 0 {
		return allowed(tadj[0], s, dists, efforts, width, height, depth+1)
	} else {
		return false
	}
}

func print(width int, height int, dists map[PHash]int, efforts map[PHash]int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			if _, ok := dists[p.hash()]; ok {
				fmt.Printf("\033[32m")
				fmt.Printf("[%3d]", dists[p.hash()])
			} else {
				fmt.Printf("\033[0m")
				fmt.Printf("[%3d]", efforts[p.hash()])
			}
		}
		fmt.Printf("\n")
	}
}

func inRow(points []Point) bool {
	xSame := true
	ySame := true
	x := points[0].x
	y := points[0].y
	for _, p := range points {
		if p.x != x {
			xSame = false
		}
		if p.y != y {
			ySame = false
		}
	}
	return xSame || ySame
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
