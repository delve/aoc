package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

type Day10 struct{}

const (
	maxCycle = 241
)

type sigStrengths struct {
	c20, c60, c100, c140, c180, c220 int
}

func (s sigStrengths) sum() int {
	return s.c20 + s.c60 + s.c100 + s.c140 + s.c180 + s.c220
}

type addRegister struct {
	cycles, amount int
	register       *int
}

func NewAddRegister(amount int, register *int) addRegister {
	return addRegister{cycles: 2, amount: amount, register: register}
}

func (n *addRegister) tick() bool {
	n.cycles += -1
	if n.cycles > 0 {
		return true
	} else {
		*n.register += n.amount
		return false
	}
}

type noop struct {
	cycles int
}

func NewNoop() noop {
	return noop{cycles: 1}
}

func (n *noop) tick() bool {
	n.cycles += -1
	if n.cycles > 0 {
		return true
	} else {
		return false
	}
}

type command interface {
	tick() bool // return false if command expired
}

func parseInput10(lines []string, register *int) []command {
	var commands []command

	for _, line := range lines {
		inLine := strings.Split(line, " ")
		switch inLine[0] {
		case "noop":
			c := NewNoop()
			commands = append(commands, &c)
		case "addx":
			a, err := strconv.Atoi(inLine[1])
			common.Check(err)
			c := NewAddRegister(a, register)
			commands = append(commands, &c)
		}
	}
	return commands
}

func (p Day10) PartA(lines []string) any {
	xReg := 1
	commands := parseInput10(lines, &xReg)

	allSig := make(map[int]int)
	comCursor := 0
	var curCommand command

	for i := 1; i < maxCycle; i++ {
		// get signal
		allSig[i] = i * xReg
		if curCommand == nil {
			if comCursor < len(commands) {
				curCommand = commands[comCursor]
				comCursor++
			}
		}
		if curCommand != nil {
			processing := curCommand.tick()
			if !processing {
				curCommand = nil
			}
		}
	}

	signals := sigStrengths{
		c20:  allSig[20],
		c60:  allSig[60],
		c100: allSig[100],
		c140: allSig[140],
		c180: allSig[180],
		c220: allSig[220],
	}
	return signals.sum()
}

func (p Day10) PartB(lines []string) any {
	xReg := 1
	commands := parseInput10(lines, &xReg)

	var crtString []string
	comCursor := 0
	var curCommand command

	for i := 1; i < maxCycle; i++ {
		pixel := (i - 1) % 40
		if pixel >= xReg-1 && pixel <= xReg+1 {
			crtString = append(crtString, "#")
		} else {
			crtString = append(crtString, ".")
		}

		if curCommand == nil {
			if comCursor < len(commands) {
				curCommand = commands[comCursor]
				comCursor++
			}
		}
		if curCommand != nil {
			processing := curCommand.tick()
			if !processing {
				curCommand = nil
			}
		}
	}

	fmt.Println("Display:")
	fmt.Printf("%v\n", crtString[0:40])
	fmt.Printf("%v\n", crtString[40:80])
	fmt.Printf("%v\n", crtString[80:120])
	fmt.Printf("%v\n", crtString[120:160])
	fmt.Printf("%v\n", crtString[160:200])
	fmt.Printf("%v\n", crtString[200:])
	return "see console"
}
