package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "02"

func solvePartOne(input string) string {
	lines := strings.Split(input, "\n")

	cubeBudget := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGamesSum := 0

	for _, line := range lines {
		gameIdAndContent := strings.Split(line, ":")
		if len(gameIdAndContent) != 2 {
			panic(fmt.Sprintf("A valid game must have format: Game 00: ...\", got: \"%s\"", line))
		}

		id, err := strconv.Atoi(strings.Split(gameIdAndContent[0], " ")[1])
		if nil != err {
			panic(fmt.Sprintf("Game id is not an int for game: \"%s\"", line))
		}

		matchValid := true

		for _, round := range strings.Split(gameIdAndContent[1], ";") {
			for _, cubes := range strings.Split(round, ",") {
				countAndColor := strings.Split(strings.Trim(cubes, " "), " ")
				if len(countAndColor) != 2 {
					panic(fmt.Sprintf("The format \"%s\" is invalid for cubes and count", cubes))
				}

				color := countAndColor[1]
				cubeCount, err := strconv.Atoi(countAndColor[0])
				if nil != err {
					panic(fmt.Sprintf("The count for \"%s\" is not a valid integer", cubes))
				}

				maxCount, ok := cubeBudget[color]
				if !ok {
					panic(fmt.Sprintf("There is not budget for color \"%s\"", color))
				}

				if cubeCount > maxCount {
					matchValid = false
					break

				}
			}

			if !matchValid {
				break
			}

		}

		if matchValid {
			possibleGamesSum += id
		}
	}

	return fmt.Sprint(possibleGamesSum)
}

func solvePartTwo(input string) string {
	lines := strings.Split(input, "\n")

	cubePower := 0

	for _, line := range lines {
		gameIdAndContent := strings.Split(line, ":")
		if len(gameIdAndContent) != 2 {
			panic(fmt.Sprintf("A valid game must have format: Game 00: ...\", got: \"%s\"", line))
		}

		highestCountsByColor := map[string]int{}

		for _, round := range strings.Split(gameIdAndContent[1], ";") {
			for _, cubes := range strings.Split(round, ",") {
				countAndColor := strings.Split(strings.Trim(cubes, " "), " ")
				if len(countAndColor) != 2 {
					panic(fmt.Sprintf("The format \"%s\" is invalid for cubes and count", cubes))
				}

				color := countAndColor[1]
				cubeCount, err := strconv.Atoi(countAndColor[0])
				if nil != err {
					panic(fmt.Sprintf("The count for \"%s\" is not a valid integer", cubes))
				}

				highestCount, ok := highestCountsByColor[color]
				if !ok {
					highestCountsByColor[color] = cubeCount
				}

				if highestCount < cubeCount {
					highestCountsByColor[color] = cubeCount
				}
			}
		}

		var gamePower int
		for _, count := range highestCountsByColor {
			if 0 == gamePower {
				gamePower = count
				continue
			}

			gamePower *= count
		}

		cubePower += gamePower
	}

	return fmt.Sprint(cubePower)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
