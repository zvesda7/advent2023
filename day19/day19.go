package day19

import (
	"advent23/utils"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	param   int  //0-3 for xmas
	oper    byte //'>' or '<' or 0 for all
	compare int
	target  string
	ptarget *Workflow
}
type Workflow struct {
	name  string
	rules []Rule
}

type Part [4]int

var XMAS = map[byte]int{
	'x': 0,
	'm': 1,
	'a': 2,
	's': 3,
}

func Run() {
	var input, _ = utils.ReadLines("day19/input.txt")
	workflows, parts := parse(input)

	wf_map := map[string]*Workflow{}
	for i := 0; i < len(workflows); i++ {
		wf_map[workflows[i].name] = &workflows[i]
	}
	for i := 0; i < len(workflows); i++ {
		for j := 0; j < len(workflows[i].rules); j++ {
			workflows[i].rules[j].ptarget = wf_map[workflows[i].rules[j].target]
		}
	}

	sum := 0
	for _, p := range parts {
		if accepted(wf_map["in"], p) {
			sum += p[0] + p[1] + p[2] + p[3]
		}
	}
	fmt.Println("Part 1", sum)

	distinct(wf_map)
}

type XMAS_REM [4][]int

func (x *XMAS_REM) copy() XMAS_REM {
	x2 := XMAS_REM{}
	for i := 0; i < 4; i++ {
		for j := 0; j < len((*x)[i]); j++ {
			x2[i] = append(x2[i], (*x)[i][j])
		}
	}
	return x2
}

func distinct(wf_map map[string]*Workflow) {
	xmas_rem := XMAS_REM{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4000; j++ {
			xmas_rem[i] = append(xmas_rem[i], j+1)
		}
	}
	outFill := []XMAS_REM{}
	filter(wf_map["in"], xmas_rem, &outFill)
	tot := 0
	for _, r := range outFill {
		tot += len(r[0]) * len(r[1]) * len(r[2]) * len(r[3])
	}
	fmt.Println("Part 2", tot)
}

func filter(wf *Workflow, xmas_rem XMAS_REM, outFill *[]XMAS_REM) {
	for _, rule := range wf.rules {
		xmas_new := xmas_rem.copy()
		if rule.oper == '>' {
			xmas_new[rule.param] = []int{}
			for i := len(xmas_rem[rule.param]) - 1; i >= 0; i-- {
				if xmas_rem[rule.param][i] > rule.compare {
					xmas_new[rule.param] = append(xmas_new[rule.param], xmas_rem[rule.param][i])
					xmas_rem[rule.param] = append(xmas_rem[rule.param][:i], xmas_rem[rule.param][i+1:]...)
				}
			}
		} else if rule.oper == '<' {
			xmas_new[rule.param] = []int{}
			for i := len(xmas_rem[rule.param]) - 1; i >= 0; i-- {
				if xmas_rem[rule.param][i] < rule.compare {
					xmas_new[rule.param] = append(xmas_new[rule.param], xmas_rem[rule.param][i])
					xmas_rem[rule.param] = append(xmas_rem[rule.param][:i], xmas_rem[rule.param][i+1:]...)
				}
			}
		}
		if rule.target == "A" {
			*outFill = append(*outFill, xmas_new)
		} else if rule.target != "R" {
			filter(rule.ptarget, xmas_new, outFill)
		}
	}
}

func accepted(wf *Workflow, p Part) bool {
	cwf := wf
	for cwf != nil {
		for i := 0; i < len(cwf.rules); i++ {
			if evalRule(&cwf.rules[i], p) {
				if cwf.rules[i].target == "A" {
					return true
				} else if cwf.rules[i].target == "R" {
					return false
				}
				cwf = cwf.rules[i].ptarget
				break
			}
		}
	}

	return false //should never get hit
}

func evalRule(rule *Rule, p Part) bool {
	if rule.oper == '>' {
		return p[rule.param] > rule.compare
	} else if rule.oper == '<' {
		return p[rule.param] < rule.compare
	}
	return true
}

func parse(input []string) ([]Workflow, []Part) {
	i := 0
	workflows := []Workflow{}
	for ; input[i] != ""; i++ {
		split1 := strings.Split(input[i], "{")

		w := Workflow{}
		w.name = split1[0]
		rules_s := split1[1][:len(split1[1])-1]
		for _, rule_s := range strings.Split(rules_s, ",") {
			r := Rule{}
			split2 := strings.Split(rule_s, ":")
			if len(split2) == 1 {
				r.target = split2[0]
			} else {
				r.target = split2[1]
				r.param = XMAS[rule_s[0]]
				r.oper = rule_s[1]
				r.compare, _ = strconv.Atoi(split2[0][2:])
			}
			w.rules = append(w.rules, r)
		}
		workflows = append(workflows, w)
	}

	parts := []Part{}
	for i++; i < len(input); i++ {
		split1 := strings.Split(input[i][1:len(input[i])-1], ",")

		p := Part{}
		for i := 0; i < 4; i++ {
			p[i], _ = strconv.Atoi(split1[i][2:])
		}
		parts = append(parts, p)
	}
	return workflows, parts
}
