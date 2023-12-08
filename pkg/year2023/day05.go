package year2023

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Day05 struct{}

type maperator struct {
	sources      []int
	destinations []int
	ranges       []int
}

func (p Day05) PartA(lines []string) any {
	inputSplit := strings.Split(lines[0], " ")
	seeds := []int{}
	for _, seedStr := range inputSplit[1:] {
		seeds = append(seeds, common.Atoi(seedStr))
	}
	// fmt.Printf("%v\n", seeds)

	maps := []maperator{}
	cursor := 2
	for i := 0; i < 7; i++ {
		// find lines slice containing map
		start := cursor + 1
		for ; len(lines[cursor]) > 0; cursor++ {
		}

		// parse map into maps
		maps = append(maps, parseMap(lines[start:cursor]))
		cursor++
	}

	seedLocs := []int{}
	for _, seed := range seeds {
		location := seed
		for _, almanac := range maps {
			location = almanac.translate(location)
		}
		seedLocs = append(seedLocs, location)
	}

	slices.Sort(seedLocs)
	// fmt.Printf("%v\n", seedLocs)

	return seedLocs[0]
}

// func (p Day05) deprecated_PartB(lines []string) any {
// 	inputSplit := strings.Split(lines[0], " ")
// 	seedInts := []int{}
// 	for i := 1; i < len(inputSplit); i++ {
// 		seedInts = append(seedInts, common.Atoi(inputSplit[i]))
// 	}

// 	seeds := []int{}
// 	for i := 0; i < len(seedInts); i = i + 2 {
// 		fmt.Printf("Growing slice by %v.\n", seedInts[i+1])
// 		seeds = slices.Grow(seeds, seedInts[i+1])
// 		for j := seedInts[i]; j < seedInts[i]+seedInts[i+1]; j++ {
// 			seeds = append(seeds, j)
// 		}
// 	}

// 	fmt.Printf("Got %v seeds.\n", len(seeds))

// 	maps := []maperator{}
// 	cursor := 2
// 	for i := 0; i < 7; i++ {
// 		fmt.Printf("Starting map %v.\n", i)
// 		// find lines slice containing map
// 		start := cursor + 1
// 		for ; len(lines[cursor]) > 0; cursor++ {
// 		}

// 		// parse map into maps
// 		maps = append(maps, parseMap(lines[start:cursor]))
// 		cursor++
// 	}

// 	smallest := math.MaxInt32
// 	// seedLocs := []int{}
// 	for _, seed := range seeds {
// 		location := seed
// 		for _, almanac := range maps {
// 			// fmt.Printf("%v\n", location)
// 			location = almanac.translate(location)
// 		}
// 		// seedLocs = append(seedLocs, location)
// 		if location < smallest {
// 			smallest = location
// 		}
// 	}

// 	// slices.Sort(seedLocs)
// 	// fmt.Printf("%v\n", seedLocs)

// 	return smallest
// }

func (p Day05) PartB(lines []string) any {
	inputSplit := strings.Split(lines[0], " ")
	seedInts := []int{}
	for i := 1; i < len(inputSplit); i++ {
		seedInts = append(seedInts, common.Atoi(inputSplit[i]))
	}
	maps := []maperator{}
	cursor := 2
	for i := 0; i < 7; i++ {
		// find lines slice containing map
		start := cursor + 1
		for ; len(lines[cursor]) > 0; cursor++ {
		}

		// parse map into maps
		maps = append(maps, parseMap(lines[start:cursor]))
		cursor++
	}

	smallest := math.MaxInt32
	fmt.Printf("%v seed groups.\n", len(seedInts)/2)
	for i := 0; i < len(seedInts); i = i + 2 {
		fmt.Printf("Searching group %v [%v - %v]. Smallest so far: %v\n", i/2, seedInts[i], seedInts[i]+seedInts[i+1], smallest)

		for seed := seedInts[i]; seed < seedInts[i]+seedInts[i+1]; seed++ {
			location := seed
			for _, almanac := range maps {
				// fmt.Printf("%v\n", location)
				location = almanac.translate(location)
			}
			// seedLocs = append(seedLocs, location)
			if location < smallest {
				smallest = location
			}
		}
	}

	return smallest
}

func parseMap(input []string) (almanac maperator) {
	almanac = maperator{sources: []int{}, destinations: []int{}, ranges: []int{}}

	for _, line := range input {
		tuple := strings.Split(line, " ")
		almanac.sources = append(almanac.sources, common.Atoi(tuple[1]))
		almanac.destinations = append(almanac.destinations, common.Atoi(tuple[0]))
		almanac.ranges = append(almanac.ranges, common.Atoi(tuple[2]))
	}
	return almanac
}

func (m maperator) translate(input int) int {
	value := input

	for idx, source := range m.sources {
		if input >= source && input <= source+m.ranges[idx] {
			offset := input - source
			value = m.destinations[idx] + offset
			break
		}

	}
	return value
}
