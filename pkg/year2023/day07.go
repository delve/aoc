package year2023

import (
	"aocgen/pkg/common"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Day07 struct{}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

const (
	joker = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

// runeCards map input runes to cards
var runeCards = map[rune]int{
	'2': two,
	'3': three,
	'4': four,
	'5': five,
	'6': six,
	'7': seven,
	'8': eight,
	'9': nine,
	'T': ten,
	'J': jack,
	'Q': queen,
	'K': king,
	'A': ace,
}

// runeCards map input runes to cards for part 2
var runeCardsP2 = map[rune]int{
	'J': joker,
	'2': two,
	'3': three,
	'4': four,
	'5': five,
	'6': six,
	'7': seven,
	'8': eight,
	'9': nine,
	'T': ten,
	'Q': queen,
	'K': king,
	'A': ace,
}

type hand struct {
	cards [5]int
	grade int // because 'type' is a syntax error here
	bet   int
}

func (p Day07) PartA(lines []string) any {
	hands := p.parseInput(lines, jack)
	// for _, v := range hands {
	// 	fmt.Printf("%v - %v\n", v, v.grade)
	// }
	slices.SortFunc(hands, cmpGrade)
	// fmt.Printf("%v\n\n", hands)
	sortedHands := []hand{}
	for i := highCard; i < fiveOfAKind+1; i++ {
		gradeSlice := getGradeSlice(hands, i)
		// fmt.Printf("%v\n", gradeSlice)
		slices.SortFunc(gradeSlice, cmpCards)
		sortedHands = append(sortedHands, gradeSlice...)
	}
	// fmt.Printf("\n\n\n")
	// for _, v := range sortedHands {
	// 	fmt.Printf("%v - %v\n", v, v.grade)
	// }

	winnings := 0
	for i := 0; i < len(sortedHands); i++ {
		handwin := (i + 1) * sortedHands[i].bet
		// fmt.Println(handwin)
		winnings += handwin
	}

	return winnings
}

func (p Day07) PartB(lines []string) any {
	hands := p.parseInput(lines, joker)
	// for _, v := range hands {
	// 	fmt.Printf("%v - %v\n", v, v.grade)
	// }
	slices.SortFunc(hands, cmpGrade)
	// fmt.Printf("%v\n\n", hands)
	sortedHands := []hand{}
	for i := highCard; i < fiveOfAKind+1; i++ {
		gradeSlice := getGradeSlice(hands, i)
		// fmt.Printf("%v\n", gradeSlice)
		slices.SortFunc(gradeSlice, cmpCards)
		sortedHands = append(sortedHands, gradeSlice...)
	}
	// fmt.Printf("\n\n\n")
	// for _, v := range sortedHands {
	// 	fmt.Printf("%v - %v\n", v, v.grade)
	// }

	winnings := 0
	for i := 0; i < len(sortedHands); i++ {
		handwin := (i + 1) * sortedHands[i].bet
		// fmt.Println(handwin)
		winnings += handwin
	}

	return winnings
}

func (d Day07) parseInput(lines []string, jType int) (hands []hand) {
	hands = []hand{}
	var cardSet map[rune]int
	if jType == jack {
		cardSet = runeCards
	} else {
		cardSet = runeCardsP2
	}

	for _, line := range lines {
		inHand := hand{}
		input := strings.Split(line, " ")
		inHand.bet = common.Atoi(input[1])

		for i, inCard := range input[0] {
			inHand.cards[i] = cardSet[inCard]
		}
		inHand.grade = getHandGrade(inHand.cards, jType)
		hands = append(hands, inHand)
	}

	return hands
}

func getHandGrade(inCards [5]int, jType int) int {
	scoreCard := map[int]int{}
	for _, card := range inCards {
		scoreCard[card]++
	}
	retGrade := 0
	// if it's jack scoring or there's no jokers:
	if jType == jack || !slices.Contains(maps.Keys(scoreCard), joker) {
		retGrade = getSimpleGrade(scoreCard)
	} else {
		// must be joker scoring with jokers.
		faces := maps.Keys(scoreCard)
		for i := 0; i < len(faces); i++ {
			if faces[i] != joker {
				trialCards := maps.Clone(scoreCard)
				trialCards[faces[i]] += scoreCard[joker]
				delete(trialCards, joker)
				trialGrade := getSimpleGrade(trialCards)
				if trialGrade > retGrade {
					retGrade = trialGrade
				}
			}
		}
		// fmt.Println("leaving joker loop, debug breakpoint.")
		// hacky catch for JJJJJ cases
		if retGrade == 0 {
			retGrade = 6
		}
	}

	return retGrade
}

func getSimpleGrade(scoreCard map[int]int) (retGrade int) {
	faceCount := len(maps.Keys(scoreCard))
	switch faceCount {
	case 1: // only one card face, must be 5oaKind
		retGrade = fiveOfAKind
	case 2: // could be 4oaKind or a full house
		retGrade = fullHouse
		for _, face := range maps.Keys(scoreCard) {
			if scoreCard[face] == 4 {
				retGrade = fourOfAKind
			}
		}
	case 3: // could be 3oaKind or two pair
		retGrade = twoPair
		for _, face := range maps.Keys(scoreCard) {
			if scoreCard[face] == 3 {
				retGrade = threeOfAKind
			}
		}
	case 4: // must be one pair
		retGrade = onePair
	case 5: // must be high card
		retGrade = highCard
	}
	return retGrade
}

func cmpGrade(a, b hand) int {
	return a.grade - b.grade
}

func getGradeSlice(hands []hand, grade int) []hand {
	sliceBounds := [2]int{}

	sliceBounds[0] = slices.IndexFunc(hands, func(h hand) bool {
		return h.grade == grade
	})
	if sliceBounds[0] < 0 { // no hands of this grade
		return []hand{}
	}

	sliceBounds[1] = slices.IndexFunc(hands[sliceBounds[0]:], func(h hand) bool {
		return h.grade != grade
	})
	if sliceBounds[1] < 0 { // grade runs to the end of the slice
		sliceBounds[1] = len(hands)
	} else { // indexfunc is working of a partial copy of the source, so add the extra count from the part we hid from it
		// do it after the slice bounds check to avoid clobbering -1 returns from index func
		sliceBounds[1] = sliceBounds[1] + sliceBounds[0]
	}

	return hands[sliceBounds[0]:sliceBounds[1]]
}

func cmpCards(a, b hand) int {
	for i := 0; i < 5; i++ {
		if a.cards[i]-b.cards[i] != 0 {
			return a.cards[i] - b.cards[i]
		}
	}
	return 0
}
