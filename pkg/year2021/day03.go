//lint:file-ignore ST1005 Ignore stupid error nazi-ing
package year2021

import (
	"aocgen/pkg/common"
	"fmt"
	"strings"
)

type Day03 struct {
}

func pivot(lines []string) [][]string {
	pivoted := [][]string{}
	for i := 0; i < len(lines[0]); i++ {
		pivoted = append(pivoted, []string{})
	}
	for _, val := range lines {
		for idx, val2 := range val {
			pivoted[idx] = append(pivoted[idx], string(val2))
		}
	}
	return pivoted
}

func reduce(filter func(string) string, lines []string) string {
	result := lines[:len(lines)-1]
	pivoted := pivot(lines)

	for i := 0; i < len(lines[0]); i++ {
		selector := filter(strings.Join(pivoted[i], ""))
		tmp := []string{}
		for j := 0; j < len(result); j++ {
			if string(result[j][i]) == selector {
				tmp = append(tmp, result[j])
			}
		}

		if len(tmp) == 1 {
			return tmp[0]
		}
		result = tmp
		pivoted = pivot(result)
	}

	if len(result) != 1 {
		common.Check(fmt.Errorf("Expecting exactly 1 result! %v", result))
	}

	return result[0]
}

func (p Day03) PartA(lines []string) any {
	pivoted := pivot(lines)
	gamma := []rune{}
	epsilon := []rune{}

	for _, val := range pivoted {
		bits := strings.Join(val, "")
		if strings.Count(bits, "1") == strings.Count(bits, "0") {
			common.Check(fmt.Errorf("Input counts are equal! %v", val))
		}
		if strings.Count(bits, "1") > strings.Count(bits, "0") {
			gamma = append([]rune(gamma), '1')
			epsilon = append([]rune(epsilon), '0')
		} else {
			gamma = append([]rune(gamma), '0')
			epsilon = append([]rune(epsilon), '1')
		}
	}

	fmt.Printf("Gamma %v\n", common.MustParseInt(string(gamma), 2))
	fmt.Printf("Epsilon %v\n", common.MustParseInt(string(epsilon), 2))

	return fmt.Sprintf("%v", common.MustParseInt(string(gamma), 2)*common.MustParseInt(string(epsilon), 2))
}

func (p Day03) PartB(lines []string) any {
	o2Selector := func(bits string) string {
		ones := strings.Count(bits, "1")
		zeroes := strings.Count(bits, "0")
		if ones >= zeroes {
			return "1"
		} else {
			return "0"
		}
	}

	co2Selector := func(bits string) string {
		ones := strings.Count(bits, "1")
		zeroes := strings.Count(bits, "0")
		if zeroes <= ones {
			return "0"
		} else {
			return "1"
		}
	}

	o2Report := reduce(o2Selector, lines)
	co2Report := reduce(co2Selector, lines)

	fmt.Printf("O2:  %v\n", o2Report)
	fmt.Printf("CO2: %v\n", co2Report)
	return fmt.Sprintf("%v", common.MustParseInt(o2Report, 2)*common.MustParseInt(co2Report, 2))
}
