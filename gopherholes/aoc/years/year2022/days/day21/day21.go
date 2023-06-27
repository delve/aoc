package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Day21 struct{}

type monkey21 struct {
	left, right, op string
}

func (p Day21) PartA(lines []string) any {
	// begin all input parsing, will need to abstract
	part1 := make(map[string]monkey21)
	opsMonkeys := make(map[string]monkey21)
	p1Resolved := make(map[string]int)
	valueMonkeys := make(map[string]string)

	opsMonkeyRE := regexp.MustCompile("([a-z]*): ([a-z]*) (.) ([a-z]*)")
	for _, line := range lines {
		regIn := opsMonkeyRE.FindAllStringSubmatch(line, -1)
		if len(regIn) > 0 {
			// opsMonkey!!
			part1[regIn[0][1]] = monkey21{
				left:  regIn[0][2],
				right: regIn[0][4],
				op:    regIn[0][3],
			}
			opsMonkeys[regIn[0][1]] = monkey21{
				left:  regIn[0][2],
				right: regIn[0][4],
				op:    regIn[0][3],
			}
		} else {
			input := strings.Split(line, ": ")
			p1Resolved[input[0]] = common.Atoi(input[1])
			valueMonkeys[input[0]] = input[1]
		}
	}
	// end input parsing

	for len(part1) > 0 {
		for k, v := range part1 {
			if left, ok := p1Resolved[v.left]; ok {
				if right, ok := p1Resolved[v.right]; ok {
					val := 0
					switch v.op {
					case "+":
						val = left + right
					case "-":
						val = left - right
					case "*":
						val = left * right
					case "/":
						val = left / right
					default:
						logrus.Fatalf("Unknown operator in %s: %s", k, v.op)
					}
					p1Resolved[k] = val
					delete(part1, k)
				}
			}
		}
	}

	fmt.Printf("Root: %d\n", p1Resolved["root"])
	fmt.Println("Part 2\n-----")
	opsMonkeys["root"] = monkey21{
		left:  opsMonkeys["root"].left,
		right: opsMonkeys["root"].right,
		op:    "=",
	}
	// i am ashamed to admit that i used an online algebra solver (https://mathsolver.microsoft.com)
	//  to deal with all those parenthesis :(
	valueMonkeys["humn"] = "3296135418820"
	// delete(valueMonkeys, "humn")
	simplify := true
	for simplify {
		simplify = false
		for k, v := range opsMonkeys {
			if k == "root" {
				continue
			}
			nLeft := v.left
			nRight := v.right
			if _, ok := valueMonkeys[v.left]; ok {
				nLeft = valueMonkeys[v.left]
			}
			if _, ok := valueMonkeys[v.right]; ok {
				nRight = valueMonkeys[v.right]
			}
			iLeft, lErr := strconv.Atoi(nLeft)
			iRight, rErr := strconv.Atoi(nRight)
			if lErr == nil && rErr == nil {
				simplify = true
				val := 0
				switch v.op {
				case "+":
					val = iLeft + iRight
				case "-":
					val = iLeft - iRight
				case "*":
					val = iLeft * iRight
				case "/":
					val = iLeft / iRight
				default:
					logrus.Fatalf("Unknown operator in %s: %s", k, v.op)
				}
				valueMonkeys[k] = fmt.Sprintf("%d", val)
				delete(opsMonkeys, k)
			} else {
				opsMonkeys[k] = monkey21{
					left:  nLeft,
					right: nRight,
					op:    v.op,
				}
			}
		}
	}

	for deadKey, deadMonkey := range opsMonkeys {
		if deadKey == "root" {
			continue
		}
		calc := fmt.Sprintf("(%s %s %s", deadMonkey.left, deadMonkey.op, deadMonkey.right)
		for updateKey, updateMonkey := range opsMonkeys {
			nLeft := strings.ReplaceAll(updateMonkey.left, deadKey, calc)
			nRight := strings.ReplaceAll(updateMonkey.right, deadKey, calc)
			opsMonkeys[updateKey] = monkey21{
				left:  nLeft,
				right: nRight,
				op:    opsMonkeys[updateKey].op,
			}
		}
	}
	fmt.Printf("opsMonkeys: %d %v\n\n", len(opsMonkeys), opsMonkeys)
	fmt.Printf("valueMonkeys: %d %v\n\n", len(valueMonkeys), valueMonkeys)
	fmt.Println(opsMonkeys["root"])
	fmt.Printf("\n\n\n%s = %s", valueMonkeys[opsMonkeys["root"].left], valueMonkeys[opsMonkeys["root"].right])
	fmt.Printf("\n\n%s %s %s\n", opsMonkeys["root"].left, opsMonkeys["root"].op, valueMonkeys[opsMonkeys["root"].right])

	return "implement_me"
}

func (p Day21) PartB(lines []string) any {
	return "implement_me"
}
