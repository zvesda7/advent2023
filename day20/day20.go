package day20

import (
	"advent23/utils"
	"fmt"
	"strings"
)

type Dir int

type Module struct {
	typ        byte // '%' flip flop, '&' conjunction, 'b' broadcast
	name       string
	outputs_s  []string
	inputs     []*Module
	outputs    []*Module
	state_ff   bool
	state_conj map[string]bool //one for each input
}

type Pulse struct {
	from   *Module
	to     *Module
	signal bool
}

type PCount struct {
	low  int
	high int
}

func Run() {
	var input, _ = utils.ReadLines("day20/input.txt")
	modules := parse(input)
	root := buildGraph(modules)
	total := PCount{}
	for i := 0; i < 1000; i++ {
		subTot := pulse(root)
		total.high += subTot.high
		total.low += subTot.low
	}
	fmt.Println("Part1", total.high*total.low)
}

func pulse(root *Module) PCount {
	button := &Module{}
	button.name = "button"
	pulses := []Pulse{{button, root, false}} //initial low pulse into broadcaster aka 'root'
	totalPulses := PCount{}
	for len(pulses) > 0 {
		newPulses := []Pulse{}
		for _, p := range pulses {

			if p.signal {
				totalPulses.high += 1
			} else {
				totalPulses.low += 1
			}
			//fmt.Println(p.from.name, p.signal, p.to.name)
			evalPulse(p, &newPulses)
		}
		pulses = newPulses
		//fmt.Println("")
	}
	return totalPulses
}

func evalPulse(p Pulse, pList *[]Pulse) {
	if p.to.typ == 'b' {
		for _, mod := range p.to.outputs {
			*pList = append(*pList, Pulse{p.to, mod, p.signal})
		}
	} else if p.to.typ == '%' {
		//do nothing unless signal is low
		if p.signal == false {
			p.to.state_ff = !p.to.state_ff
			for _, mod := range p.to.outputs {
				*pList = append(*pList, Pulse{p.to, mod, p.to.state_ff})
			}
		}
	} else if p.to.typ == '&' {
		p.to.state_conj[p.from.name] = p.signal
		allHigh := true
		for _, s := range p.to.state_conj {
			if !s {
				allHigh = false
			}
		}
		toSend := true
		if allHigh && len(p.to.state_conj) == len(p.to.inputs) {
			toSend = false
		}
		for _, mod := range p.to.outputs {
			*pList = append(*pList, Pulse{p.to, mod, toSend})
		}
	}
}

func buildGraph(mod_map map[string]*Module) *Module {
	for _, m := range mod_map {
		for _, output_s := range m.outputs_s {
			if _, ok := mod_map[output_s]; !ok {
				m := Module{}
				m.typ = 'o'
				m.name = output_s
				mod_map[m.name] = &m
			}
			m.outputs = append(m.outputs, mod_map[output_s])
			mod_map[output_s].inputs = append(mod_map[output_s].inputs, m)
			mod_map[output_s].state_conj = map[string]bool{}
		}
	}

	return mod_map["roadcaster"]
}

func parse(input []string) map[string]*Module {
	mod_map := map[string]*Module{}
	for i := 0; i < len(input); i++ {
		m := Module{}
		m.outputs_s = strings.Split(strings.Split(input[i], " -> ")[1], ", ")
		m.typ = input[i][0]
		m.name = strings.Split(input[i], " -> ")[0][1:]
		mod_map[m.name] = &m
	}
	return mod_map
}
