package main

import (
	"strconv"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "01"

var mappingTable = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func solvePartOne(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		runes := []rune(line)

		digitA := ""
		digitB := ""
		for i := 0; i < len(runes); i++ {
			if _, err := strconv.Atoi(string(runes[i])); err != nil {
				continue
			}

			if digitA == "" {
				digitA = string(runes[i])
				digitB = string(runes[i])

				continue
			}

			digitB = string(runes[i])
		}

		// Safeguard against empty strings to int conversion
		if digitA == "" || digitB == "" {
			continue
		}

		sum += util.StringToInt(digitA + digitB)
	}

	return sum
}

func solvePartTwo(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		runes := []rune(line)

		digitA := ""
		digitB := ""
		for i := 0; i < len(runes); i++ {
			currentCharacter := getDigitIfStringStartingWithNumberOrNumberWord(string(runes[i:]))

			if currentCharacter == "" {
				continue
			}

			if digitA == "" {
				digitA = string(currentCharacter)
				digitB = string(currentCharacter)

				continue
			}

			digitB = string(currentCharacter)
		}

		// Safeguard against empty strings to int conversion
		if digitA == "" || digitB == "" {
			continue
		}

		sum += util.StringToInt(digitA + digitB)
	}

	return sum
}

func getDigitIfStringStartingWithNumberOrNumberWord(str string) string {
	if len(str) == 0 {
		return ""
	}

	if _, err := strconv.Atoi(string(str[0])); err == nil {
		return string(str[0])
	}

	for key, value := range mappingTable {
		if strings.HasPrefix(str, key) {
			return value
		}
	}

	return ""
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
