package year2024

import (
	"aocgen/pkg/common"
	"fmt"

	"golang.org/x/exp/maps"
)

type Day10 struct {
	island     map[complex128]int // map x,y to elevation
	trailHeads map[complex128]int // trails as x,y to score (trails start at 0, score is how many summits can be reached)
	peaks      []complex128       // all worthy peaks (9 elevation)
	highBounds complex128
	lowBounds  complex128
}

func (p *Day10) parseInput(lines []string) {
	p.lowBounds = complex(0.0, 0.0)
	p.highBounds = complex(float64(len(lines)), float64(len(lines[0])))
	p.island = map[complex128]int{}
	p.trailHeads = map[complex128]int{}
	p.peaks = []complex128{}
	for row, cols := range lines {
		for location, height := range cols {
			elevation := common.Atoi(string(height))
			coords := complex(float64(row), float64(location))
			p.island[coords] = elevation
			switch elevation {
			case 0:
				p.trailHeads[coords] = 0
			case 9:
				p.peaks = append(p.peaks, coords)
			}
		}
	}
}

func (p Day10) printMap() {
	fmt.Printf("Trailheads: %d\nPeaks: %d\n", len(maps.Keys(p.trailHeads)), len(p.peaks))
	rowHeader := "%2d:"
	rowNum := 0
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; {
		fmt.Printf(rowHeader, rowNum)
		rowNum++
		for moreColumns := true; moreColumns; {
			out := p.island[position]
			mark := " "
			if _, ok := p.trailHeads[position]; ok {
				mark = "_"
			}
			for _, coord := range p.peaks {
				if coord == position {
					mark = "*"
				}
			}
			fmt.Printf("%s%d", mark, out)
			if _, moreColumns = p.island[position+1i]; moreColumns {
				position += 1i
			}
		}
		fmt.Print("\n")
		if _, moreRows = p.island[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}
}

func (p Day10) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	p.printMap()
	return "implement_me"
}

func (p Day10) PartB(lines []string) any {
	return "implement_me"
}
