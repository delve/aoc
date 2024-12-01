package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
	leftList    []float64
	rightList   []float64
	leftString  string
	rightString string
}

func (p *Day01) parseInput(lines []string) {
	left := []string{}
	right := []string{}
	for _, line := range lines[:len(lines)-1] {
		ids := strings.Split(line, "   ")
		if id, err := strconv.ParseFloat(ids[0], 64); err == nil {
			p.leftList = append(p.leftList, id)
			left = append(left, ids[0])
		}
		if id, err := strconv.ParseFloat(ids[1], 64); err == nil {
			p.rightList = append(p.rightList, id)
			right = append(right, ids[1])
		}
	}

	p.leftString = strings.Join(left, " ")
	p.rightString = strings.Join(right, " ")
	sort.Float64s(p.leftList)
	sort.Float64s(p.rightList)
	if len(p.leftList) != len(p.rightList) {
		panic(fmt.Errorf("fuck you go i'll capitalize this if i want.\nLength mismatch\n%v\n%v", p.leftList, p.rightList))
	}
}

func (p Day01) PartA(lines []string) any {
	p.parseInput(lines)
	// fmt.Printf("%v\n%v", p.leftList, p.rightList)
	distance := 0.0
	for idx := range p.leftList {
		distance += math.Abs(p.leftList[idx] - p.rightList[idx])
	}

	return fmt.Sprintf("%f\n", distance)
}

func (p Day01) PartB(lines []string) any {
	p.parseInput(lines)
	similarity := 0
	iterator := strings.Split(p.leftString, " ")
	for _, val := range iterator {
		count := strings.Count(p.rightString, val)
		similarity += count * common.Atoi(val)
	}
	return similarity
}
