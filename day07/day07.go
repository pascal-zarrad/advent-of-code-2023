package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "07"

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

func findHandType(camelCardsHand *CamelCardsHand) HandType {
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

func rankHands(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	hands := map[HandType][]*CamelCardsHand{}

	for _, hand := range camelCardsHands {
		handType := findHandType(hand)

		weightCards := []rune(hand.cards)
		for i := 0; i < len(weightCards); i++ {
			weightCards[i] = cardWeightMappingPart1[weightCards[i]]
		}
		hand.weightCards = string(weightCards)

		handsByType, ok := hands[handType]
		if !ok {
			handsByType = make([]*CamelCardsHand, 0)
		}

		hands[handType] = append(handsByType, hand)
	}

	for _, typedHands := range hands {
		// Order by hand type
		sort.Slice(typedHands, func(i, j int) bool {
			if typedHands[i].handType == 0 {
				typedHands[i].handType = findHandType(typedHands[i])
			}

			if typedHands[j].handType == 0 {
				typedHands[j].handType = findHandType(typedHands[j])
			}

			if camelCardsHands[i].handType < camelCardsHands[j].handType {
				return true
			}

			return false
		})
	}

	for _, typedHands := range hands {
		// Order by hand type
		sort.Slice(typedHands, func(i, j int) bool {
			return typedHands[i].weightCards < typedHands[j].weightCards
		})
	}

	sortedHands := make([]*CamelCardsHand, 0)
	for i := 1; i <= 7; i++ {
		sortedHands = append(sortedHands, hands[HandType(i)]...)
	}

	return sortedHands
}

func solvePartOne(input string) string {
	rounds := strings.Split(input, "\n")

	camelCardsHand := convertRoundsToHands(rounds)
	sum := 0

	// Two-step process is easier than doing a multi-condition sort
	// Additionally, no need
	rankedHands := rankHands(camelCardsHand)

	for rank, hand := range rankedHands {
		sum += (rank + 1) * hand.bid
	}

	return fmt.Sprint(sum)
}

func calculateCardCountsPart2(hand string) map[rune]int {
	cardCounts := map[rune]int{}

	for _, char := range hand {
		if _, ok := cardCounts[char]; !ok {
			cardCounts[char] = 1

			continue
		}

		cardCounts[char]++
	}

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
	cardCounts := calculateCardCountsPart2(camelCardsHand.cards)

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

func rankHandsPart2(camelCardsHands []*CamelCardsHand) []*CamelCardsHand {
	hands := map[HandType][]*CamelCardsHand{}

	for _, hand := range camelCardsHands {
		handType := findHandTypePart2(hand)

		weightCards := []rune(hand.cards)
		for i := 0; i < len(weightCards); i++ {
			weightCards[i] = cardWeightMappingPart2[weightCards[i]]
		}
		hand.weightCards = string(weightCards)

		handsByType, ok := hands[handType]
		if !ok {
			handsByType = make([]*CamelCardsHand, 0)
		}

		hands[handType] = append(handsByType, hand)
	}

	for _, typedHands := range hands {
		// Order by hand type
		sort.Slice(typedHands, func(i, j int) bool {
			if typedHands[i].handType == 0 {
				typedHands[i].handType = findHandTypePart2(typedHands[i])
			}

			if typedHands[j].handType == 0 {
				typedHands[j].handType = findHandTypePart2(typedHands[j])
			}

			if camelCardsHands[i].handType < camelCardsHands[j].handType {
				return true
			}

			return false
		})
	}

	for _, typedHands := range hands {
		// Order by hand type
		sort.Slice(typedHands, func(i, j int) bool {
			return typedHands[i].weightCards < typedHands[j].weightCards
		})
	}

	sortedHands := make([]*CamelCardsHand, 0)
	for i := 1; i <= 7; i++ {
		sortedHands = append(sortedHands, hands[HandType(i)]...)
	}

	return sortedHands
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

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
