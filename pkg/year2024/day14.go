package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	termbox "github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
)

type Day14 struct {
	robots []*robot
	bounds [2]int // maxRow, maxColumn
}

type robot struct {
	id       int
	position complex128
	velocity complex128
}

func (p *Day14) parseInput(lines []string) {
	bounds := strings.Split(lines[0], ",")
	p.bounds[0] = common.Atoi(bounds[0])
	p.bounds[1] = common.Atoi(bounds[1])

	stateRex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	id := 0
	for _, rState := range lines[1:] {
		state := stateRex.FindSubmatch([]byte(rState))
		position := complex(common.MustFloat(string(state[2])), common.MustFloat(string(state[1])))
		velocity := complex(common.MustFloat(string(state[4])), common.MustFloat(string(state[3])))
		p.robots = append(p.robots, &robot{id: id, position: position, velocity: velocity})
		id++
	}
}

func (p *Day14) walkRobots(seconds float64) {
	for _, bot := range p.robots {
		bot.walk(seconds, p.bounds)
	}
}

func (p Day14) calcSafetyFactor() int {
	quadRows := p.bounds[0] / 2
	quadColumns := p.bounds[1] / 2
	botCounts := [4]int{0, 0, 0, 0}
	for _, bot := range p.robots {
		// q0
		if row(bot.position) < quadRows && column(bot.position) < quadColumns {
			botCounts[0]++
			continue
		}
		// q1
		if row(bot.position) < quadRows && column(bot.position) > quadColumns {
			botCounts[1]++
			continue
		}
		// q2
		if row(bot.position) > quadRows && column(bot.position) > quadColumns {
			botCounts[2]++
			continue
		}
		// q3
		if row(bot.position) > quadRows && column(bot.position) < quadColumns {
			botCounts[3]++
			continue
		}
	}
	return botCounts[0] * botCounts[1] * botCounts[2] * botCounts[3]
}

func (r *robot) walk(seconds float64, bounds [2]int) {
	// full line method
	newPos := r.position + complex(seconds*real(r.velocity), seconds*imag(r.velocity))
	newRow := math.Mod(real(newPos), float64(bounds[0]))
	if newRow < 0 {
		newRow = float64(bounds[0]) + newRow
	}
	if newRow > float64(bounds[0]) {
		newRow = newRow - float64(bounds[0])
	}
	newColumn := math.Mod(imag(newPos), float64(bounds[1]))
	if newColumn < 0 {
		newColumn = float64(bounds[1]) + newColumn
	}
	if newColumn > float64(bounds[1]) {
		newColumn = newColumn - float64(bounds[1])
	}

	r.position = complex(newRow, newColumn)
}

func (p Day14) getRobotMap() map[complex128]int {
	area := map[complex128]int{}
	for _, bot := range p.robots {
		if _, ok := area[bot.position]; !ok {
			area[bot.position] = 1
		} else {
			area[bot.position] += 1
		}
	}
	return area
}

func (p Day14) getAreaMap() string {
	area := p.getRobotMap()

	cursor := complex(0, 0)
	var botMap strings.Builder

	for row := 0; row < p.bounds[0]; row++ {
		for col := 0; col < p.bounds[1]; col++ {
			if count, ok := area[cursor]; ok {
				botMap.WriteString(fmt.Sprintf("%d", count))
			} else {
				botMap.WriteString(".")
			}
			cursor += 1i
		}
		botMap.WriteString("\n")
		cursor -= complex(0, imag(cursor))
		cursor += 1
	}
	return botMap.String()
}

// input prep: in order to support running sample in debug and full input when run, and the
// change in arena size between the two for this day, prepend both inputs with a line "maxRow,maxColumn"
// for the appropriate scenario
func (p Day14) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])

	// println(p.getAreaMap())
	p.walkRobots(100)
	// println(p.getAreaMap())
	return p.calcSafetyFactor()
}

func (p Day14) PartB(lines []string) any {
	running := true
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	kb := termbox.New()
	kb.Bind(func() {
		fmt.Println("pressed space!")
		running = false
	}, "space")

	for running {
		kb.Poll(term.PollEvent())
	}

	p.parseInput(lines[:len(lines)-1])
	// println(p.getAreaMap())
	var out strings.Builder
	for i := 0; i < 2000; i++ {
		p.walkRobots(1)
		if i >= 1000 {
			out.WriteString("\n\n\n")
			out.WriteString(fmt.Sprintf("Seconds: %d \n", i+1))
			out.WriteString(p.getAreaMap())
		}
	}
	err = os.WriteFile("/tmp/out", []byte(out.String()), 0644)
	common.Check(err)

	return "implement_me"
}
