package year2022

import (
	"aocgen/pkg/common"
	"math"
	"regexp"
)

// had to compare w/ https://github.com/aarneng/AdventOfCode2022/blob/main/day14/main.go to find an issue. so glad i learned about complex numbers though!!!!!
type Day14 struct{}

// output tile states (falling just traces the last path)
const (
	kindEmpty = iota
	kindFalling
	kindSand
	kindRock
)

// sand states
const (
	sandFalling = iota
	sandStopped
	sandEscaped
	sandBlocked
)

// kindRunes map tile kinds to output runes
// still with the damned name collisions >.<
var kindRunes14 = map[int]rune{
	kindEmpty:   '.',
	kindFalling: '~',
	kindSand:    'o',
	kindRock:    '#',
}

type sand struct {
	xyCoord complex128 // coords
	state   int        // const iota representing state
}

func newSand() sand {
	return sand{
		xyCoord: complex(500, 0.0),
		state:   sandFalling,
	}
}

func (s *sand) fall(cave map[complex128]int, yMax float64) {
	if s.xyCoord == complex(500, 0) && cave[s.xyCoord] > kindFalling {
		s.state = sandBlocked
		return
	}
	switch {
	case imag(s.xyCoord) >= yMax:
		// escaped
		s.state = sandEscaped
	case cave[s.xyCoord+1i] < kindSand:
		// straight down
		s.xyCoord += 1i
		s.state = sandFalling
	case cave[s.xyCoord-1+1i] < kindSand:
		// down left
		s.xyCoord += -1 + 1i
		s.state = sandFalling
	case cave[s.xyCoord+1+1i] < kindSand:
		// down right
		s.xyCoord += 1 + 1i
		s.state = sandFalling
	default:
		// fell through to here, nowhere to go
		s.state = sandStopped
	}
}

func parseInput14(lines []string) (map[complex128]int, float64, float64, float64) {
	cave := map[complex128]int{}
	xMin, xMax, floor := math.MaxFloat64, 0.0, 0.0
	re := regexp.MustCompile("[0-9]+")

	for _, line := range lines {
		inputs := re.FindAllString(line, -1)
		x1 := common.MustFloat(inputs[0])
		y1 := common.MustFloat(inputs[1])
		// check cave extents
		if y1 > floor {
			floor = y1
		}
		if x1 < xMin {
			xMin = x1
		}
		if x1 > xMax {
			xMax = x1
		}

		for i := 2; i < len(inputs); i += 2 {
			xNext := common.MustFloat(inputs[i])
			yNext := common.MustFloat(inputs[i+1])
			// check cave extents
			if yNext > floor {
				floor = yNext
			}
			if xNext < xMin {
				xMin = xNext
			}
			if xNext > xMax {
				xMax = xNext
			}

			if xNext == x1 {
				yL := y1
				yH := yNext
				if yL > yH {
					yL, yH = yH, yL
				}
				for y := yL; y <= yH; y++ {
					cave[complex(x1, y)] = kindRock
				}
			} else {
				xL := x1
				xH := xNext
				if xL > xH {
					xH, xL = xL, xH
				}
				for x := xL; x <= xH; x++ {
					cave[complex(x, y1)] = kindRock
				}
			}
			x1, y1 = xNext, yNext
		}
	}

	// adjust for 0 indexing
	floor++

	return cave, xMin, xMax, floor
}

func (p Day14) PartA(lines []string) any {
	cave, _, _, caveFloor := parseInput14(lines)
	var grains int64
	grains = 0
	for {
		grains++
		grain := newSand()
		for ; grain.state == sandFalling; grain.fall(cave, caveFloor) {
			cave[grain.xyCoord] = kindFalling
		}

		if grain.state != sandEscaped {
			cave[grain.xyCoord] = kindSand
		} else {
			// don't count the last grain, since it escaped
			grains += -1
			break
		}
	}

	return grains
}

func (p Day14) PartB(lines []string) any {
	cave, xMin, xMax, caveFloor := parseInput14(lines)
	// the floor is deeper, and solid!!
	caveFloor += 2
	var grains int64
	grains = 0
	for {
		grains++
		grain := newSand()
		for ; grain.state == sandFalling; grain.fall(cave, caveFloor) {
			if cave[grain.xyCoord] == kindEmpty {
				cave[grain.xyCoord] = kindFalling
			}
			// JIT-build the proper floor for this x val
			if imag(grain.xyCoord) == caveFloor-2 {
				cave[complex(real(grain.xyCoord), caveFloor-1)] = kindRock
				cave[complex(real(grain.xyCoord)+1, caveFloor-1)] = kindRock
				cave[complex(real(grain.xyCoord)-1, caveFloor-1)] = kindRock
				if real(grain.xyCoord)-1 < xMin {
					xMin = real(grain.xyCoord) - 1
				}
				if real(grain.xyCoord)+1 > xMax {
					xMax = real(grain.xyCoord) + 1
				}
			}
		}

		if grain.state < sandEscaped {
			cave[grain.xyCoord] = kindSand
		} else {
			// don't count the last grain, because it escaped or
			//  was blocked by the previous grain
			grains += -1
			break
		}
	}

	return grains
}
