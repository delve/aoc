package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"regexp"
)

type Day13 struct {
	games []game
}

type game struct {
	buttonA       complex128
	buttonB       complex128
	aCost         int
	bCost         int
	prize         complex128
	solutions     [][2]int
	cheapestSolve int
	itrCount      float64
}

func (p *Day13) parseInput(lines []string) {
	buttonRex := regexp.MustCompile(`Button .: X\+([0-9]*), Y\+([0-9]*)`)
	prizeRex := regexp.MustCompile(`Prize: X=([0-9]*), Y=([0-9]*)`)
	for i := 0; i < len(lines); i++ {
		machine := game{aCost: 3, bCost: 1, cheapestSolve: math.MaxInt64, itrCount: 0,solutions: [][2]int{}}
		match := buttonRex.FindStringSubmatch(lines[i])
		machine.buttonA = complex(float64(common.Atoi(match[1])), float64(common.Atoi(match[2])))
		i++
		match = buttonRex.FindStringSubmatch(lines[i])
		machine.buttonB = complex(float64(common.Atoi(match[1])), float64(common.Atoi(match[2])))
		i++
		match = prizeRex.FindStringSubmatch(lines[i])
		machine.prize = complex(float64(common.Atoi(match[1])), float64(common.Atoi(match[2])))
		i++ // skip blank line
		p.games = append(p.games, machine)
	}
}

func (p *Day13) playGames() (tokenCost int) {
	tokenCost = 0

	for _, game := range p.games {
		game.solve(0, 0, complex(0, 0))
		if len(game.solutions) > 0 {
			tokenCost += game.cheapestSolve
		}
	}

	return tokenCost
}

func (g *game) solve(aCount, bCount int, curLoc complex128) {
	g.itrCount++
	if math.Mod(g.itrCount, 1000) == 0 {
		fmt.Printf("Iteration %0f, aCount %d | bCount %d | location %v | target %v\n", g.itrCount, aCount, bCount, curLoc, g.prize)
	}
	if real(curLoc) > real(g.prize) || imag(curLoc) > imag(g.prize) {
		return // too far, just back up
	}
	if curLoc == g.prize {
		// got it, calc cost and update solves
		g.solutions = append(g.solutions, [2]int{aCount, bCount})
		cost := g.aCost*aCount + g.bCost*bCount
		if cost < g.cheapestSolve {
			g.cheapestSolve = cost
		}
		return
	}
	if aCount == 100 || bCount == 100 {
		return // puzzle says no more than 100 pushes each
	}
	// else follow all branches
	g.solve(aCount+1, bCount, curLoc+g.buttonA)
	g.solve(aCount, bCount+1, curLoc+g.buttonB)
}

func (p Day13) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])

	p.playGames()
	fmt.Printf("%v", p.games)

	return "implement_me"
}

func (p Day13) PartB(lines []string) any {
	return "implement_me"
}
