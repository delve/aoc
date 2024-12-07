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
	// i don't understand why inputs[2:] isn't causing an out of bounds error here when len(inputs) == 2
	if ok, equation := p.checkValidity(testVal, append([]int{inputs[0] * inputs[1]}, inputs[2:]...), equation); ok {
		equation = "*" + equation
		return true, equation
	}
	if ok, equation := p.checkValidity(testVal, append([]int{inputs[0] + inputs[1]}, inputs[2:]...), equation); ok {
		equation = "+" + equation
		return true, equation
	}

	return false, equation
}

func (p Day07) PartA(lines []string) any {
	p.parseInput(lines[:len(lines)-1])
	sum := 0
	for testVal, nums := range p.calibrations {
		sanitySum := nums[0]
		for _, num := range nums[1:] {
			sanitySum = sanitySum * num
		}
		if sanitySum < testVal { // largest possible result is too small, skip it
			continue
		}
		sanitySum = nums[0]
		for _, num := range nums[1:] {
			sanitySum = sanitySum + num
		}
		if sanitySum > testVal { // smallest possible result is too large, skip it
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
			println(out)
			sum += testVal
		} else {
			fmt.Printf("%d = %v\n", testVal, nums)
		}
	}
	return sum
}

func (p Day07) PartB(lines []string) any {
	return "implement_me"
}
