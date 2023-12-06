package year2023

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
	"strings"
)

type Day06 struct {
}

type input struct {
	times, distances []string
}

func (p Day06) PartA(lines []string) any {
	inputs := p.parseInput(lines)

	marginOfError := 1
	fmt.Printf("%v\n%v\n", inputs.times, inputs.distances)
	for idx, time := range inputs.times {
		raceMargin := p.bruteForceIt(common.Atoi(time), common.Atoi(inputs.distances[idx]))
		// fmt.Printf("%v\n", raceMargin)
		marginOfError = marginOfError * raceMargin
	}
	return marginOfError
}

func (p Day06) PartB(lines []string) any {
	inputs := p.parseInput(lines)

	fmt.Printf("%v\n%v\n", inputs.times, inputs.distances)
	rTime := strings.Join(inputs.times[:], "")
	rRecord := strings.Join(inputs.distances[:], "")
	fmt.Printf("%v\n%v\n", rTime, rRecord)
	marginOfError := p.bruteForceIt(common.Atoi(rTime), common.Atoi(rRecord))
	return marginOfError
}

func (p Day06) parseInput(lines []string) input {
	r := regexp.MustCompile(`(?P<number>\d+)`)
	return input{
		times:     r.FindAllString(lines[0], -1),
		distances: r.FindAllString(lines[1], -1),
	}
}

func (p Day06) bruteForceIt(timeLimit, distanceRecord int) (margin int) {
	margin = 0

	for i := 0; i < timeLimit; i++ {
		if i*(timeLimit-i) > distanceRecord {
			margin = i
			break
		}
	}
	for i := timeLimit; i > 0; i-- {
		if i*(timeLimit-i) > distanceRecord {
			margin = i + 1 - margin
			break
		}
	}

	return margin
}
