package year2024

import "fmt"

type Day06 struct {
	lab map[complex128]rune
}

func (p *Day06) buildLabMap(input []string) {
	p.lab = map[complex128]rune{}
	position := complex(0.0, 0.0)
	for _, row := range input {
		for _, column := range row {
			p.lab[position] = column
			position += 1i
		}
		position += 1.0
	}
}

func (p Day06) printMap() {
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; position += 1.0 {
		for moreColumns := true; moreColumns; position += 1i {
			fmt.Printf("%s", string(p.lab[position]))
			_, moreColumns = p.lab[position+1i]
		}
		fmt.Print("\n")
		_, moreRows = p.lab[position+1.0]
	}
}

func (p Day06) getStartPosition() complex128 {
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; position += 1.0 {
		for moreColumns := true; moreColumns; position += 1i {
			if p.lab[position] == '^' {
				return position
			}
			_, moreColumns = p.lab[position+1i]
		}
		_, moreRows = p.lab[position+1.0]
	}
	panic(fmt.Errorf("no start position found"))
}

func (p *Day06) walkGuard() (steps int) {
	steps = 0
	direction := 0 // 0 = up, 0-3 corresponding to right turn cardinal directions
	position := p.getStartPosition()
	return steps
}

func (p Day06) PartA(lines []string) any {
	p.buildLabMap(lines[:len(lines)-1])
	p.printMap()
	steps := p.walkGuard()
	p.printMap()
	return steps
}

func (p Day06) PartB(lines []string) any {
	return "implement_me"
}
