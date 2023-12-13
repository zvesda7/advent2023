package day13

import (
	"advent23/utils"
	"fmt"
)

type Board struct {
	width  int
	height int
	data   [][]byte
}

func Run() {

	var input, _ = utils.ReadLines("day13/input.txt")
	boards := parse(input)

	sum := 0
	for _, b := range boards {
		vert := findVerticalReflect(b, 0)
		horz := findHorizontalReflect(b, 0)
		sum += vert + horz*100
	}
	fmt.Println("part1", sum)

	sum = 0
	for _, b := range boards {
		oldVert := findVerticalReflect(b, 0)
		oldHorz := findHorizontalReflect(b, 0)

		done := false
		for i := 0; i < b.height && !done; i++ {
			for j := 0; j < b.width && !done; j++ {
				b.data[i][j] = invert(b.data[i][j])
				newVert := findVerticalReflect(b, oldVert)
				newHorz := findHorizontalReflect(b, oldHorz)
				b.data[i][j] = invert(b.data[i][j])
				if newVert > 0 || newHorz > 0 {
					sum += newVert + newHorz*100
					done = true
				}
			}
		}

	}
	fmt.Println("part2", sum)
}

func invert(b byte) byte {
	if b == '#' {
		return '.'
	} else {
		return '#'
	}
}

func findVerticalReflect(board Board, reject int) int {
	for i := 1; i < board.width; i++ {
		mirror := true
		for a, b := i-1, i; a >= 0 && b < board.width; a, b = a-1, b+1 {
			for r := 0; r < board.height; r++ {
				if board.data[r][a] != board.data[r][b] {
					mirror = false
					break
				}
			}
		}
		if mirror && i != reject {
			return i
		}
	}
	return 0
}

func findHorizontalReflect(board Board, reject int) int {
	for i := 1; i < board.height; i++ {
		mirror := true
		for a, b := i-1, i; a >= 0 && b < board.height; a, b = a-1, b+1 {
			for r := 0; r < board.width; r++ {
				if board.data[a][r] != board.data[b][r] {
					mirror = false
					break
				}
			}
		}
		if mirror && i != reject {
			return i
		}
	}
	return 0
}

func parse(input []string) []Board {
	cboards := []Board{}
	cboard := Board{}
	for _, x := range input {
		if len(x) == 0 {
			cboards = append(cboards, cboard)
			cboard = Board{}
		} else {
			cboard.data = append(cboard.data, []byte(x))
		}
	}
	cboards = append(cboards, cboard)

	for i := 0; i < len(cboards); i++ {
		cboards[i].height = len(cboards[i].data)
		cboards[i].width = len(cboards[i].data[0])
	}

	return cboards
}
