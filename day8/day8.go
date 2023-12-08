package day8

import (
	"advent23/utils"
	"fmt"
)

type Node struct {
	id      int
	leftId  int
	rightId int
	left    *Node
	right   *Node
}

func Run() {
	var input, _ = utils.ReadLines("day8/input.txt")
	dirs, nodes := parse(input)

	//build tree
	lookup := map[int]*Node{}
	for i := 0; i < len(nodes); i++ {
		lookup[nodes[i].id] = &nodes[i]
	}

	for _, n := range lookup {
		n.left = lookup[n.leftId]
		n.right = lookup[n.rightId]
	}

	stepI := 0
	stepCount := 0
	cNode := lookup[0]
	for cNode.id != (25*26*26 + 25*26 + 25) {
		if stepI == len(dirs) {
			stepI = 0
		}
		if dirs[stepI] == 1 {
			cNode = cNode.right
		} else {
			cNode = cNode.left
		}
		stepI++
		stepCount++
	}
	fmt.Println("Part 1", stepCount)

	cNodes := []*Node{}
	for _, n := range lookup {
		if n.id%26 == 0 {
			cNodes = append(cNodes, n)
		}
	}
	stepI = 0
	stepCount = 0
	zCountByStartNode := map[int]int{}
	for {
		if stepI == len(dirs) {
			stepI = 0
		}
		for i := 0; i < len(cNodes); i++ {
			if dirs[stepI] == 1 {
				cNodes[i] = cNodes[i].right
			} else {
				cNodes[i] = cNodes[i].left
			}
		}
		stepI++
		stepCount++

		done := false
		for i, cNode := range cNodes {
			if cNode.id%26 == 25 {
				zCountByStartNode[i] = stepCount
				if len(zCountByStartNode) == len(cNodes) {
					done = true
				}
			}
		}
		if done {
			break
		}
	}
	nums := []int{}
	for _, x := range zCountByStartNode {
		nums = append(nums, x)
	}
	fmt.Println("Part 2", LCM(nums[0], nums[1], nums[2:]))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], []int{})
	}

	return result
}

func parse(input []string) ([]int, []Node) {
	dirs := []int{}
	for _, x := range input[0] {
		if x == 'R' {
			dirs = append(dirs, 1)
		} else {
			dirs = append(dirs, -1)
		}
	}

	nodes := []Node{}
	for i := 2; i < len(input); i++ {
		x := [3]string{}
		x[0] = input[i][0:3]
		x[1] = input[i][7:10]
		x[2] = input[i][12:15]
		nodes = append(nodes, Node{parseToNum(x[0]), parseToNum(x[1]), parseToNum(x[2]), nil, nil})
	}
	return dirs, nodes
}

func parseToNum(s string) int {
	return int(s[0]-'A')*26*26 + int(s[1]-'A')*26 + int(s[2]-'A')
}
