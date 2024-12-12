package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"

	as "github.com/beefsack/go-astar"
	"golang.org/x/exp/maps"
)

type Day10 struct {
	island     world              // map x,y to elevation
	trailHeads map[complex128]int // trails as x,y to score (trails start at 0, score is how many summits can be reached)
	peaks      []complex128       // all worthy peaks (9 elevation)
	highBounds complex128
	lowBounds  complex128
}

// world is a 2d map of tiles
type world map[complex128]*tile

type tile struct {
	// elevation is how tall a tile is potentially affecting movement
	elevation int
	// XY coords
	position complex128
	// W is a reference to the world that the tile is a part of
	w world
}

func (p *Day10) parseInput(lines []string) {
	p.lowBounds = complex(0.0, 0.0)
	p.highBounds = complex(float64(len(lines)), float64(len(lines[0])))
	p.island = map[complex128]*tile{}
	p.trailHeads = map[complex128]int{}
	p.peaks = []complex128{}
	for row, cols := range lines {
		for location, height := range cols {
			elevation := common.Atoi(string(height))
			coords := complex(float64(row), float64(location))
			nTile := tile{elevation: elevation, position: coords, w: p.island}
			p.island[coords] = &nTile
			switch elevation {
			case 0:
				p.trailHeads[coords] = 0
			case 9:
				p.peaks = append(p.peaks, coords)
			}
		}
	}
}

// borkeneded atm
//
//lint:ignore U1000 not always output
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
			fmt.Printf("%s%d", mark, out.elevation)
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

// PathNeighbors returns the neighbors of the tile excluding...
//
//	tiles off the edge of the map, and tiles too tall to reach
func (t *tile) PathNeighbors() []as.Pather {
	neighbors := []as.Pather{}
	offsets := []complex128{complex(-1, 0), complex(1, 0), complex(0, -1), complex(0, 1)}
	for _, offset := range offsets {
		// does this neighbor exist and is it 1 unit higher
		if n, ok := t.w[t.position+offset]; ok && n.elevation == t.elevation+1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// neighbors returns the neighbors of the tile excluding...
//
//	tiles off the edge of the map, and tiles too tall to reach
func (t *tile) neighbors() []*tile {
	neighbors := []*tile{}
	offsets := []complex128{complex(-1, 0), complex(1, 0), complex(0, -1), complex(0, 1)}
	for _, offset := range offsets {
		// does this neighbor exist and is it 1 unit higher
		if n, ok := t.w[t.position+offset]; ok && n.elevation == t.elevation+1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the neighboring tile (always 1!)
func (t *tile) PathNeighborCost(to as.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
//
//	between non-adjacent tiles. this is meaningless here but required
func (t *tile) PathEstimatedCost(to as.Pather) float64 {
	toT := to.(*tile)
	absX := math.Abs(float64(column(toT.position) - column(t.position)))
	absY := math.Abs(float64(row(toT.position) - row(t.position)))
	return absX + absY
}

func findAllPaths(world map[complex128]*tile, start *tile) (rank int) {
	rank = 0
	neighbors := start.neighbors()
	for _, neighbor := range neighbors {
		if neighbor.elevation == 9 {
			rank++
		} else {
			rank += findAllPaths(world, neighbor)
		}
	}
	return rank
}

// commented for speed
func (p Day10) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	for trailHead := range p.trailHeads {
		for _, peak := range p.peaks {
			if _, _, found := as.Path(p.island[trailHead], p.island[peak]); found {
				p.trailHeads[trailHead]++
			}
		}
	}
	sum := 0
	for _, score := range p.trailHeads {
		sum += score
	}
	return sum
}

func (p Day10) PartB(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	sum := 0
	for trailHead := range p.trailHeads {
		sum += findAllPaths(p.island, p.island[trailHead])
	}
	return sum
}
