package year2024

import "fmt"

type Day06 struct {
	lab      map[complex128]rune
	startPos complex128
}

func (p *Day06) buildLabMap(input []string) {
	p.lab = map[complex128]rune{}
	p.startPos = complex(-1.0, -1.0)
	position := complex(0.0, 0.0)
	for _, row := range input {
		for _, place := range row {
			p.lab[position] = place
			if place == '^' {
				p.startPos = position
			}
			position += 1i
		}
		position = complex(real(position)+1.0, 0i)
	}
	if p.startPos == complex(-1.0, -1.0) {
		panic(fmt.Errorf("no start position found"))
	}
}

//lint:ignore U1000 BECAUSE I SAID SO YOU ANAL RETARD
func (p Day06) printMap() {
	fmt.Printf("Start Position %d, %d\n", row(p.startPos), column(p.startPos))
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; {
		for moreColumns := true; moreColumns; {
			fmt.Printf("%s", string(p.lab[position]))
			if _, moreColumns = p.lab[position+1i]; moreColumns {
				position += 1i
			}
		}
		fmt.Print("\n")
		if _, moreRows = p.lab[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}
}

func (p Day06) countSteps() int {
	steps := 0
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; {
		for moreColumns := true; moreColumns; {
			if p.lab[position] == 'X' {
				steps++
			}
			if _, moreColumns = p.lab[position+1i]; moreColumns {
				position += 1i
			}
		}
		if _, moreRows = p.lab[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}
	return steps
}

func (p Day06) getNextPos(start complex128, direction int) complex128 {
	switch direction {
	case 0:
		return start - 1
	case 1:
		return start + 1i
	case 2:
		return start + 1
	case 3:
		return start - 1i
	}
	panic("unknown direction, you fall through the floor")
}

func (p *Day06) walkGuard() bool {
	direction := 0 // 0 = up, 0-3 corresponding to right turn cardinal directions
	position := p.startPos
	p.lab[position] = 'X'
	place := ' '
	visits := map[string]bool{}

	for inMap := true; inMap; {
		travelHash := fmt.Sprintf("%v%d", position, direction)
		if _, beenHere := visits[travelHash]; beenHere {
			return true
		} else {
			visits[travelHash] = true
		}
		nextPos := p.getNextPos(position, direction)

		place, inMap = p.lab[nextPos]
		if !inMap {
			return false
		}
		if place == '#' {
			direction = (direction + 1) % 4
			continue
		}
		if place == '.' {
			p.lab[nextPos] = 'X'
			position = nextPos
		}
		if place == 'X' {
			position = nextPos
		}
	}

	return false
}

func (p *Day06) obstructGuard() (loopingPaths int) {
	loopingPaths = 0
	obstructions := map[string]bool{}
	direction := 0 // 0 = up, 0-3 corresponding to right turn cardinal directions
	position := p.startPos
	p.lab[position] = 'X'
	place := ' '
	for inMap := true; inMap; {
		nextPos := p.getNextPos(position, direction)

		place, inMap = p.lab[nextPos]
		if !inMap {
			loopingPaths = len(obstructions)
			return loopingPaths
		}
		if place == '#' {
			direction = (direction + 1) % 4
			continue
		}
		if place != '#' {
			p.lab[nextPos] = '#'
			if p.walkGuard() {
				obstructions[fmt.Sprintf("%v\n", p.getNextPos(position, direction))] = true
			}
			p.lab[nextPos] = 'X'
			position = nextPos
		}
	}
	return loopingPaths
}

func (p Day06) PartA(lines []string) any {
	p.buildLabMap(lines[:len(lines)-1])
	// p.printMap()
	p.walkGuard()
	uniqueCount := p.countSteps()
	// p.printMap()
	return uniqueCount
}

func (p Day06) PartB(lines []string) any {
	p.buildLabMap(lines[:len(lines)-1])
	// p.printMap()
	obstructions := p.obstructGuard()
	// p.printMap()

	return obstructions
}
