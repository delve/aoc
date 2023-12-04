package year2023

import (
	"aocgen/pkg/common"
	"regexp"
)

type Day03 struct{}

func (p Day03) PartA(lines []string) any {
	r := regexp.MustCompile(`(?P<number>\d+)`)
	symbol := regexp.MustCompile(`[^\d.]`)

	partSum := 0
	for inputIndex, line := range lines {
		matches := r.FindAllStringIndex(line, -1)
		// fmt.Printf("%#v\n", matches)
		// return slice indices, use to test for part and capture number if needed
		for _, match := range matches {
			leftBoundInc, rightBoundExc := 0, len(line)
			symbolsFound := 0
			if match[0] > 0 {
				leftBoundInc = match[0] - 1
			}
			if match[1] < len(line) {
				rightBoundExc = match[1] + 1
			}
			if inputIndex > 0 {
				// check the previous line for symbols
				symbolsFound += len(symbol.FindAllString(lines[inputIndex-1][leftBoundInc:rightBoundExc], -1))
			}
			if inputIndex < len(lines)-1 {
				// check the next line for symbols
				symbolsFound += len(symbol.FindAllString(lines[inputIndex+1][leftBoundInc:rightBoundExc], -1))
			}
			// check current line for symbols
			symbolsFound += len(symbol.FindAllString(line[leftBoundInc:rightBoundExc], -1))

			if symbolsFound > 0 {
				// fmt.Printf("%s", line[match[0]:match[1]])
				partSum += common.Atoi(line[match[0]:match[1]])
			}
		}
	}

	return partSum
}

func (p Day03) PartB(lines []string) any {
	gearRatioSum := 0
	star := regexp.MustCompile(`\*`)

	for inputIndex, line := range lines {
		gearCandidates := star.FindAllStringIndex(line, -1)
		for _, candidate := range gearCandidates {
			leftBoundInc, rightBoundExc := 0, len(line)
			parts := []int{}
			if candidate[0] > 0 {
				leftBoundInc = candidate[0] - 1
			}
			if candidate[1] < len(line) {
				rightBoundExc = candidate[1] + 1
			}
			gearRange := [2]int{leftBoundInc, rightBoundExc}
			if inputIndex > 0 {
				parts = append(parts, checkAdjacentParts(lines[inputIndex-1], [2]int(gearRange))...)
			}
			if inputIndex < len(lines)-1 {
				parts = append(parts, checkAdjacentParts(lines[inputIndex+1], [2]int(gearRange))...)
			}
			parts = append(parts, checkAdjacentParts(line, [2]int(gearRange))...)

			if len(parts) == 2 {
				gearRatioSum += parts[0] * parts[1]
			}
		}
	}

	return gearRatioSum
}

func checkAdjacentParts(line string, gearRange [2]int) []int {
	partNums := []int{}

	// look for numbers overlapping gearRange, return it/them as int(s)
	number := regexp.MustCompile(`(?P<number>\d+)`)
	parts := number.FindAllStringIndex(line, -1)
	for _, part := range parts {
		// THIS DOES NOT CATCH long numbers where the ends are both outside the range
		if part[1] > gearRange[0] && part[1] <= gearRange[1] ||
			part[0] >= gearRange[0] && part[0] < gearRange[1] ||
			part[0] <= gearRange[0] && part[1] >= gearRange[1] {
			partNums = append(partNums, common.Atoi(line[part[0]:part[1]]))
		}
	}

	return partNums
}
