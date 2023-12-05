package main

import (
	"fmt"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "05"

func applyAlmanacMapping(rawMapping string, currentStepsIds []int) []int {
	newStepIds := make([]int, len(currentStepsIds))
	copy(newStepIds, currentStepsIds)

	mappingLines := strings.Split(rawMapping, "\n")[1:]

	for _, instruction := range mappingLines {
		columns := strings.Split(instruction, " ")

		destinationStart := util.StringToInt(columns[0])
		sourceStart := util.StringToInt(columns[1])
		mappingRange := util.StringToInt(columns[2])
		upperBound := sourceStart + (mappingRange - 1)

		for index := 0; len(currentStepsIds) > index; index++ {
			currentId := currentStepsIds[index]
			if currentId < sourceStart || currentId > upperBound {
				continue
			}

			destinationId := destinationStart + (currentId - sourceStart)

			newStepIds[index] = destinationId
		}
	}

	return newStepIds
}

func solvePartOne(input string) string {
	almanacMaps := strings.Split(input, "\n\n")

	// First one is seeds -> start value
	seeds := strings.Split(almanacMaps[0], " ")[1:]

	currentStepIds := make([]int, 0)
	for _, seed := range seeds {
		currentStepIds = append(currentStepIds, util.StringToInt(seed))
	}

	for _, rawMapping := range almanacMaps[1:] {
		currentStepIds = applyAlmanacMapping(rawMapping, currentStepIds)
	}

	lowestNumber := currentStepIds[0]
	for _, id := range currentStepIds[1:] {
		if id < lowestNumber {
			lowestNumber = id
		}
	}

	return fmt.Sprint(lowestNumber)
}

func convertSeedsToInts(seeds []string) []int {
	numberSeeds := make([]int, len(seeds))

	for index, seed := range seeds {
		numberSeeds[index] = util.StringToInt(seed)
	}

	return numberSeeds
}

// I didn't have time to implement proper range mapping
// and decided to just bruteforce this day.
// Some proper range mapping could solve this in a few ms.
//
// The following code took between 4GB and 23GB! RAM.
// Thanks to processing each range on its own, this is an ok amount
// of resource consumption.
//
// Without processing each range on its own, resource consumption can take
// up to >~70GB RAM. It barely ran through on an 64GB server by using some SWAP.
// The current version should run fine on any 32GB system with 25GB free (or some SWAP).
//
// # Time on an i7-7700 server took around 3m38.994s
//
// Still way faster than doing the proper implementation.
func solvePartTwo(input string) string {
	almanacMaps := strings.Split(input, "\n\n")

	// First one is seeds -> start value
	seeds := strings.Split(almanacMaps[0], " ")[1:]
	numberSeeds := convertSeedsToInts(seeds)

	lowestNumber := numberSeeds[0]

	seedStart := -1
	for _, seed := range numberSeeds {
		if seedStart == -1 {
			seedStart = seed

			continue
		}

		currentStepIds := make([]int, 0)
		for i := 0; i < seed; i++ {
			currentStepIds = append(currentStepIds, seedStart+i)
		}

		for _, rawMapping := range almanacMaps[1:] {
			currentStepIds = applyAlmanacMapping(rawMapping, currentStepIds)
		}

		for index, id := range currentStepIds[1:] {
			if index%2 != 0 {
				continue
			}

			if id < lowestNumber {
				lowestNumber = id
			}
		}

		seedStart = -1
	}

	return fmt.Sprint(lowestNumber)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
