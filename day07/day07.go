package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "07"

////////// COMMON

type HandType int

const (
	FIVE_OF_A_KIND  HandType = 7
	FOUR_OF_A_KIND  HandType = 6
	FULL_HOUSE      HandType = 5
	THREE_OF_A_KIND HandType = 4
	TWO_PAIR        HandType = 3
	ONE_PAIR        HandType = 2
	HIGH_CARD       HandType = 1
)

var cardWeightMappingPart1 = map[rune]rune{
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
	'T': 'A',
	'J': 'B',
	'Q': 'C',
	'K': 'D',
	'A': 'E',
}

type CamelCardsHand struct {
	cards       string
	weightCards string
	bid         int
	handType    HandType
}

func convertRoundsToHands(rounds []string) []*CamelCardsHand {
	camelCardsHands := make([]*CamelCardsHand, 0)

	for _, round := range rounds {
		handAndBid := strings.Split(round, " ")

		camelCardsHands = append(camelCardsHands, &CamelCardsHand{
			cards: handAndBid[0],
			bid:   util.StringToInt(handAndBid[1]),
		})
	}

	return camelCardsHands
}

func calculateCardCounts(hand string) map[rune]int {
	cardCounts := map[rune]int{}

	for _, char := range hand {
		if _, ok := cardCounts[char]; !ok {
			cardCounts[char] = 1

			continue
		}

		cardCounts[char]++
	}

	return cardCounts
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}

////////// PART 1

func solvePartOne(input string) string {
	rounds := strings.Split(input, "\n")

	camelCardsHand := convertRoundsToHands(rounds)
	sum := 0

	// Two-step process is easier than doing a multi-condition sort
	// Additionally, no need
	rankedHands := rankHandsPart1(camelCardsHand)

	for rank, hand := range rankedHands {
		sum += (rank + 1) * hand.bid
	}

	return fmt.Sprint(sum)
}

func findHandTypePart1(camelCardsHand *CamelCardsHand) HandType {
	cardCounts := calculateCardCounts(camelCardsHand.cards)

	switch len(cardCounts) {
	case 5:
		return HIGH_CARD
	case 4:
		return ONE_PAIR
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return THREE_OF_A_KIND
			}
		}

		return TWO_PAIR
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return FOUR_OF_A_KIND
			}
		}

		return FULL_HOUSE
	case 1:
		return FIVE_OF_A_KIND
	}

	panic("Got invalid hand!")
}

func prepareHandPart1(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	for _, hand := range camelCardsHands {
		hand.handType = findHandTypePart1(hand)

		weightCards := []rune(hand.cards)
		for i := 0; i < len(weightCards); i++ {
			weightCards[i] = cardWeightMappingPart1[weightCards[i]]
		}
		hand.weightCards = string(weightCards)
	}

	return camelCardsHands
}

func rankHandsPart1(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	camelCardsHands = prepareHandPart1(camelCardsHands)

	sort.Slice(camelCardsHands, func(i, j int) bool {
		if camelCardsHands[i].handType < camelCardsHands[j].handType {
			return true
		}

		if camelCardsHands[i].handType > camelCardsHands[j].handType {
			return false
		}
		return camelCardsHands[i].weightCards < camelCardsHands[j].weightCards
	})

	return camelCardsHands
}

////////// PART 2

var cardWeightMappingPart2 = map[rune]rune{
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
	'T': 'A',
	'J': '1',
	'Q': 'C',
	'K': 'D',
	'A': 'E',
}

func solvePartTwo(input string) string {
	rounds := strings.Split(input, "\n")

	camelCardsHand := convertRoundsToHands(rounds)
	sum := 0

	// Two-step process is easier than doing a multi-condition sort
	// Additionally, no need
	rankedHands := rankHandsPart2(camelCardsHand)

	for rank, hand := range rankedHands {
		sum += (rank + 1) * hand.bid
	}

	return fmt.Sprint(sum)
}

func adjustCardCountsPart2(cardCounts map[rune]int) map[rune]int {
	if count, ok := cardCounts['J']; ok {
		highestChar := 'J'
		highestCount := 0
		for char, charCount := range cardCounts {
			if char == 'J' {
				continue
			}

			if charCount <= highestCount {
				continue
			}

			highestChar = char
			highestCount = charCount
		}

		if highestChar != 'J' {
			cardCounts[highestChar] += count
			delete(cardCounts, 'J')
		}
	}

	return cardCounts
}

func findHandTypePart2(camelCardsHand *CamelCardsHand) HandType {
	cardCounts := calculateCardCounts(camelCardsHand.cards)
	cardCounts = adjustCardCountsPart2(cardCounts)

	switch len(cardCounts) {
	case 5:
		return HIGH_CARD
	case 4:
		return ONE_PAIR
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return THREE_OF_A_KIND
			}
		}

		return TWO_PAIR
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return FOUR_OF_A_KIND
			}
		}

		return FULL_HOUSE
	case 1:
		return FIVE_OF_A_KIND
	}

	panic("Got invalid hand!")
}

func prepareHandPart2(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	for _, hand := range camelCardsHands {
		hand.handType = findHandTypePart2(hand)

		weightCards := []rune(hand.cards)
		for i := 0; i < len(weightCards); i++ {
			weightCards[i] = cardWeightMappingPart2[weightCards[i]]
		}
		hand.weightCards = string(weightCards)
	}

	return camelCardsHands
}

func rankHandsPart2(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	camelCardsHands = prepareHandPart2(camelCardsHands)

	sort.Slice(camelCardsHands, func(i, j int) bool {
		if camelCardsHands[i].handType < camelCardsHands[j].handType {
			return true
		}

		if camelCardsHands[i].handType > camelCardsHands[j].handType {
			return false
		}
		return camelCardsHands[i].weightCards < camelCardsHands[j].weightCards
	})

	return camelCardsHands
}
