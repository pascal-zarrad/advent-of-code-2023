package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "04"

func cleanUpAndConvertNumbers(numbers []string) []int {
	cleanNumbers := make([]int, 0)

	for _, number := range numbers {
		number := strings.Trim(number, " ")

		if number == "" {
			continue
		}

		if iNumber, err := strconv.Atoi(number); err == nil {
			cleanNumbers = append(cleanNumbers, iNumber)
		}
	}

	return cleanNumbers
}

func solvePartOne(input string) string {
	cards := strings.Split(input, "\n")

	points := 0

	for _, card := range cards {
		gameIdCardContentTuple := strings.Split(strings.Trim(card, "\n"), ":")
		card := strings.Split(gameIdCardContentTuple[1], "|")

		winningNumbers := cleanUpAndConvertNumbers(strings.Split(card[0], " "))
		gameNumbers := cleanUpAndConvertNumbers(strings.Split(card[1], " "))

		currentCardPoints := 0

		for _, winningNumber := range winningNumbers {
			for _, gameNumber := range gameNumbers {
				if winningNumber == gameNumber {
					if currentCardPoints == 0 {
						currentCardPoints = 1
						break
					}

					currentCardPoints *= 2
					break
				}
			}
		}

		points += currentCardPoints
	}

	return fmt.Sprint(points)
}

type Scratchcard struct {
	winningsNumbers []int
	gameNumbers     []int
	count           int
}

func solvePartTwo(input string) string {
	cards := strings.Split(input, "\n")

	scratchcards := make([]*Scratchcard, 0)

	for _, card := range cards {
		gameIdCardContentTuple := strings.Split(strings.Trim(card, "\n"), ":")
		card := strings.Split(gameIdCardContentTuple[1], "|")

		winningNumbers := cleanUpAndConvertNumbers(strings.Split(card[0], " "))
		gameNumbers := cleanUpAndConvertNumbers(strings.Split(card[1], " "))

		scratchcards = append(scratchcards, &Scratchcard{
			winningsNumbers: winningNumbers,
			gameNumbers:     gameNumbers,
			count:           1,
		})
	}

	wonCardCount := 0

	for index, card := range scratchcards {
		cardId := index + 1
		wonCardCount += card.count

		currentWinCount := 0
		for _, winningNumber := range card.winningsNumbers {
			for _, gameNumber := range card.gameNumbers {
				if winningNumber == gameNumber {
					currentWinCount++
				}
			}
		}

		for i := 0; i < currentWinCount; i++ {
			if cardId+i >= len(scratchcards) {
				break
			}

			scratchcards[cardId+i].count += card.count

		}
	}

	return fmt.Sprint(wonCardCount)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
