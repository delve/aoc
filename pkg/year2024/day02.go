package year2024

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strings"
)

type Day02 struct{}

func checkReport(report []string) (bool, int, string) {
	safeReport := true
	monotonicity := 0
	comment := "Safe"
	if common.Atoi(report[0])-common.Atoi(report[len(report)-1]) < 1 {
		monotonicity = 1
	} else {
		monotonicity = -1
	}
	badLevel := -1
	for idx, level := range report {
		if idx == 0 {
			continue
		}
		val := common.Atoi(level)
		change := val - common.Atoi(report[idx-1])
		if float64(change)/math.Abs(float64(change)) != float64(monotonicity) {
			safeReport = false
			badLevel = idx
			comment = fmt.Sprintf("\t%s->%s not monotonic %d", report[idx-1], level, monotonicity)
			break
		}
		if math.Abs(float64(change)) > 3 {
			safeReport = false
			badLevel = idx
			comment = fmt.Sprintf("\t%s->%s change %d greater than 3", report[idx-1], level, change)
			break
		}
	}
	return safeReport, badLevel, comment
}

func (p Day02) PartA(lines []string) any {
	safeCount := 0
	for _, inputs := range lines[:len(lines)-1] {
		report := strings.Split(inputs, " ")
		safe, _, _ := checkReport(report)
		// fmt.Printf("report %d: %v [%d]\n", idx, safe, badLevel)
		if safe {
			safeCount++
		}
	}
	return fmt.Sprintf("%v", safeCount)
}

func (p Day02) PartB(lines []string) any {
	safeCount := 0
	for _, inputs := range lines[:len(lines)-1] {
		report := strings.Split(inputs, " ")
		levelCount := len(report)
		safe, badLevel, comment := checkReport(report)
		// fmt.Printf("report %d: %v [%d]\n", idx, safe, badLevel)
		if !safe {
			report = append(report[:badLevel], report[badLevel+1:]...)
			safe, _, comment = checkReport(report)
		}
		if !safe {
			report = strings.Split(inputs, " ")[1:]
			safe, _, comment = checkReport(report)
		}
		if !safe {
			report = strings.Split(inputs, " ")[:levelCount-1]
			safe, _, comment = checkReport(report)
		}

		if safe {
			safeCount++
		} else {
			fmt.Printf("Unsafe: %v\n%s\n", report, comment)
		}
	}
	return fmt.Sprintf("%v", safeCount)
}
