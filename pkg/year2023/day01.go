package year2023

import (
	"aocgen/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	lineDigits := make([][]rune, len(lines))

	for lineI, line := range lines {
		for _, rne := range line {
			if _, err := strconv.Atoi(string(rne)); err == nil {
				lineDigits[lineI] = append(lineDigits[lineI], rne)
			}
		}
	}

	calVals := getCalVals(lineDigits)
	// why is Sum not defined? the docs say it exists??? >:|
	// return math.Sum(calVals)
	retval := 0
	for _, v := range calVals {
		retval += v
	}
	return retval
}

func (p Day01) PartB(lines []string) any {
	lineDigits := make([][]rune, len(lines))

	for lineI, line := range lines {
		for i, rne := range line {
			if _, err := strconv.Atoi(string(rne)); err == nil {
				// if it's a digit then grab it
				lineDigits[lineI] = append(lineDigits[lineI], rne)
			} else {
				switch {
				case strings.HasPrefix(line[i:], "one"):
					lineDigits[lineI] = append(lineDigits[lineI], '1')
				case strings.HasPrefix(line[i:], "two"):
					lineDigits[lineI] = append(lineDigits[lineI], '2')
				case strings.HasPrefix(line[i:], "three"):
					lineDigits[lineI] = append(lineDigits[lineI], '3')
				case strings.HasPrefix(line[i:], "four"):
					lineDigits[lineI] = append(lineDigits[lineI], '4')
				case strings.HasPrefix(line[i:], "five"):
					lineDigits[lineI] = append(lineDigits[lineI], '5')
				case strings.HasPrefix(line[i:], "six"):
					lineDigits[lineI] = append(lineDigits[lineI], '6')
				case strings.HasPrefix(line[i:], "seven"):
					lineDigits[lineI] = append(lineDigits[lineI], '7')
				case strings.HasPrefix(line[i:], "eight"):
					lineDigits[lineI] = append(lineDigits[lineI], '8')
				case strings.HasPrefix(line[i:], "nine"):
					lineDigits[lineI] = append(lineDigits[lineI], '9')

				}
			}
		}
	}

	calVals := getCalVals(lineDigits)
	// why is Sum not defined? the docs say it exists??? >:|
	// return math.Sum(calVals)
	retval := 0
	for _, v := range calVals {
		retval += v
	}
	return retval

}

func getCalVals(lineDigits [][]rune) (calVals []int) {
	calVals = []int{}
	for _, line := range lineDigits {
		val := ""
		// it's literal first and last. if they're the same it's duplicated
		if len(line) >= 1 {
			val = fmt.Sprintf("%s%s", string(line[0]), string(line[len(line)-1]))
		} else {
			// protect from panics in part 2 sample with no embedded ints
			val = "0"
		}
		calVals = append(calVals, common.Atoi(val))
	}

	return calVals
}
