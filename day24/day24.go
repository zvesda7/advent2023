package day24

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

type Hail struct {
	x  float64
	y  float64
	dx float64
	dy float64
}

func Run() {
	var input, _ = utils.ReadLines("day24/input.txt")
	hails := parse(input)
	fmt.Println(hails)

	//min := float64(7)
	//max := float64(27)
	min := float64(200000000000000)
	max := float64(400000000000000)

	cnt := float64(0)
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			intersects, future, x, y := calcIntersect(hails[i], hails[j])
			fmt.Println(intersects, future, x, y)
			if intersects && future && x >= min && y >= min && x <= max && y <= max {
				cnt++
			}
		}
	}
	fmt.Println("Part 1", cnt)
}

func calcIntersect(h1 Hail, h2 Hail) (bool, bool, float64, float64) {
	m1 := h1.dy / h1.dx
	m2 := h2.dy / h2.dx
	b1 := h1.y - m1*h1.x
	b2 := h2.y - m2*h2.x
	if m1 == m2 {
		return false, false, 0, 0
	}
	x := (b2 - b1) / (m1 - m2)
	y := m1*x + b1
	future := true
	if x > h1.x && h1.dx < 0 {
		future = false
	} else if x < h1.x && h1.dx > 0 {
		future = false
	} else if x > h2.x && h2.dx < 0 {
		future = false
	} else if x < h2.x && h2.dx > 0 {
		future = false
	}
	return true, future, x, y
}

func parse(input []string) []Hail {
	hails := []Hail{}
	for _, r := range input {
		r = strings.Replace(r, " @ ", ", ", -1)
		splits := strings.Split(r, ", ")
		h := Hail{}
		h.x, _ = strconv.ParseFloat(strings.TrimSpace(splits[0]), 64)
		h.y, _ = strconv.ParseFloat(strings.TrimSpace(splits[1]), 64)
		h.dx, _ = strconv.ParseFloat(strings.TrimSpace(splits[3]), 64)
		h.dy, _ = strconv.ParseFloat(strings.TrimSpace(splits[4]), 64)
		hails = append(hails, h)
	}
	return hails
}
