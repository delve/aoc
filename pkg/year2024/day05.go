package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Day05 struct {
	rules   map[int]map[int]bool
	updates [][]int
}

func (p *Day05) parseInput(lines []string) {
	p.rules = map[int]map[int]bool{}
	p.updates = [][]int{}
	breakLine := 0
	ruleRe := regexp.MustCompile(`(\d+)\|(\d+)`)
	for idx, input := range lines {
		breakLine = idx
		if len(strings.TrimSpace(input)) == 0 {
			break
		}
		if match := ruleRe.FindSubmatch([]byte(input)); match != nil {
			ruleKey := common.Atoi(string(match[1]))
			ruleVal := common.Atoi(string(match[2]))
			if _, contained := p.rules[ruleKey]; !contained {
				p.rules[ruleKey] = map[int]bool{}
			}
			p.rules[ruleKey][ruleVal] = true
		} else {
			panic(fmt.Errorf("no regex match: %v", input))
		}
	}

	for _, input := range lines[breakLine+1 : len(lines)-1] {
		list := strings.Split(input, ",")
		set := []int{}
		for _, val := range list {
			set = append(set, common.Atoi(val))
		}
		p.updates = append(p.updates, set)
	}
}

func isOrdered(update []int, rules map[int]map[int]bool) bool {
	positions := map[int]int{}
	for idx, val := range update {
		positions[val] = idx
	}
	for lowVal, highVals := range rules {
		if _, contained := positions[lowVal]; contained {
			for highVal := range highVals {
				if _, contained := positions[highVal]; contained {
					if positions[lowVal] > positions[highVal] {
						return false
					}
				}
			}
		}
	}
	return true
}

func fixOrder(update []int, rules map[int]map[int]bool) []int {
	// positionByVal := map[int]*list.Element{}
	// set := list.New()
	// for _, pag := range update {
	// 	positionByVal[pag] = set.PushBack(pag)
	// }

	// for lowVal, highVals := range rules {
	// 	if _, contained := positionByVal[lowVal]; contained {
	// 		for highVal := range highVals {
	// 			if _, contained := positionByVal[highVal]; contained {
	// 				if positionByVal[lowVal] > positionByVal[highVal]  {
	// 					for i := 0; i < positionByVal[lowVal] - positionByVal[highVal]; i++ {

	// 						valByPosition[positionByVal[lowVal]-i] = valByPosition[positionByVal[lowVal]-i-1]
	// 						positionByVal[valByPosition[positionByVal[lowVal]-i]] = positionByVal[positionByVal[lowVal]-i]+1
	// 					}
	// 					valByPosition[positionByVal[highVal]] = lowVal
	// 					positionByVal[valByPosition[positionByVal[highVal]]] = positionByVal[highVal]
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	positionByVal := map[int]int{}
	valByPosition := map[int]int{}
	for idx, val := range update {
		positionByVal[val] = idx
		valByPosition[idx] = val
	}

	shiftLeft := func(positionByVal map[int]int, valByPosition map[int]int, targetPos int, distance int) (posMap map[int]int, valMap map[int]int) {
		targetVal := valByPosition[targetPos]

		for i := 0; i < distance; i++ {

			valByPosition[targetPos-i] = valByPosition[targetPos-i-1]
			positionByVal[valByPosition[targetPos-i]] = targetPos - i
		}
		// set the last position
		valByPosition[targetPos-distance] = targetVal
		positionByVal[targetVal] = targetPos - distance

		return positionByVal, valByPosition
	}

	for lowVal, highVals := range rules {
		if _, contained := positionByVal[lowVal]; contained {
			for highVal := range highVals {
				if _, contained := positionByVal[highVal]; contained {
					if positionByVal[lowVal] > positionByVal[highVal] {
						shiftLeft(positionByVal, valByPosition, positionByVal[lowVal], positionByVal[lowVal]-positionByVal[highVal])
					}
				}
			}
		}
	}

	output := []int{}
	for i := 0; i < len(update); i++ {
		output = append(output, valByPosition[i])
	}
	return output
}

func (p Day05) PartA(lines []string) any {
	p.parseInput(lines)
	sum := 0
	for _, update := range p.updates {
		if isOrdered(update, p.rules) {
			mid := int(math.Ceil(float64(len(update))/2.0)) - 1 //zero based position
			sum += update[mid]
		}
	}
	return sum
}

func (p Day05) PartB(lines []string) any {
	p.parseInput(lines)
	sum := 0
	for _, update := range p.updates {
		if !isOrdered(update, p.rules) {
			newUpdate := fixOrder(update, p.rules)
			mid := int(math.Ceil(float64(len(newUpdate))/2.0)) - 1 //zero based position
			sum += newUpdate[mid]
		}
	}
	return sum
}
