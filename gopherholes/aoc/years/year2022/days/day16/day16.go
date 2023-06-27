package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
	"strings"

	"github.com/dominikbraun/graph"
)

type Day16 struct{}

type valve struct {
	id             int // probably not actually use, leftover from a different graph lib
	name           string
	rate, openedAt int
	opened         bool
	tunnels        []string
}

func (p Day16) PartA(lines []string) any {
	valves := make(map[string]valve)

	re := regexp.MustCompile("Valve ([A-Z]{2}) has flow rate=([-0-9]+).*valves? (.*)")

	for _, line := range lines {
		regIn := re.FindAllStringSubmatch(line, -1)
		input := regIn[0]

		rate := common.Atoi(input[2])
		tunnels := strings.Split(input[3], ", ")
		valves[input[1]] = valve{
			id:       len(valves),
			name:     input[1],
			rate:     rate,
			openedAt: -1,
			opened:   false,
			tunnels:  tunnels,
		}
	}

	vGraph := graph.New(graph.StringHash)

	for k, v := range valves {
		vGraph.AddVertex(k)
		for _, to := range v.tunnels {
			vGraph.AddEdge(k, to)
		}
	}

	// file,_:=os.Create("./mygraph.gv")
	// _ = draw.DOT(vGraph,file)

	fullPath := ""
	released := 0
	start := "AA"
	totalTime := 30
	candidate := ""
	for i := 0; i < totalTime; i++ {
		expPress := 0
		pathTaken := []string{}
		timeUsed := 0
		for k, v := range valves {
			path, _ := graph.ShortestPath(vGraph, start, k)
			timeUsed = len(path)
			vPress := v.rate * (totalTime - i - timeUsed)
			fmt.Printf("time %d: %s to %s via %v steps %d, rate %d for pressure tot %d", i, start, k, path, timeUsed, v.rate, vPress)
			if vPress > expPress && timeUsed < totalTime-i {
				expPress = vPress
				candidate = k
				pathTaken = path[1:]
				fmt.Printf(" %s is the candidate\n", k)
			} else {
				fmt.Println()
			}
		}
		if start != candidate && candidate != "" {
			fmt.Printf("Move to %s\n\n", candidate)
			fullPath += strings.Join(pathTaken, " ") + " "
			released += expPress
			i += timeUsed
			delete(valves, candidate)
			start = candidate
			candidate = ""
		} else {
			fmt.Println("No viable moves")
			fmt.Println()
		}
	}

	return released
}

func (p Day16) PartB(lines []string) any {
	return "implement_me"
}
