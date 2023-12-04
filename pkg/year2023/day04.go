package year2023

import (
	"aocgen/pkg/common"
	"math"
	"slices"
	"strings"
)

type Day04 struct{}

type scratchCard struct {
	winners, numbers []int
	id, winCount     int
}

func (p Day04) PartA(lines []string) any {
	cards := []scratchCard{}
	for _, input := range lines {
		cards = append(cards, parseCard(input))
	}
	retScore := 0
	for _, card := range cards {
		if card.winCount > 0 {
			retScore += int(math.Pow(2, float64(card.winCount-1)))
		}
	}
	return retScore
}

func (p Day04) PartB(lines []string) any {
	plays := map[int]scratchCard{}
	allCards := []scratchCard{}

	for _, input := range lines {
		play := parseCard(input)
		plays[play.id] = play
		allCards = append(allCards, play)
	}

	for i := 0; i < len(allCards); i++ {
		if allCards[i].winCount > 0 {
			for j := 1; j <= allCards[i].winCount; j++ {
				allCards = append(allCards, plays[allCards[i].id+j])
			}
		}
	}

	return len(allCards)
}

func parseCard(input string) (card scratchCard) {
	card = scratchCard{winners: []int{}, numbers: []int{}, winCount: 0, id: 0}
	split1 := strings.Split(input, ":")
	cardHeader := strings.Split(split1[0], " ")
	card.id = common.Atoi(cardHeader[len(cardHeader)-1])
	nums := strings.Split(split1[1], "|")
	for _, winner := range strings.Split(nums[0], " ") {
		if winner != "" {
			card.winners = append(card.winners, common.Atoi(winner))
		}
	}

	for _, playString := range strings.Split(nums[1], " ") {
		if playString != "" {
			playNum := common.Atoi(playString)
			card.numbers = append(card.winners, playNum)
			if slices.Contains(card.winners, playNum) {
				card.winCount++
			}
		}
	}

	return card
}
