package year2022

import (
	"aocgen/pkg/common"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Day11 struct{}

type monkey struct {
	items          []big.Float
	itemsInspected int
	test           int
	trueOp         int
	falseOp        int
	inspectOp      struct {
		string
		float64
	}
}

func (m *monkey) inspect(relief bool) {
	if len(m.items) == 0 {
		logrus.Fatalf("item inspection overrun on monkey!! %v", m)
	}
	old := m.items[0]
	var operand big.Float
	if m.inspectOp.float64 == 0 {
		// edge case for 'old' operand
		operand = old
	} else {
		operand = *big.NewFloat(m.inspectOp.float64)
	}

	switch m.inspectOp.string {
	case "+":
		m.items[0] = *old.Add(&old, &operand)
	case "*":
		m.items[0] = *old.Mul(&old, &operand)
	default:
		logrus.Fatalf("unknown operations in monket: %v", m)
	}
	if relief {
		m.items[0] = *m.items[0].Quo(&m.items[0], big.NewFloat(3))
		m.items[0] = *m.items[0].SetMode(big.ToZero)
	}
	// increment monkey business
	m.itemsInspected++
}

func (m *monkey) testAndThrow(mob map[int]*monkey) {
	var targetMonkey int
	if m.items[0].Quo(&m.items[0], big.NewFloat(float64(m.test))).IsInt() {
		targetMonkey = m.trueOp
	} else {
		targetMonkey = m.falseOp
	}

	var inFlight big.Float
	inFlight, m.items = m.items[0], m.items[1:]
	mob[targetMonkey].items = append(mob[targetMonkey].items, inFlight)
}

func parseInput11(lines []string) map[int]*monkey {
	monkeys := make(map[int]*monkey)
	for i := 0; i < len(lines); i++ {
		common.PrefixOrDie("Monkey ", lines[i])
		input := strings.Split(lines[i], " ")
		mI, err := strconv.Atoi((strings.Replace(input[1], ":", "", -1)))
		common.Check(err)
		monkeys[mI] = &monkey{itemsInspected: 0}

		// get the monkey's items
		i++
		common.PrefixOrDie("  Starting items: ", lines[i])
		input = strings.Split(lines[i], ":")
		for _, item := range strings.Split(input[1], ",") {
			intItem := common.Atoi(item)
			monkeys[mI].items = append(monkeys[mI].items, *big.NewFloat(float64(intItem)))
		}

		//get inspectOp
		i++
		common.PrefixOrDie("  Operation: new = old ", lines[i])
		input = strings.Split(strings.TrimSpace(strings.Split(lines[i], "=")[1]), " ")
		// id operand == old 0 is left as a flag for 'square existing value'
		operand := 0
		if input[2] != "old" {
			operand = common.Atoi(input[2])
		}
		monkeys[mI].inspectOp = struct {
			string
			float64
		}{input[1], float64(operand)}

		// get test
		i++
		common.PrefixOrDie("  Test: divisible by ", lines[i])
		input = strings.Split(strings.TrimSpace(lines[i]), " ")
		monkeys[mI].test = common.Atoi(input[3])

		// get trueOp
		i++
		common.PrefixOrDie("    If true: throw to monkey ", lines[i])
		input = strings.Split(strings.TrimSpace(lines[i]), " ")
		monkeys[mI].trueOp = common.Atoi(input[5])

		// get falseOp
		i++
		common.PrefixOrDie("    If false: throw to monkey ", lines[i])
		input = strings.Split(strings.TrimSpace(lines[i]), " ")
		monkeys[mI].falseOp = common.Atoi(input[5])

		// skip the blank line
		i++
	}

	return monkeys
}

func playRound(monkeys map[int]*monkey, relief bool) map[int]*monkey {
	for i := 0; i < len(monkeys); i++ {
		l := len(monkeys[i].items) // required to maintain count, slice is edited inside the loop
		for j := 0; j < l; j++ {
			monkeys[i].inspect(relief)
			monkeys[i].testAndThrow(monkeys)
		}
	}
	return monkeys
}

func sortByBusyness(monkeys map[int]*monkey) []int {
	keys := make([]int, 0, len(monkeys))
	for key := range monkeys {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return monkeys[keys[i]].itemsInspected > monkeys[keys[j]].itemsInspected
	})
	return keys
}

func (p Day11) PartA(lines []string) any {
	relief := true
	monkeys := parseInput11(lines)
	totalRounds := 20
	for i := 0; i < totalRounds; i++ {
		monkeys = playRound(monkeys, relief)
	}

	// for some reason we're not getting the right monkey business since converting to `big`.
	// sort by busyness
	ordered := sortByBusyness(monkeys)
	// multiply the item inspection count of the 2 busiest monkeys
	monkeyBusiness := monkeys[ordered[0]].itemsInspected * monkeys[ordered[1]].itemsInspected
	return monkeyBusiness
}

func (p Day11) PartB(lines []string) any {
	// relief := false
	// monkeys := parseInput11(lines)
	// totalRounds := 1000
	return "implement_me"
}
