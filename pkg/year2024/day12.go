package year2024

import (
	"fmt"
	"math"

	"golang.org/x/exp/maps"
)

type Day12 struct {
	plots   map[complex128]plot  // [coord]plot
	regions map[int]gardenRegion // [id]area; regions[-1] is invalid, the unknown state
}

type plot struct {
	kind   rune
	region int
}

// corners are numbered
// 3 0
//
//	X
//
// 2 1
type gardenRegion struct {
	area        int
	perimeter   int
	sides       int
	cornerPlots map[complex128][]int // coord, corners
}

func (p *Day12) parseInput(lines []string) {
	p.plots = map[complex128]plot{}
	p.regions = map[int]gardenRegion{}
	for rowNum, vals := range lines {
		for colNum, cell := range vals {
			coord := complex(float64(rowNum), float64(colNum))
			p.plots[coord] = plot{kind: cell, region: -1}
		}
	}
	p.identifyRegions()
}

func (p *Day12) identifyRegions() {
	position := complex(0.0, 0.0)
	regionId := -1
	for moreRows := true; moreRows; {
		for moreColumns := true; moreColumns; {
			if p.plots[position].region < 0 {
				regionId++
				region := gardenRegion{area: 0, perimeter: 0, sides: 0, cornerPlots: map[complex128][]int{}}
				neighbors := map[bool]map[complex128]bool{false: {position: false}, true: {}} // [searched][coord]ignored value
				for len(maps.Keys(neighbors[false])) > 0 {
					coord := maps.Keys(neighbors[false])[0]
					// remove from unsearched, add to searched
					delete(neighbors[false], coord)
					neighbors[true][coord] = true

					thisNeighbors, corners := p.getNeighborsAndCorners(coord)
					// add neighbors to unsearched, if not already searched
					for _, neighbor := range thisNeighbors {
						if _, searched := neighbors[true][neighbor]; !searched {
							neighbors[false][neighbor] = false
						}
					}

					if len(corners) > 0 {
						region.sides += len(corners)
						region.cornerPlots[coord] = corners
					}

					// add perimter
					region.perimeter += 4 - len(thisNeighbors)
				}

				// update plot region membership
				for _, coord := range maps.Keys(neighbors[true]) {
					update := p.plots[coord]
					update.region = regionId
					p.plots[coord] = update
				}

				// final region calcs, and add to Day12 map
				region.area = len(maps.Keys(neighbors[true]))
				p.regions[regionId] = region
			}

			if _, moreColumns = p.plots[position+1i]; moreColumns {
				position += 1i
			}
		}
		if _, moreRows = p.plots[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}

}

//lint:ignore U1000 not always output
func (p Day12) printMap() {
	fmt.Print("Regions:\n")
	for id, region := range p.regions {
		fmt.Printf("ID: %d Area: %d Perimeter: %d Sides: %d Corners:\n%v\n", id, region.area, region.perimeter, region.sides, region.cornerPlots)
	}
	rowHeader := "%2d:"
	rowNum := 0
	position := complex(0.0, 0.0)
	for moreRows := true; moreRows; {
		fmt.Printf(rowHeader, rowNum)
		rowNum++
		for moreColumns := true; moreColumns; {
			fmt.Printf("%s%d ", string(p.plots[position].kind), p.plots[position].region)
			if _, moreColumns = p.plots[position+1i]; moreColumns {
				position += 1i
			}
		}

		fmt.Print("\n")
		if _, moreRows = p.plots[position+1.0]; moreRows {
			position = complex(real(position)+1.0, 0i)
		}
	}
}

func (p Day12) getNeighborsAndCorners(coord complex128) ([]complex128, []int) {
	neighbors := []complex128{}
	corners := []int{}
	offsets := []complex128{complex(-1, 0), complex(0, 1), complex(1, 0), complex(0, -1)}
	cornerOffsets := []complex128{complex(-1, 1), complex(1, 1), complex(1, -1), complex(-1, -1)}

	for idx, offset := range offsets {
		// does this coord exist and is it the same kind
		neighbor, exists := p.plots[coord+offset]
		if exists && neighbor.kind == p.plots[coord].kind {
			neighbors = append(neighbors, coord+offset)
		} else {
			cornerNeighbor, cornerExists := p.plots[coord+cornerOffsets[idx]]
			nextNeighbor, nextExists := p.plots[coord+offsets[int(math.Mod(float64(idx+1), 4.0))]]
			// convex corner
			if !nextExists || nextNeighbor.kind != p.plots[coord].kind {
				corners = append(corners, idx)
				continue
			}
			// concave corner
			if cornerExists && cornerNeighbor.kind == p.plots[coord].kind {
				corners = append(corners, idx)
				continue
			}
		}
	}
	return neighbors, corners
}

func (p Day12) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	// p.printMap()
	fenceCost := 0
	for _, region := range p.regions {
		fenceCost += region.area * region.perimeter
	}
	return fenceCost
}

func (p Day12) PartB(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	// p.printMap()

	fenceCost := 0
	for _, region := range p.regions {
		fenceCost += region.area * region.sides
	}
	return fenceCost
}
