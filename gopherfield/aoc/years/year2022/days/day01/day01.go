package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"sort"
	"strconv"
	"unicode/utf8"
)

type Day01 struct{}

func getElves(l []string) []int {
	elfCalories := []int{}

	currentCalories := 0
	for _, v := range l {
		if utf8.RuneCountInString(v) > 0 {
			// accumulate this elf's calories
			calorieEntry, err := strconv.Atoi(v)
			common.Check(err)
			currentCalories += calorieEntry
		} else {
			// blank line indicates a new elf, store & reset
			elfCalories = append(elfCalories, currentCalories)
			currentCalories = 0
		}
	}
	return elfCalories
}

func (p Day01) PartA(lines []string) any {
	elves := getElves(lines)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	ret := fmt.Sprintf("Total entries: %d\n", len(lines))
	ret += fmt.Sprintf("Elves: %d\n", len(elves))
	ret += fmt.Sprintf("Top calories: %d", elves[0])
	return ret
}

func (p Day01) PartB(lines []string) any {
	elves := getElves(lines)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	return fmt.Sprintf("Sum of top 3 elves: %d", elves[0]+elves[1]+elves[2])
}
