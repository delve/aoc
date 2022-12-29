package year2022

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Day05 struct{}

type cargo struct {
	stacks [9]stack.Stack
}

// implements a one-at-a-time move
func (c *cargo) move(count, source, target int) {
	for i := 0; i < count; i++ {
		c.stacks[target-1].Push(c.stacks[source-1].Pop())
	}
}

// implements moving as a stack
func (c *cargo) stackMove(count, source, target int) {
	load := stack.New()
	for i := 0; i < count; i++ {
		load.Push(c.stacks[source-1].Pop())
	}
	for i := 0; i < count; i++ {
		c.stacks[target-1].Push(load.Pop())
	}
}

type order struct {
	count, source, target int
}

func parseInput05(lines []string) (cargo, []order) {
	cargo := cargo{
		stacks: [9]stack.Stack{},
	}
	orders := make([]order, 10)
	orderRex := regexp.MustCompile(`move (?P<count>\d+) from (?P<source>\d+) to (?P<target>\d+)`)
	flagLine := regexp.MustCompile(`( \d  )+`)
	passedDivider := false
	invertCache := []string{}
	for _, line := range lines {
		if !passedDivider {
			// check for the dividing line
			passedDivider = flagLine.Match([]byte(line))
			if passedDivider {
				// read & store the cargo data
				for i := len(invertCache) - 1; i >= 0; i-- {
					for j := 0; j < 9; j++ {
						// read each column on the line
						val := strings.TrimSpace(string(invertCache[i][j*4+1]))
						if len(val) > 0 { // empty spaces are trimmed to 0
							cargo.stacks[j].Push(val)
						}
					}
				}
				// and skip the divider line
				continue
			}

			// read cargo into the inversion cache for later processing
			invertCache = append(invertCache, line)
		} else {
			if line == "" {
				continue // skip the blank line
			}
			// read orders
			regOut := orderRex.FindStringSubmatch(line)
			count, _ := strconv.Atoi(regOut[1])
			source, _ := strconv.Atoi(regOut[2])
			target, _ := strconv.Atoi(regOut[3])
			orders = append(orders, order{
				count:  count,
				source: source,
				target: target,
			})
		}
	}
	return cargo, orders
}

func (p Day05) PartA(lines []string) any {
	cargo, orders := parseInput05(lines)
	for _, order := range orders {
		cargo.move(order.count, order.source, order.target)
	}
	topCrates := ""
	for i := 0; i < len(cargo.stacks); i++ {
		topCrates += cargo.stacks[i].Pop().(string)
	}
	return topCrates
}

func (p Day05) PartB(lines []string) any {
	cargo, orders := parseInput05(lines)
	for _, order := range orders {
		cargo.stackMove(order.count, order.source, order.target)
	}
	topCrates := ""
	for i := 0; i < len(cargo.stacks); i++ {
		topCrates += cargo.stacks[i].Pop().(string)
	}
	return topCrates
}
