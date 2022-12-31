package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

type Day15 struct{}

const (
	// ARRRRRGHHH!!!! with the name collisions! i've got to sort this shit out.
	kind15Empty = iota
	kind15Sensor
	kind15Beacon
	kind15Coverage
)

var kind15Runes = map[int]rune{
	kind15Empty:    '.',
	kind15Sensor:   'S',
	kind15Beacon:   'B',
	kind15Coverage: '#',
}

func getX(c complex128) float64 {
	return real(c)
}

func getY(c complex128) float64 {
	return imag(c)
}

// getDist returns the Manhattan distance between two coords expressed as complex numbers
func getDist(a, b complex128) float64 {
	v1 := getX(a)
	v2 := getX(b)
	xDist := v1 - v2
	if xDist < 0 {
		xDist = xDist * -1
	}

	v1 = getY(a)
	v2 = getY(b)
	yDist := v1 - v2
	if yDist < 0 {
		yDist = yDist * -1
	}

	return yDist + xDist
}

func isSensed(cell complex128, sensors map[complex128]float64) bool {
	for senseCoord, senseRange := range sensors {
		if senseRange >= getDist(cell, senseCoord) {
			return true
		}
	}
	return false
}

type caveData struct {
	// a map representation of the contents of the cave
	//  complex coords, int from the set of kind consts
	contents map[complex128]int
	// map of sensors with complex coords, range as float
	sensors map[complex128]float64
	// map of beacons with complex coords, value is meaningless (map for easy boolean logic elsewhere)
	beacons map[complex128]int
	// cave extents
	xMin, yMin, xMax, yMax float64
}

func parseInput15(lines []string) caveData {
	cave := caveData{
		contents: map[complex128]int{},
		sensors:  map[complex128]float64{},
		beacons:  map[complex128]int{},
		xMin:     math.MaxFloat64,
		yMin:     math.MaxFloat64,
		xMax:     0.0,
		yMax:     0.0,
	}
	re := regexp.MustCompile("[-0-9]+")

	for _, line := range lines {
		inputs := re.FindAllString(line, -1)
		if len(inputs) != 4 {
			logrus.Fatalf("parse fail on: %s", line)
		}
		xS := common.MustFloat(inputs[0])
		yS := common.MustFloat(inputs[1])
		xB := common.MustFloat(inputs[2])
		yB := common.MustFloat(inputs[3])

		// add the sensor
		senseRange := getDist(complex(xS, yS), complex(xB, yB))
		cave.sensors[complex(xS, yS)] = senseRange
		cave.contents[complex(xS, yS)] = kind15Sensor
		// add the beacon
		cave.beacons[complex(xB, yB)] = 1.0
		cave.contents[complex(xB, yB)] = kind15Beacon

		// update extents
		// xMin
		if xS-senseRange < cave.xMin {
			cave.xMin = xS - senseRange
		}
		if xB < cave.xMin {
			// this feels redundant. all sensors have range to touch beaons?
			cave.xMin = xB
		}
		// xMax
		if xS+senseRange > cave.xMax {
			cave.xMax = xS + senseRange
		}
		if xB > cave.xMax {
			cave.xMax = xB
		}
		// yMin
		if yS-senseRange < cave.yMin {
			cave.yMin = yS - senseRange
		}
		if yB < cave.yMin {
			cave.yMin = yB
		}
		// yMax
		if yS+senseRange > cave.yMax {
			cave.yMax = yS + senseRange
		}
		if yB > cave.yMax {
			cave.yMax = yB
		}
	}

	// xMin and yMax are NOT coming out of this right?!?
	return cave
}

func computeCoverage(cave caveData) {
	for y := cave.yMin; y < cave.yMax; y++ {
		for x := cave.xMin; x < cave.xMax; x++ {
			if cave.contents[complex(x, y)] == kind15Empty && isSensed(complex(x, y), cave.sensors) {
				cave.contents[complex(x, y)] = kind15Coverage
			}
		}
	}
}

func printCave15(cave caveData) []string {
	rows := make([]string, int(cave.yMax+1))
	row := fmt.Sprintf("%5d", int(cave.xMin))
	rows = append(rows, row)
	for y := cave.yMin; y <= cave.yMax; y++ {
		row = fmt.Sprintf("%3d", int(y))
		for x := cave.xMin; x <= cave.xMax; x++ {
			r := kind15Runes[cave.contents[complex(x, y)]]
			row += string(r)
		}
		rows = append(rows, row)
	}

	return rows
}

func (p Day15) PartA(lines []string) any {
	cave := parseInput15(lines)
	// fmt.Printf("Cave is\n%f\t%f -> %f\nV\n%f\n", cave.yMin, cave.xMin, cave.xMax, cave.yMax)

	// for sample data
	y := 10.0
	// part1
	y = 2000000.0

	count := 0
	for x := cave.xMin; x < cave.xMax; x++ {
		if (cave.contents[complex(x, y)] == kind15Empty || cave.contents[complex(x, y)] == kind15Coverage) &&
			isSensed(complex(x, y), cave.sensors) {
			count++
		}
	}
	// fmt.Println(strings.Join(printCave15(cave), "\n"))
	return count
}

func (p Day15) PartB(lines []string) any {
	/* if each row takes 2.5 seconds, which the timing below suggests on this gitpod host
	then this algorithm will run for ~115 days.
	that ain't gonna work.
	*/
	cave := parseInput15(lines)
	// boundaries, min 0, sample values in init
	xBound := 20.0
	yBound := 20.0
	xBound = 4000000.0
	yBound = 4000000.0
	tuningFreq := 0.0

	for y := 0.0; y <= yBound && tuningFreq == 0.0; y++ {
		fmt.Printf("y=%f", y)
		start := time.Now()
		for x := 0.0; x <= xBound && tuningFreq == 0.0; x++ {
			if !isSensed(complex(x, y), cave.sensors) {
				tuningFreq = (x * 4000000.0) + y
			}
		}
		fmt.Printf(" took %s\n", time.Since(start))
	}

	return tuningFreq
}
