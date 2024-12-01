package year2021

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strings"
)

type Day02 struct {
	horz  int
	depth int
	aim int
}

func (p Day02) PartA(lines []string) any {
	for _, input := range lines {
		direction := strings.Split(input, " ")
		switch direction[0] {
		case "forward":
			p.horz += common.Atoi(direction[1])
		case "up":
			p.depth += common.Atoi(direction[1])
		case "down":
			p.depth -= common.Atoi(direction[1])
		}
	}

	return fmt.Sprintf("%v", p.horz*int(math.Abs(float64(p.depth))))
}

func (p Day02) PartB(lines []string) any {
	for _, input := range lines {
		direction := strings.Split(input, " ")
		switch direction[0] {
		case "forward":
			p.horz += common.Atoi(direction[1])
			p.depth += common.Atoi(direction[1]) * p.aim
		case "up":
			p.aim -= common.Atoi(direction[1])
		case "down":
			p.aim += common.Atoi(direction[1])
		}
	}

	return fmt.Sprintf("%v", p.horz*int(math.Abs(float64(p.depth))))
}
