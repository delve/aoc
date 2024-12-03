package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
)

type Day03 struct{}

func (p Day03) PartA(lines []string) any {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	for _, line := range lines {
		instructions := re.FindAllStringSubmatchIndex(line, -1)
		for _, instruction := range instructions {
			fmt.Printf("%s * %s\n", line[instruction[2]:instruction[3]], line[instruction[4]:instruction[5]])
			sum += common.Atoi(line[instruction[2]:instruction[3]]) * common.Atoi(line[instruction[4]:instruction[5]])
		}
	}
	return sum
}

func (p Day03) PartB(lines []string) any {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	for _, line := range lines {
		instructions := re.FindAllStringSubmatchIndex(line, -1)
		for _, instruction := range instructions {
			if line[instruction[0]:instruction[1]] == "don't()" {
				enabled = false
			}
			if line[instruction[0]:instruction[1]] == "do()" {
				enabled = true
			}
			if enabled && line[instruction[0]:instruction[1]][:3] == "mul" {
				num1 := common.Atoi(line[instruction[2]:instruction[3]])
				num2 := common.Atoi(line[instruction[4]:instruction[5]])
				fmt.Printf("%d * %d\n", num1, num2)
				sum += num1 * num2
			}
		}
	}
	return sum
}
