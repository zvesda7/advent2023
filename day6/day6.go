package day6

import (
	"advent23/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time int64
	dist int64
}

func Run() {
	var input, _ = utils.ReadLines("day6/input.txt")
	races, bigRace := parse(input)
	fmt.Println(races, bigRace)

	bestCounts := map[int]int64{}
	for raceI, race := range races {
		bestCounts[raceI] = bestCount(race)
	}
	prod := int64(1)
	for _, cnt := range bestCounts {
		prod *= cnt
	}
	fmt.Println("Part 1", prod)

	fmt.Println("Part 2", bestCount(bigRace))
}

func bestCount(race Race) int64 {
	bestCount := int64(0)
	for ms := int64(1); ms < race.time; ms++ {
		dist := ms * (race.time - ms)
		if dist > race.dist {
			bestCount++
		}
	}
	return bestCount
}

func parse(lines []string) ([]Race, Race) {
	getNums := regexp.MustCompile("[0-9]+")

	times := getNums.FindAllString(lines[0], -1)
	dists := getNums.FindAllString(lines[1], -1)

	races := []Race{}
	for i := 0; i < len(times); i++ {
		race := Race{}
		race.time, _ = strconv.ParseInt(times[i], 10, 64)
		race.dist, _ = strconv.ParseInt(dists[i], 10, 64)
		races = append(races, race)
	}

	bigRace := Race{}
	lines[0] = strings.Replace(lines[0], " ", "", -1)
	lines[1] = strings.Replace(lines[1], " ", "", -1)
	bigRace.time, _ = strconv.ParseInt(getNums.FindAllString(lines[0], -1)[0], 10, 64)
	bigRace.dist, _ = strconv.ParseInt(getNums.FindAllString(lines[1], -1)[0], 10, 64)

	return races, bigRace
}
