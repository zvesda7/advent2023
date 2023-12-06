package day5

import (
	"advent23/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Route struct {
	start  int64
	end    int64
	offset int64
}

type Range struct {
	start int64
	count int64
}

func Run() {
	var input, _ = utils.ReadLines("day5/input.txt")
	seeds, seedRanges, maps := parse(input)

	x := reduce(maps)

	lowest := int64(math.MaxInt64)
	for _, seed := range seeds {
		val := route2(seed, x)
		if val < lowest {
			lowest = val
		}
	}

	fmt.Println("Part1", lowest)

	fakeRoutesForSeeds := []Route{}
	for _, r := range seedRanges {
		fakeRoutesForSeeds = append(fakeRoutesForSeeds, Route{r.start, r.start + r.count - 1, 0})
	}
	x = reducePair(fakeRoutesForSeeds, x, true)

	lowest = int64(math.MaxInt64)
	for _, route := range x {

		val := route.start + route.offset
		if val < lowest {
			lowest = val
		}
	}
	fmt.Println("Part2", lowest)
}

func sortRoutes(r *[]Route) {
	sort.Slice(*r, func(i, j int) bool {
		return (*r)[i].start < (*r)[j].start
	})
}

func sortRoutesByDest(r *[]Route) {
	sort.Slice(*r, func(i, j int) bool {
		return (*r)[i].start+(*r)[i].offset < (*r)[j].start+(*r)[j].offset
	})
}

func reduce(maps [][]Route) []Route {
	final := maps[0]
	for i := 1; i < len(maps); i++ {
		final = reducePair(final, maps[i], false)
	}
	return final
}
func reducePair(x []Route, y []Route, noYOnly bool) []Route {
	sortRoutesByDest(&x)
	sortRoutes(&y)

	final := []Route{}
	i, j := 0, 0
	a, b := Route{}, Route{}
	for i < len(x) && j < len(y) {
		a, b = x[i], y[j]
		adstart := a.start + a.offset
		adend := a.end + a.offset

		if adstart < b.start && adend < b.start {
			final = append(final, a)
			i++
		} else if b.start < adstart && b.end < adstart {
			if !noYOnly {
				final = append(final, b)
			}
			j++
		} else {
			//overlap

			final = append(final, Route{max(adstart, b.start) - a.offset, min(adend, b.end) - a.offset, a.offset + b.offset})
			if adstart < b.start {
				final = append(final, Route{a.start, b.start - 1 - a.offset, a.offset})
			} else if b.start < adstart && !noYOnly {
				final = append(final, Route{b.start, adstart - 1, b.offset})
			}
			if adend > b.end {
				x[i].start = b.end + 1 - a.offset
				j++
			} else if b.end > adend {
				y[j].start = adend + 1
				i++
			} else {
				i++
				j++
			}

		}
	}
	for ; i < len(x); i++ {
		final = append(final, x[i])
	}
	if !noYOnly {
		for ; j < len(y); j++ {
			final = append(final, y[j])
		}
	}
	//fmt.Println(final)
	return final
}

func route2(seed int64, maps []Route) int64 {
	//find correct route, if any
	for _, route := range maps {
		if seed >= route.start && seed <= route.end {
			return seed + route.offset
		}
	}

	return 0
}

func parse(input []string) ([]int64, []Range, [][]Route) {
	rowPos := 0
	seeds := []int64{}
	seeds_s := strings.Split(input[rowPos], " ")
	for i := 1; i < len(seeds_s); i++ {
		seed, _ := strconv.ParseInt(seeds_s[i], 10, 64)
		seeds = append(seeds, seed)
	}
	seedRanges := []Range{}
	for i := 1; i < len(seeds_s); i += 2 {
		seedRange := Range{}
		seedRange.start, _ = strconv.ParseInt(seeds_s[i], 10, 64)
		seedRange.count, _ = strconv.ParseInt(seeds_s[i+1], 10, 64)
		seedRanges = append(seedRanges, seedRange)
	}
	rowPos += 1
	maps := [][]Route{}
	for rowPos < len(input) {
		rowPos += 2
		mp := []Route{}
		for ; rowPos < len(input) && len(input[rowPos]) > 0; rowPos++ {
			route_s := strings.Split(input[rowPos], " ")
			dest, _ := strconv.ParseInt(route_s[0], 10, 64)
			src, _ := strconv.ParseInt(route_s[1], 10, 64)
			count, _ := strconv.ParseInt(route_s[2], 10, 64)
			mp = append(mp, Route{src, src + count - 1, dest - src})
		}
		maps = append(maps, mp)
	}
	return seeds, seedRanges, maps
}
