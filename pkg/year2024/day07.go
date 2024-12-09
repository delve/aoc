package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"strings"
)

type Day07 struct {
	calibrations map[int][]int
}

func (p *Day07) parseInput(inputs []string) {
	p.calibrations = map[int][]int{}
	for _, equation := range inputs {
		colon := strings.IndexRune(equation, ':')
		testVal := common.Atoi(equation[:colon])
		p.calibrations[testVal] = []int{}
		for _, val := range strings.Split(equation[colon+2:], " ") {
			p.calibrations[testVal] = append(p.calibrations[testVal], common.Atoi(val))
		}
	}
}

func (p Day07) checkValidity(testVal int, inputs []int, equation string) (bool, string) {
	if len(inputs) == 1 && inputs[0] == testVal {
		return true, equation
	}
	if len(inputs) == 1 || inputs[0] > testVal {
		return false, equation
	}

	nextIter := make([]int, len(inputs)-1)
	copy(nextIter, inputs[1:])

	nextIter[0] = inputs[0] * inputs[1]
	if ok, equation := p.checkValidity(testVal, nextIter, equation); ok {
		equation = "*" + equation
		return true, equation
	}

	nextIter[0] = inputs[0] + inputs[1]
	if ok, equation := p.checkValidity(testVal, nextIter, equation); ok {
		equation = "+" + equation
		return true, equation
	}

	return false, equation
}

func (p Day07) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	sum := 0
	for testVal, nums := range p.calibrations {
		sanitySum := 1
		for _, num := range nums {
			sanitySum = sanitySum * num
		}
		if sanitySum < testVal { // largest possible result is too small, skip it
			fmt.Printf("too small: %d > %s = %d\n", testVal, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), "*"), "[]"), sanitySum)
			continue
		}
		sanitySum = 0
		for _, num := range nums {
			sanitySum = sanitySum + num
		}
		if sanitySum > testVal { // smallest possible result is too large, skip it
			fmt.Printf("too large: %d < %s = %d\n", testVal, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), "+"), "[]"), sanitySum)
			continue
		}
		if ok, equation := p.checkValidity(testVal, nums, ""); ok {
			out := fmt.Sprintf("%d = ", testVal)
			testSum := nums[0]
			for idx, op := range equation {
				out += fmt.Sprintf("%d %s ", nums[idx], string(op))
				if op == '+' {
					testSum += nums[idx+1]
				}
				if op == '*' {
					testSum = testSum * nums[idx+1]
				}
			}
			out += fmt.Sprintf("%d = %d", nums[len(nums)-1], testSum)
			if testSum != testVal {
				out += "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
				println(out)
				panic("!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			}
			fmt.Printf("Valid: %s\n", out)
			sum += testVal
		} else {
			fmt.Printf("impossible: %d = %v\n", testVal, nums)
		}
	}
	if sum == 5751665302886 {
		fmt.Printf("\n\nnot it (high? low? who knows!)\n\n")
	}
	return sum
}

func (p Day07) PartB(lines []string) any {
	return "implement_me"
}
