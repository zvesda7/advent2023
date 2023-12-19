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

type Node struct {
	parent *Node
	point  Point
	cost   int
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

	p0 := Point{0, 0}
	px := Point{-1, 0}
	py := Point{0, -1}
	rootx := Node{nil, px, 0}
	rooty := Node{nil, py, 0}
	root1 := Node{&rootx, p0, 0}
	root2 := Node{&rooty, p0, 0}
	targ := Point{width - 1, height - 1}

	carts := []Node{root1, root2}

	bestToNode := map[int]int{}

	lc := 0
	for len(carts) > 0 {
		if lc++; lc%1000 == 0 {
			fmt.Println(lc, len(carts), carts[0].cost)
		}
		if carts[0].point.x == targ.x && carts[0].point.y == targ.y {
			fmt.Println(lc, carts[0].cost)
			printRoute(&carts[0], width, height, efforts)
			break
		}

		adj := neighbors(carts[0].point, width, height)
		for j := 0; j < len(adj); j++ {
			cost := carts[0].cost + efforts[adj[j].hash()]
			already := onPath(adj[j], &carts[0])
			tooMany := tooMany(adj[j], &carts[0])
			if !already && !tooMany {
				if isBest(adj[j], &carts[0], cost, bestToNode) {
					n := Node{&carts[0], adj[j], cost}
					carts = append(carts, n)
				}
			}
		}

		carts = carts[1:]
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].cost < carts[j].cost
		})
	}
}

func isBest(p Point, from *Node, cost int, bestMap map[int]int) bool {
	dir := 0
	pt := p
	for i := 0; i < 3 && from != nil; i++ {
		dir *= 10
		if pt.x > from.point.x {
			dir += 1
		} else if pt.y < from.point.y {
			dir += 2
		} else if pt.x < from.point.x {
			dir += 3
		}
		pt = from.point
		from = from.parent
	}
	best, ok := bestMap[int(p.hash())*1000000+dir]
	if !ok || cost < best {
		bestMap[int(p.hash())*1000000+dir] = cost
		return true
	}
	return false
}

func tooMany(p Point, n *Node) bool {
	xDiff := false
	yDiff := false
	x := p.x
	y := p.y
	cnt := 0
	for ; cnt < 4 && n != nil; cnt++ {
		if n.point.x != x {
			xDiff = true
		}
		if n.point.y != y {
			yDiff = true
		}
		if xDiff && yDiff {
			return false
		}
		n = n.parent
	}
	if cnt == 4 {
		return true
	} else {
		return false
	}
}

func onPath(p Point, n *Node) bool {
	if n.parent != nil {
		n = n.parent
		if p.x == n.point.x && p.y == n.point.y {
			return true
		}
	}
	return false
}

func printRoute(n *Node, width int, height int, efforts map[PHash]int) {
	path := map[PHash]int{}
	for n != nil {
		path[n.point.hash()] = n.cost
		n = n.parent
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			if _, ok := path[p.hash()]; ok {
				fmt.Printf("\033[32m")
				fmt.Printf("[%3d]", path[p.hash()])
			} else {
				fmt.Printf("\033[0m")
				fmt.Printf("[%3d]", efforts[p.hash()])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
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
