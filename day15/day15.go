package day15

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

type HashMapEntry struct {
	key string
	val int
}

type Step struct {
	key  string
	oper byte
	val  int
}

type HashMap [256][]HashMapEntry

func (m *HashMap) Set(key string, val int) {
	box := &(*m)[hash(key)]
	for i := 0; i < len(*box); i++ {
		if (*box)[i].key == key {
			(*box)[i].val = val
			return
		}
	}
	*box = append(*box, HashMapEntry{key, val})
}

func (m *HashMap) Delete(key string) {
	box := &(*m)[hash(key)]
	for i := 0; i < len(*box); i++ {
		if (*box)[i].key == key {
			*box = append((*box)[:i], (*box)[i+1:]...)
			return
		}
	}
}

func Run() {

	var input, _ = utils.ReadLines("day15/input.txt")

	steps := strings.Split(input[0], ",")

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println("part 1", sum)

	hm := HashMap{}
	for _, step := range steps {
		s := parseStep(step)
		if s.oper == '=' {
			hm.Set(s.key, s.val)
		} else {
			hm.Delete(s.key)
		}
	}

	sum = 0
	for i := 0; i < len(hm); i++ {
		for j := 0; j < len(hm[i]); j++ {
			power := (i + 1) * (j + 1) * hm[i][j].val
			sum += power
		}
	}
	fmt.Println("part 2", sum)
}

func parseStep(str string) Step {
	s := Step{}
	split1 := strings.Split(str, "=")
	if len(split1) == 2 {
		s.key = split1[0]
		s.val, _ = strconv.Atoi(split1[1])
		s.oper = '='
	} else {
		s.key = str[:len(str)-1]
		s.oper = '-'
	}
	return s
}

func hash(s string) int {
	cv := 0
	for _, c := range s {
		cv += int(c)
		cv *= 17
		cv = cv % 256
	}
	return cv
}
