package year2022

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Day02 struct{}

type scoreAccumulator struct {
	score, wins, losses, draws, rock, paper, scissors int
	decryptedScore                                    int
}

func parseShape(s string) int {
	ret := 0
	switch s {
	case "A":
		ret = 1
	case "X":
		ret = 1
	case "B":
		ret = 2
	case "Y":
		ret = 2
	case "C":
		ret = 3
	case "Z":
		ret = 3
	default:
		logrus.Fatalf("Unknown shape %s", s)
	}
	return ret
}

// returns 0: s1 loses, 0: draw, 2: s1 wins
func parseGame(s1, s2 int) int {
	// 1<2<3<1
	if s1 == s2 {
		return 1
	}
	if s1 == 3 && s2 == 1 {
		return 0
	}
	if s1 == 1 && s2 == 3 {
		return 2
	}
	if s1 > s2 {
		return 2
	}
	if s1 < s2 {
		return 0
	}

	logrus.Fatalf("Can't tell who won: %d : %d", s1, s2)
	return 4 // we should never reach here anyway
}

func getWin(s int) int {
	switch s {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 1
	}
	logrus.Fatalf("unknown input %d", s)
	return 4 // we should never reach here anyway
}

func getDraw(s int) int {
	return s
}

func getLoss(s int) int {
	switch s {
	case 1:
		return 3
	case 2:
		return 1
	case 3:
		return 2
	}
	logrus.Fatalf("unknown input %d", s)
	return 4 // we should never reach here anyway
}

func parseInput(lines []string) scoreAccumulator {
	accumulator := scoreAccumulator{
		score:          0,
		wins:           0,
		losses:         0,
		draws:          0,
		rock:           0,
		paper:          0,
		scissors:       0,
		decryptedScore: 0,
	}

	for i, line := range lines {
		shapes := strings.Split(line, " ")

		if len(shapes) != 2 {
			logrus.Fatalf("couldn't find player's shapes in %s on line # %d", line, i)
		}
		shape1 := parseShape(strings.TrimSpace(shapes[0]))
		shape2 := parseShape(strings.TrimSpace(shapes[1]))
		accumulator.score += shape2

		switch shape2 {
		case 1:
			accumulator.rock++
		case 2:
			accumulator.paper++
		case 3:
			accumulator.scissors++
		}

		switch parseGame(shape2, shape1) {
		case 0:
			// player loses
			accumulator.losses++
		case 1:
			// draw
			accumulator.score += 3
			accumulator.draws++
		case 2:
			// player wins
			accumulator.wins++
			accumulator.score += 6
		}

		// calculate for part 2
		switch shapes[1] {
		case "X":
			accumulator.decryptedScore += 0 + getLoss(shape1)
		case "Y":
			accumulator.decryptedScore += 3 + getDraw(shape1)
		case "Z":
			accumulator.decryptedScore += 6 + getWin(shape1)
		}
	}
	return accumulator
}

func (p Day02) PartA(lines []string) any {
	accumulator := parseInput(lines)

	return accumulator.score
}

func (p Day02) PartB(lines []string) any {
	accumulator := parseInput(lines)

	return accumulator.decryptedScore
}
