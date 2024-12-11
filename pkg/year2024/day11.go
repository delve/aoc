package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day11 struct {
	stones []int
}

func (p *Day11) parseInput(lines []string) {
	if len(lines) > 1 {
		panic("too many lines")
	}
	p.stones = make([]int, 0, 1000000000)
	for _, stone := range strings.Split(lines[0], " ") {
		p.stones = append(p.stones, common.Atoi(stone))
	}

}

func (p *Day11) blink() {
	stoneCount := len(p.stones)
	for idx := 0; idx < stoneCount; idx++ {
		if p.stones[idx] == 0 {
			p.stones[idx] = 1
			continue
		}
		if p.splitStone(idx) {
			idx++
			stoneCount++
			continue
		}
		p.stones[idx] = p.stones[idx] * 2024
	}
}

func (p *Day11) splitStone(index int) bool {
	txt := strconv.Itoa(p.stones[index])
	if math.Mod(float64(len(txt)), 2) == 0 {
		newLen := len(txt) / 2
		newStones := []int{
			common.Atoi(txt[:newLen]),
			common.Atoi(txt[newLen:]),
		}
		p.stones = append(p.stones[:index], append(newStones, p.stones[index+1:]...)...)
		return true
	}

	return false
}

// commented for speed
func (p Day11) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	// for i := 0; i < 25; i++ {
	// 	p.blink()
	// }
	return len(p.stones)
}

func (p Day11) PartB(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	for i := 0; i < 75; i++ {
		if math.Mod(float64(i), 10) == 0 {
			fmt.Printf("%d\n", len(p.stones))
		}
		p.blink()
	}
	return len(p.stones)
}
