package year2023

import (
	"fmt"
	"math"
	"regexp"
)

type Day08 struct{}

type node struct {
	left, right string
}

func (p Day08) PartA(lines []string) any {
	return 0
	steps := 0
	directions := lines[0]
	network := p.parseNetwork(lines[2:])
	curNode := "AAA"
	// fmt.Printf("%v\n", network)
	for i := 0; true; i++ {
		dir := directions[int(math.Mod(float64(i), float64(len(directions))))]
		switch dir {
		case 'R':
			curNode = network[curNode].right
			// fmt.Printf("%v: Right to %v\n", i, curNode)
		case 'L':
			curNode = network[curNode].left
			// fmt.Printf("%v: Left to %v\n", i, curNode)
		}
		if curNode == "ZZZ" {
			steps = i + 1
			break
		}
	}
	return steps
}

func (p Day08) PartB(lines []string) any { // currently never seems to terminate.
	steps := 0
	directions := lines[0]
	network, curNodes := p.parseNetwork2(lines[2:])
	// fmt.Printf("%v\n%v\n", network, curNodes)
	fmt.Printf("%v\n", curNodes)
	// number that's too high as determined by submitting guesses
	cap := 1000000000000000
TRAVERSENETWORK:
	for i := 0; i < cap; i++ {
		if math.Mod(float64(i), float64(cap/1000)) == 0 {
			fmt.Printf("%.2f%% complete Directions looped %v times. Steps: %v\n", 100*(float64(i)/float64(cap)), i/len(directions), i+1)
		}
		dir := directions[int(math.Mod(float64(i), float64(len(directions))))]
		switch dir {
		case 'R':
			for i, curNode := range curNodes {
				curNodes[i] = network[curNode].right
			}
			// fmt.Printf("%v: Right to %v\n", i, curNode)
		case 'L':
			for i, curNode := range curNodes {
				curNodes[i] = network[curNode].left
			}
			// fmt.Printf("%v: Left to %v\n", i, curNode)
		}
		for _, curNode := range curNodes {
			if curNode[2] != 'Z' {
				continue TRAVERSENETWORK
			}
			// fmt.Printf("%v\n", curNodes)
		}
		// we slipped out of the above for, meaning everything ends with Z
		steps = i + 1
		break
	}
	return steps
}

func (p Day08) parseNetwork(lines []string) (network map[string]node) {
	network = map[string]node{}
	r := regexp.MustCompile("([A-Z]{3})")
	for _, line := range lines {
		nodeMatch := r.FindAllStringSubmatch(line, -1)
		network[nodeMatch[0][0]] = node{left: nodeMatch[1][0], right: nodeMatch[2][0]}
	}
	return network
}

func (p Day08) parseNetwork2(lines []string) (network map[string]node, startNodes []string) {
	network = map[string]node{}
	startNodes = []string{}
	r := regexp.MustCompile("([A-Z]{3})")
	for _, line := range lines {
		nodeMatch := r.FindAllStringSubmatch(line, -1)
		network[nodeMatch[0][0]] = node{left: nodeMatch[1][0], right: nodeMatch[2][0]}
		if nodeMatch[0][0][2] == 'A' { // if the node name ends with A then it's a starting node, assumes names are always 3 letters
			startNodes = append(startNodes, nodeMatch[0][0])
		}
	}
	return network, startNodes
}
