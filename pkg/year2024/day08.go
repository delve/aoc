package year2024

import (
	"fmt"
	"math"

	"golang.org/x/exp/maps"
)

type Day08 struct {
	antinodes  map[complex128][]rune
	antennas   map[rune][]complex128
	city       map[complex128]rune
	highBounds complex128
}

func (p *Day08) parseInput(lines []string) {
	p.antinodes = map[complex128][]rune{}
	p.antennas = map[rune][]complex128{}
	p.city = map[complex128]rune{}

	rows, columns := 0, 0
	for row, positions := range lines {
		cCount := 0
		for column, position := range positions {
			cCount++
			coord := complex(float64(row), float64(column))
			p.city[coord] = position
			if position != '.' {
				p.antennas[position] = append(p.antennas[position], coord)
			}
		}
		columns = cCount
		rows++
	}
	p.highBounds = complex(float64(rows), float64(columns))
}

//lint:ignore U1000 BECAUSE I SAID SO YOU ANAL RETARD
func (p Day08) printMap(includeAntinodes bool) {
	rowHeader := "%2d:"
	rowNum := 0
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; {
		fmt.Printf(rowHeader, rowNum)
		rowNum++
		for moreColumns := true; moreColumns; {
			out := string(p.city[position])
			if includeAntinodes {
				if _, ok := p.antinodes[position]; ok {
					out = "#"
				}
			}
			fmt.Printf("%s", out)
			if _, moreColumns = p.city[position+1i]; moreColumns {
				position += 1i
			}
		}
		fmt.Print("\n")
		if _, moreRows = p.city[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}
	fmt.Printf("Antennas: %v\n", p.antennas)
	fmt.Printf("Antinodes: %v\n ", p.antinodes)
}

func (p *Day08) getAntinode(frequency rune, antenna complex128, partners []complex128, includeAll bool) {
	for _, partner := range partners {
		if includeAll {
			p.antinodes[antenna] = append(p.antinodes[antenna], frequency)
		}
		count := 1
		if includeAll {
			count = -1
		}
		outOfBounds := false
		dist := antenna - partner
		for ; count != 0 && !outOfBounds; count-- {
			antinode := antenna + dist*complex(math.Abs(float64(count)), 0)

			// bounds check
			if row(antinode) >= 0 && row(antinode) < row(p.highBounds) &&
				column(antinode) >= 0 && column(antinode) < column(p.highBounds) {

				p.antinodes[antinode] = append(p.antinodes[antinode], frequency)
			} else {
				outOfBounds = true
			}
		}
	}
}

func (p Day08) getAntinodes() {
	for _, frequency := range maps.Keys(p.antennas) {
		for idx, position := range p.antennas[frequency] {
			partnerAntennas := make([]complex128, len(p.antennas[frequency]))
			copy(partnerAntennas, p.antennas[frequency])
			partnerAntennas = append(partnerAntennas[:idx], partnerAntennas[idx+1:]...)
			p.getAntinode(frequency, position, partnerAntennas, false)
		}
	}
}

func (p Day08) getAntinodes2() {
	for _, frequency := range maps.Keys(p.antennas) {
		for idx, position := range p.antennas[frequency] {
			partnerAntennas := make([]complex128, len(p.antennas[frequency]))
			copy(partnerAntennas, p.antennas[frequency])
			partnerAntennas = append(partnerAntennas[:idx], partnerAntennas[idx+1:]...)
			p.getAntinode(frequency, position, partnerAntennas, true)
		}
	}
}

func (p Day08) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	// p.printMap(false)
	p.getAntinodes()
	// println("XXX")
	// p.printMap(true)
	return len(p.antinodes)
}

func (p Day08) PartB(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	// p.printMap(false)
	p.getAntinodes2()
	// println("XXX")
	// p.printMap(true)
	return len(p.antinodes)
}
