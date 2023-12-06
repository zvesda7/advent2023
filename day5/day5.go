package day5

import (
	"advent23/utils"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Route struct {
	start  int64
	end    int64
	offset int64
}

func Run() {
	var input, _ = utils.ReadLines("day5/test.txt")
	seeds, maps := parse(input)

	lowest := int64(math.MaxInt64)
	for _, seed := range seeds {
		val := route(seed, maps)
		if val < lowest {
			lowest = val
		}
	}
	fmt.Println(maps)
	slices.Reverse(maps)
	for _, mp := range maps {
		sortRoutes(&mp)
		for _, rt := range mp {
			fmt.Println(rt.start, rt.end, rt.offset)
		}
		fmt.Println()
	}
	fmt.Println(lowest)
}

func sortRoutes(r *[]Route) {
	sort.Slice(*r, func(i, j int) bool {
		return (*r)[i].start < (*r)[j].end
	})
}

func reduce(maps [][]Route) []Route {
	final := maps[0]
	for i := 1; i < len(maps); i++ {
		final = reducePair(final, maps[i])
	}
	return final
}

func reducePair(a []Route, b []Route) []Route {
	final := []Route{}
	i, j := 0, 0
	a, b := Route{}, Route{}
	for i < len(a) && j < len(b) {
		a, b = a[i], b[j]
		if a.start < b.start && a.end < b.end {
			final = append(final, a)
			i++
		} else if b.start < a.start && b.end < a.end {
			final = append(final, b)
			j++
		} else if  {
		} else if b.start < a.start && b. {
		}
	}
	sortRoutes(&final)
	return final
}

func route(seed int64, maps [][]Route) int64 {
	//find correct route, if any
	cval := seed
	for _, mp := range maps {
		var croute *Route
		for _, route := range mp {
			if cval >= route.start && cval <= route.end {
				croute = &route
				break
			}
		}
		if croute != nil {
			cval += croute.offset
		}
	}
	return cval
}

func parse(input []string) ([]int64, [][]Route) {
	rowPos := 0
	seeds := []int64{}
	seeds_s := strings.Split(input[rowPos], " ")
	for i := 1; i < len(seeds_s); i++ {
		seed, _ := strconv.ParseInt(seeds_s[i], 10, 64)
		seeds = append(seeds, seed)
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
	return seeds, maps
}

//

//56 92 4
//93 96 -37
//
//0 68 1
//69 69 -69
//
//0 55 4
//56 68 5
//69 69 -65
//70 92 4
//93 96 -37





//56 92 4
//93 96 -37
//
//0 68 1
//69 69 -69
//
//45 63 36
//64 76 4
//77 99 -32
//
//18 24 70
//25 94 -7
//
//0 6 42
//7 10 50
//11 52 -11
//53 60 -4
//
//0 14 39
//15 51 -15
//52 53 -15
//
//50 97 2
//98 99 -48