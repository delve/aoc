package year2022

import (
	"aocgen/pkg/common"
	"math"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sirupsen/logrus"
)

type Day09 struct{}

type rope struct {
	knots     [][2]int           // xy coords
	tailTrail mapset.Set[[2]int] // set of xy coords
}

func (r *rope) move(d string, c int) {
	motion := func() {}
	switch d {
	case "U":
		motion = r.moveU
	case "D":
		motion = r.moveD
	case "L":
		motion = r.moveL
	case "R":
		motion = r.moveR
	default:
		logrus.Fatalf("Unknown dir: %s", d)
	}
	for i := 0; i < c; i++ {
		motion()
	}
}

func (r *rope) moveU() {
	r.knots[0][1] += 1
	r.catchupRope()
}
func (r *rope) moveD() {
	r.knots[0][1] += -1
	r.catchupRope()
}
func (r *rope) moveL() {
	r.knots[0][0] += -1
	r.catchupRope()
}
func (r *rope) moveR() {
	r.knots[0][0] += 1
	r.catchupRope()
}

func (r *rope) catchupRope() {
	for i, knot := range r.knots {
		if i > 0 { // skip the first knot aka head
			r.knots[i] = r.catchupKnot(r.knots[i-1], knot)
		}
		r.markTailTrail()
	}
}

func (r *rope) markTailTrail() {
	// tail is the very last knot
	tail := r.knots[len(r.knots)-1]
	r.tailTrail.Add(tail)
}

func (r *rope) catchupKnot(head, tail [2]int) [2]int {
	xOffset := int(math.Abs(float64(head[0] - tail[0])))
	yOffset := int(math.Abs(float64(head[1] - tail[1])))

	if xOffset > 1 || yOffset > 1 {
		if xOffset != 0 && yOffset != 0 {
			// diagonal motion
			if head[0] > tail[0] {
				tail[0] += 1
			} else {
				tail[0] += -1
			}
			if head[1] > tail[1] {
				tail[1] += 1
			} else {
				tail[1] += -1
			}
		} else {
			// straight motion
			if xOffset > 1 {
				if head[0] > tail[0] {
					tail[0] += xOffset - 1
				} else {
					tail[0] += (-1 * xOffset) + 1
				}
			} else {
				if head[1] > tail[1] {
					tail[1] += yOffset - 1
				} else {
					tail[1] += (-1 * yOffset) + 1
				}
			}
		}
	} // x and y offsets are 0 or 1, there's no catching up to do

	return tail
}

func parseInput09(r rope, lines []string) rope {
	for _, line := range lines {
		inLine := strings.Split(line, " ")
		distance, err := strconv.Atoi(inLine[1])
		common.Check(err)
		r.move(inLine[0], distance)
	}
	return r
}

func (p Day09) PartA(lines []string) any {
	positionsTailVisited := 0
	knotCount := 2
	rope := rope{tailTrail: mapset.NewSet[[2]int]()}
	for i := 0; i < knotCount; i++ {
		rope.knots = append(rope.knots, [2]int{0, 0})
	}
	rope.markTailTrail()

	parseInput09(rope, lines)
	positionsTailVisited = rope.tailTrail.Cardinality()
	return positionsTailVisited
}

func (p Day09) PartB(lines []string) any {
	positionsTailVisited := 0
	// OHNO 10 knots!
	knotCount := 10
	rope := rope{tailTrail: mapset.NewSet[[2]int]()}
	for i := 0; i < knotCount; i++ {
		rope.knots = append(rope.knots, [2]int{0, 0})
	}
	rope.markTailTrail()

	parseInput09(rope, lines)
	positionsTailVisited = rope.tailTrail.Cardinality()
	return positionsTailVisited
}
