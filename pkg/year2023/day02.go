package year2023

import (
	"aocgen/pkg/common"
	"fmt"
	"regexp"
	"strings"
)

type Day02 struct{}

type game struct {
	id, maxB, maxR, maxG int
	samples              []sample
}

type sample struct {
	red, blue, green int
}

var rgxRed = regexp.MustCompile(`(?P<Count>\d+) red`)
var rgxBlue = regexp.MustCompile(`(?P<Count>\d+) blue`)
var rgxGreen = regexp.MustCompile(`(?P<Count>\d+) green`)

func (p Day02) PartA(lines []string) any {
	maxCubes := sample{red: 12, green: 13, blue: 14}

	games := getGames(lines)

	gameSum:=0
	for _, game := range games {
		if game.couldBePlayed(maxCubes) {
			gameSum += game.id
		}
	}
	return gameSum
}

func (p Day02) PartB(lines []string) any {
	games := getGames(lines)

	powerSum:=0
	for _, game := range games {
		powerSum += game.power()
	}
	return powerSum
}

func (g game) couldBePlayed(max sample) bool {
	if g.maxB > max.blue {
		return false
	}
	if g.maxG > max.green {
		return false
	}
	if g.maxR > max.red {
		return false
	}

	return true
}

func (g game) power() int {
	return g.maxB*g.maxG*g.maxR
}

func getGames(lines []string) (retGames []game) {
	retGames = []game{}

	for _, line := range lines {
		retGames = append(retGames, parseGame(line))
	}

	return
}

func parseGame(line string) (retGame game) {
	retGame = game{id: 0, maxB: 0, maxR: 0, maxG: 0, samples: []sample{}}

	split := strings.Split(line, ":")
	fmt.Sscanf(split[0], "Game %d:", &retGame.id)
	split = strings.Split(split[1], ";")
	for _, smpl := range split {
		sample := sample{green: 0, blue: 0, red: 0}
		// fmt.Printf("%#v\n", smpl)
		// fmt.Printf("%#v\n", rgxRed.FindStringSubmatch(smpl))
		// fmt.Printf("%#v\n", rgxBlue.FindStringSubmatch(smpl))
		// fmt.Printf("%#v\n", rgxGreen.FindStringSubmatch(smpl))
		if match := rgxRed.FindStringSubmatch(smpl); match != nil {
			sample.red = common.Atoi(match[1])
			if sample.red > retGame.maxR{
				retGame.maxR = sample.red
			}
		}
		if match := rgxBlue.FindStringSubmatch(smpl); match != nil {
			sample.blue = common.Atoi(match[1])
			if sample.blue > retGame.maxB{
				retGame.maxB = sample.blue
			}
		}
		if match := rgxGreen.FindStringSubmatch(smpl); match != nil {
			sample.green = common.Atoi(match[1])
			if sample.green > retGame.maxG{
				retGame.maxG = sample.green
			}
		}
		retGame.samples = append(retGame.samples, sample)
	}
	return
}
