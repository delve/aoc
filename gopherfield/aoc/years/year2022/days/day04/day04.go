package year2022

import (
	"aocgen/pkg/common"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sirupsen/logrus"
)

type Day04 struct{}

type regionPair struct {
	elf1, elf2 mapset.Set[int]
	intersect  mapset.Set[int]
}

func makeRange(min, max int) []int {
	r := make([]int, max-min+1)
	for i := range r {
		r[i] = min + i
	}
	return r
}

func parseLine04(s string) ([]int, []int) {
	regions := strings.Split(s, ",")
	if len(regions) > 2 {
		logrus.Fatalf("found more than 2 regions in %s", s)
	}

	r := make([][]int, 2)
	for i := 0; i < 2; i++ {
		bounds := strings.Split(regions[i], "-")
		if len(bounds) > 2 {
			logrus.Fatalf("found more than 2 numbers in %s", bounds)
		}

		iBounds := make([]int, 2)
		for i := 0; i < 2; i++ {
			e := *new(error)
			iBounds[i], e = strconv.Atoi(bounds[i])
			common.Check(e)
		}
		r[i] = makeRange(iBounds[0], iBounds[1])
	}
	return r[0], r[1]
}

func parseInput04(lines []string) []regionPair {
	dataAccumulator := []regionPair{}

	for _, line := range lines {
		pair := regionPair{}
		elf1, elf2 := parseLine04(line)

		pair.elf1 = mapset.NewSet(elf1...)
		pair.elf2 = mapset.NewSet(elf2...)
		pair.intersect = pair.elf1.Intersect(pair.elf2)

		dataAccumulator = append(dataAccumulator, pair)
	}
	return dataAccumulator
}

func (p Day04) PartA(lines []string) any {
	dataAccumulator := parseInput04(lines)

	superSets := 0
	for i := 0; i < len(dataAccumulator); i++ {
		inter := dataAccumulator[i].intersect.Cardinality()
		if dataAccumulator[i].elf1.Cardinality() == inter || dataAccumulator[i].elf2.Cardinality() == inter {
			superSets++
		}
	}
	return superSets
}

func (p Day04) PartB(lines []string) any {
	dataAccumulator := parseInput04(lines)

	intersections := 0
	for i := 0; i < len(dataAccumulator); i++ {
		if dataAccumulator[i].intersect.Cardinality() > 0 {
			intersections++
		}
	}
	return intersections
}
