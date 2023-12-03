package day2

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

type Draw struct {
	red   int
	green int
	blue  int
}

type Game struct {
	gameNum int
	draws   []Draw
}

func Run() {
	games := parse("day2/input.txt")

	sumIds := 0
	sumPowers := 0
	for _, game := range games {
		if possible(game, 12, 13, 14) {
			sumIds += game.gameNum
		}
		sumPowers += power(game)
	}
	fmt.Println("Answer 1", sumIds)
	fmt.Println("Answer 2", sumPowers)
}

func possible(game Game, maxRed int, maxGreen int, maxBlue int) bool {
	for _, draw := range game.draws {
		if draw.red > maxRed {
			return false
		}
		if draw.green > maxGreen {
			return false
		}
		if draw.blue > maxBlue {
			return false
		}
	}
	return true
}

func power(game Game) int {
	maxred := 0
	maxgreen := 0
	maxblue := 0
	for _, draw := range game.draws {
		if draw.red > maxred {
			maxred = draw.red
		}
		if draw.green > maxgreen {
			maxgreen = draw.green
		}
		if draw.blue > maxblue {
			maxblue = draw.blue
		}
	}
	return maxred * maxgreen * maxblue
}

func parse(file string) []Game {
	var instr, _ = utils.ReadLines(file)
	var games []Game

	for _, row_s := range instr {
		game := Game{}
		split1 := strings.Split(row_s, ":")
		game.gameNum, _ = strconv.Atoi(split1[0][5:])
		for _, draw_s := range strings.Split(split1[1], ";") {
			draw := Draw{}
			for _, cnts_s := range strings.Split(draw_s, ",") {
				split2 := strings.Split(cnts_s, " ")
				cnt, _ := strconv.Atoi(split2[1])
				if split2[2] == "red" {
					draw.red = cnt
				} else if split2[2] == "green" {
					draw.green = cnt
				} else {
					draw.blue = cnt
				}
			}
			//fmt.Println(draw.red, draw.green, draw.blue)
			game.draws = append(game.draws, draw)
		}
		games = append(games, game)
	}

	return games
}
