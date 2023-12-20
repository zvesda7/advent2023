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

type Dir int

const (
	UP    Dir = 0
	RIGHT Dir = 1
	DOWN  Dir = 2
	LEFT  Dir = 3
)

type Node struct {
	point Point
	dir   Dir
	inRow int
	cost  int
	index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost > pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type NHash int

func (n *Node) hash() NHash {
	return NHash(int(n.point.hash())*1000 + int(n.dir)*100 + n.inRow)
}

func Run() {
	runPart(false)
	runPart(true)
}

func runPart(part2 bool) {
	var input, _ = utils.ReadLines("day17/input.txt")

	width := len(input[0])
	height := len(input)
	efforts := map[PHash]int{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			p := Point{x, y}
			efforts[p.hash()] = int(input[y][x] - '0')
		}
	}
	targ := Point{width - 1, height - 1}

	carts := []Node{{Point{0, 0}, DOWN, 0, 0, 0}, {Point{0, 0}, RIGHT, 0, 0, 0}}

	//bestCost := map[NHash]int{}
	//bestCost[carts[0].hash()] = 0
	//bestCost[carts[1].hash()] = 0
	visited := map[NHash]bool{}
	parents := map[NHash]Node{}

	lc := 0
	for len(carts) > 0 {
		if lc++; lc%1000 == 0 {
			fmt.Println(lc, len(carts), carts[0])
		}
		if carts[0].point.x == targ.x && carts[0].point.y == targ.y {
			if !part2 || carts[0].inRow >= 4 {
				fmt.Println(lc, carts[0])
				//printRoute(carts[0], width, height, efforts, parents, bestCost)
				break
			}
		}

		adj := neighbors(&carts[0], width, height, part2)
		for j := 0; j < len(adj); j++ {
			if !visited[adj[j].hash()] {
				adj[j].cost = carts[0].cost + efforts[adj[j].point.hash()]
				visited[adj[j].hash()] = true
				parents[adj[j].hash()] = carts[0]
				carts = append(carts, adj[j])
			}
		}

		carts = carts[1:]
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].cost < carts[j].cost
		})
	}
}

func pointAddDir(p Point, d Dir) Point {
	p2 := p
	if d < 0 {
		d = 3
	} else if d > 3 {
		d = 0
	}
	if d == UP {
		p2.y += -1
	} else if d == RIGHT {
		p2.x += 1
	} else if d == DOWN {
		p2.y += 1
	} else { //left or up-1 (-1)
		p2.x -= 1
	}
	return p2
}

func inBounds(p Point, width int, height int) bool {
	if p.x < 0 || p.x >= width || p.y < 0 || p.y >= height {
		return false
	}
	return true
}

func clockDir(d Dir) Dir {
	if d == 3 {
		return 0
	} else {
		return d + 1
	}
}
func cclockDir(d Dir) Dir {
	if d == 0 {
		return 3
	} else {
		return d - 1
	}
}

func neighbors(n *Node, width int, height int, part2 bool) []Node {
	ns := []Node{}

	maxInRow := 3
	if part2 {
		maxInRow = 10
	}

	pForw := pointAddDir(n.point, n.dir)
	if inBounds(pForw, width, height) && n.inRow < maxInRow {
		ns = append(ns, Node{pForw, n.dir, n.inRow + 1, 0, 0})
	}
	if !part2 || n.inRow >= 4 {
		pLeft := pointAddDir(n.point, cclockDir(n.dir))
		if inBounds(pLeft, width, height) {
			ns = append(ns, Node{pLeft, cclockDir(n.dir), 1, 0, 0})
		}
		pRight := pointAddDir(n.point, clockDir(n.dir))
		if inBounds(pRight, width, height) {
			ns = append(ns, Node{pRight, clockDir(n.dir), 1, 0, 0})
		}
	}

	return ns
}

func printRoute(n Node, width int, height int, efforts map[PHash]int, parents map[NHash]Node, bestCost map[NHash]int) {
	path := map[PHash]int{}

	for n.point.x != 0 || n.point.y != 0 {
		path[n.point.hash()] = bestCost[NHash(n.hash())]
		n = parents[n.hash()]
	}
	fmt.Println(n)

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
