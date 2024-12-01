package year2021

import (
	"aocgen/pkg/common"
	"fmt"
)

type Day01 struct {
	depths []int
}

func (p *Day01) parseInput(lines []string) {
	for _, v := range lines[:len(lines)-1] {
		p.depths = append(p.depths, common.Atoi(v))
	}
}

func (p Day01) PartA(lines []string) any {
	p.parseInput(lines)
	deepening := 0
	// fmt.Print(p.depths)
	for idx, val := range p.depths {
		if idx > 0 {
			if p.depths[idx-1] < val {
				deepening++
			}
		}
	}
	return fmt.Sprintf("%v", deepening)
}

func (p Day01) PartB(lines []string) any {
	p.parseInput(lines)
	deepening := 0
	// fmt.Print(p.depths)
	for idx := range p.depths {
		if idx >= 3 {
			window1 := p.depths[idx-3] + p.depths[idx-2] + p.depths[idx-1]
			window2 := p.depths[idx-2] + p.depths[idx-1] + p.depths[idx-0]
			if window1 < window2 {
				deepening++
			}
		}
	}
	return fmt.Sprintf("%v", deepening)
}
