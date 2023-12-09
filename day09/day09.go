package main

import (
	"fmt"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "09"

////////// COMMON

func prepareReadings(lines []string) [][]int {
	readingList := make([][]int, len(lines))

	for index, line := range lines {
		spittedLine := strings.Split(line, " ")

		readingList[index] = util.CleanUpAndConvertIntArray(spittedLine)
	}

	return readingList
}

func generateNextRow(sequence []int) ([]int, bool) {
	nextRow := make([]int, 0)
	sequenceLength := len(sequence)

	allZero := true
	for index, num := range sequence {
		if index+1 > sequenceLength-1 {
			break
		}

		diff := sequence[index+1] - num

		nextRow = append(nextRow, diff)

		if diff != 0 {
			allZero = false
		}
	}

	return nextRow, allZero
}

func generateSequences(readings []int) [][]int {
	sequences := make([][]int, 1)
	sequences[0] = readings

	currentRowPointer := 0

	for {
		currentSequence := sequences[currentRowPointer]

		row, allZero := generateNextRow(currentSequence)

		sequences = append(sequences, row)
		currentRowPointer++

		if allZero {
			break
		}
	}

	return sequences
}

////////// PART 1

func solvePartOne(input string) string {
	lines := strings.Split(input, "\n")

	readingList := prepareReadings(lines)

	sum := 0

	extrapolatedNumbers := computeExtrapolatedNextReading(readingList)
	for _, num := range extrapolatedNumbers {
		sum += num
	}

	return fmt.Sprint(sum)
}

func computeExtrapolatedNextReading(readingList [][]int) []int {
	extrapolatedNumbers := make([]int, len(readingList))

	for _, reading := range readingList {
		sequences := generateSequences(reading)
		nextNumber := calculateNextNumber(sequences)

		extrapolatedNumbers = append(extrapolatedNumbers, nextNumber)
	}

	return extrapolatedNumbers
}

func calculateNextNumber(sequences [][]int) int {
	// Last row is only 0's anyway, I just
	// handle it like a 0 has been added :)
	for i := len(sequences) - 2; i >= 0; i-- {
		currentSequence := sequences[i]
		pastSequence := sequences[i+1]

		currentLastValue := currentSequence[len(currentSequence)-1]
		lastLastValue := pastSequence[len(pastSequence)-1]
		newValue := currentLastValue + lastLastValue

		sequences[i] = append(currentSequence, newValue)
	}

	return sequences[0][len(sequences[0])-1]
}

////////// PART 2

func solvePartTwo(input string) string {
	lines := strings.Split(input, "\n")

	readingList := prepareReadings(lines)

	sum := 0

	extrapolatedNumbers := computeExtrapolatedPreviousReading(readingList)
	for _, num := range extrapolatedNumbers {
		sum += num
	}

	return fmt.Sprint(sum)
}

func computeExtrapolatedPreviousReading(readingList [][]int) []int {
	extrapolatedNumbers := make([]int, len(readingList))

	for _, reading := range readingList {
		sequences := generateSequences(reading)
		nextNumber := calculatePreviousNumber(sequences)

		extrapolatedNumbers = append(extrapolatedNumbers, nextNumber)
	}

	return extrapolatedNumbers
}

func calculatePreviousNumber(sequences [][]int) int {
	// Last row is only 0's anyway, I just
	// handle it like a 0 has been added :)
	for i := len(sequences) - 2; i >= 0; i-- {
		currentSequence := sequences[i]
		pastSequence := sequences[i+1]

		currentLastValue := currentSequence[0]
		lastLastValue := pastSequence[0]
		newValue := currentLastValue - lastLastValue

		newSequence := make([]int, len(currentSequence)+1)
		newSequence[0] = newValue

		sequences[i] = append(newSequence, currentSequence...)
	}

	return sequences[0][0]
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
