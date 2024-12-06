package year2024

import "fmt"

type Day06 struct {
	lab      map[complex128]rune
	startPos complex128
}

// TODO: extract cplx coordinate map & walk into library. walk accepts pointer reciever func to avoid repeating loops everywhere

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

func (p *Day06) walkGuard() (uniqueCount int) {
	uniqueCount = 1
	direction := 0 // 0 = up, 0-3 corresponding to right turn cardinal directions
	position := p.startPos
	p.lab[position] = 'X'
	place := ' '
	for inMap := true; inMap; {
		var nextPos complex128
		switch direction {
		case 0:
			nextPos = position - 1
		case 1:
			nextPos = position + 1i
		case 2:
			nextPos = position + 1
		case 3:
			nextPos = position - 1i
		}

		place, inMap = p.lab[nextPos]
		if !inMap {
			return uniqueCount
		}
		if place == '#' {
			direction = (direction + 1) % 4
			continue
		}
		if place == '.' {
			p.lab[nextPos] = 'X'
			uniqueCount++
			position = nextPos
		}
		if place == 'X' {
			position = nextPos
		}
	}

	return uniqueCount
}

func (p Day06) PartA(lines []string) any {
	p.buildLabMap(lines[:len(lines)-1])
	// p.printMap()
	uniqueCount := p.walkGuard()
	// p.printMap()
	return uniqueCount
}

func (p Day06) PartB(lines []string) any {
	return "implement_me"
}
