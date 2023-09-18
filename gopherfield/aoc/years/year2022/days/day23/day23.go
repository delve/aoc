package year2022

import (
	"fmt"
	"math"
	"strings"
)

type Day23 struct{}

type elf struct {
	xy complex128
}

// proposeMove implements the movement logic for an elf
func proposeMove(elf complex128, elves map[complex128]int) (complex128, bool) {
	/* if there is no elf in the N, NE, or NW adjacent positions,
	   the elf proposes moving north 1 step. */
	_, hasElf1 := elves[elf+complex(-1, -1)]
	_, hasElf2 := elves[elf+complex(0, -1)]
	_, hasElf3 := elves[elf+complex(1, -1)]
	if !hasElf1 && !hasElf2 && !hasElf3 {
		return elf + complex(0, -1), true // move north
	}

	/* if there is no elf in the S, SE, or SW adjacent positions,
	   the elf proposes moving south 1 step. */
	_, hasElf1 = elves[elf+complex(-1, 1)]
	_, hasElf2 = elves[elf+complex(0, 1)]
	_, hasElf3 = elves[elf+complex(1, 1)]
	if !hasElf1 && !hasElf2 && !hasElf3 {
		return elf + complex(0, 1), true // move south
	}

	/* if there is no elf in the W, NW, or SW adjacent positions,
	   the elf proposes moving west 1 step. */
	_, hasElf1 = elves[elf+complex(-1, -1)]
	_, hasElf2 = elves[elf+complex(-1, 0)]
	_, hasElf3 = elves[elf+complex(-1, 1)]
	if !hasElf1 && !hasElf2 && !hasElf3 {
		return elf + complex(-1, 0), true // move west
	}

	/* if there is no elf in the E, NE, or SE adjacent positions,
	   the elf proposes moving east 1 step. */
	_, hasElf1 = elves[elf+complex(1, -1)]
	_, hasElf2 = elves[elf+complex(1, 0)]
	_, hasElf3 = elves[elf+complex(1, 1)]
	if !hasElf1 && !hasElf2 && !hasElf3 {
		return elf + complex(1, 0), true // move east
	}

	// finally, if no options are open, don't move
	return 0, false
}

// getBounds returns two complex numbers representing the minimum and maximum x,y extent of elves in an array as real,imag
func getBounds(elves map[complex128]int) (complex128, complex128) {
	// something in this function is wrong! it --sometimes-- gets the minimums wrong?

	var min, max complex128
	min = complex(math.MaxFloat64, math.MaxFloat64)
	for e := range elves {
		// update X
		if real(e) < real(min) {
			min = complex(real(e), imag(min))
		} else if real(e) >= real(max) {
			max = complex(real(e)+1, imag(max))
		}
		// update Y
		if imag(e) < imag(min) {
			min = complex(real(min), imag(e))
		} else if imag(e) >= imag(max) {
			max = complex(real(max), imag(e))
		}
	}

	return min, max
}

// drawGrove returns a string slice representing the positions of all elves
func drawGrove(e map[complex128]int) []string {
	grove := []string{}
	min, max := getBounds(e)
	for y := imag(min); y < imag(max); y++ {
		row := ""
		for x := real(min); x < real(max); x++ {
			if _, ok := e[complex(x, y)]; ok {
				row += "#"
			} else {
				row += "."
			}
		}
		grove = append(grove, row)
	}
	return grove
}

func (p Day23) PartA(lines []string) any {
	elves := make(map[complex128]int)
	for y, line := range lines {
		for x, v := range line {
			if v == '#' {
				elves[complex(float64(x), float64(y))] = 1
				fmt.Printf("new elf at %d %d\n", x, y)
				bMin, bMax := getBounds(elves)
				fmt.Printf("Bounds: %v %v\n", bMin, bMax)
			}
		}
	}
	// initial state
	// bMin, bMax := getBounds(elves)
	// fmt.Println(drawGrove(elves))

	numRounds := 10
	for i := 0; i < numRounds; i++ {
		proposals := make(map[complex128][]complex128)
		for elf := range elves {
			if target, move := proposeMove(elf, elves); move {
				proposals[target] = append(proposals[target], elf)
			}
		}
		for target, sources := range proposals {
			if len(sources) == 1 {
				elves[target] = 1
				delete(elves, sources[0])
			} // else more than 1 elf proposed this target, so none move
		}
	}
	fmt.Printf("After %d rounds of movement:\n", numRounds)
	bMin, bMax := getBounds(elves)
	fmt.Printf("Bounds: %v %v\n", bMin, bMax)
	fmt.Println(strings.Join(drawGrove(elves),"\n"))

	return "implement_me"
}

func (p Day23) PartB(lines []string) any {
	return "implement_me"
}
