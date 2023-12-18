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
	steps  int
}

func Run() {

	var input, _ = utils.ReadLines("day17/test2.txt")

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
	root := Node{nil, p0, 0, 0}
	root2 := Node{&root, p0, 0, 0}
	carts := []Node{root2}

	for len(carts) > 0 {
		if carts[0].point.x == width-1 && carts[0].point.y == 0 {
			for i := 0; i < len(carts); i++ {
				fmt.Println(carts[i].steps, carts[i].cost)
				printRoute(&carts[i], width, height, efforts)
			}
			break
		}
		adj := neighbors(carts[0].point, width, height)
		for j := 0; j < len(adj); j++ {
			already := onPath(adj[j], &carts[0])
			tooMany := tooMany(adj[j], &carts[0])
			if !already && !tooMany {
				n := Node{&carts[0], adj[j], carts[0].cost + efforts[adj[j].hash()], carts[0].steps + 1}
				carts = append(carts, n)
			}
		}
		carts = carts[1:]
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].cost < carts[j].cost
		})
		//for i := 0; i < len(carts); i++ {
		//	fmt.Println(i, carts[i].point, carts[i].cost, carts[i].steps)
		//}
		fmt.Println(len(carts), carts[0].cost, carts[0].steps)
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
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
	for n != nil {
		if p.x == n.point.x && p.y == n.point.y {
			return true
		}
		n = n.parent
	}
	return false
}

func printRoute(n *Node, width int, height int, efforts map[PHash]int) {
	path := map[PHash]bool{}
	for n != nil {
		path[n.point.hash()] = true
		n = n.parent
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			if _, ok := path[p.hash()]; ok {
				fmt.Printf("\033[32m")
			}
			fmt.Printf("[%3d]", efforts[p.hash()])
			fmt.Printf("\033[0m")
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
