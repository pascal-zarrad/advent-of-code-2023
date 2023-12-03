package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "03"

// Todays solutions are not that great, there are better way :D

func solvePartOne(input string) string {
	lines := strings.Split(input, "\n")

	schematic := make([][]rune, len(lines))

	for lineIndex, line := range lines {
		schematic[lineIndex] = []rune(line)
	}

	partNumberSum := 0

	for rowIndex, schemaRow := range schematic {
		numberStartPointer := -1
		numberEndPointer := -1

		currentNumberString := ""

		for characterIndex, character := range schemaRow {
			if _, err := strconv.Atoi(string(character)); nil == err {
				currentNumberString += string(character)

				if numberStartPointer == -1 {
					numberStartPointer = characterIndex
				}

				numberEndPointer = characterIndex

				// Numeric character at end of line must be handled as non-int
				// and loop must proceed
				if characterIndex != len(schemaRow)-1 {
					continue
				}
			}

			if numberStartPointer == -1 || numberEndPointer == -1 {
				continue
			}

			// With start and end pointer defined, lets scan for symbols!
			if scanForSymbols(schematic, rowIndex, numberStartPointer-1, numberEndPointer+1) {
				currentNumber, err := strconv.Atoi(currentNumberString)
				if nil != err {
					panic(fmt.Sprintf("\"%s\" is not a valid integer", currentNumberString))
				}

				partNumberSum += currentNumber
			}

			// Reset, multiple parts per row are possible
			numberStartPointer = -1
			numberEndPointer = -1
			currentNumberString = ""
		}

	}

	return fmt.Sprint(partNumberSum)
}

func scanForSymbols(schematic [][]rune, rowIndex int, scanStartPointer int, scanEndPointer int) bool {
	if scanRowForSymbols(schematic, rowIndex-1, scanStartPointer, scanEndPointer) {
		return true
	}

	if scanRowForSymbols(schematic, rowIndex, scanStartPointer, scanEndPointer) {
		return true
	}
	if scanRowForSymbols(schematic, rowIndex+1, scanStartPointer, scanEndPointer) {
		return true
	}

	return false
}

func scanRowForSymbols(schematic [][]rune, rowIndex int, scanStartPointer int, scanEndPointer int) bool {
	if rowIndex < 0 || rowIndex > len(schematic)-1 {
		return false
	}

	row := schematic[rowIndex]

	if scanStartPointer < 0 {
		scanStartPointer = 0
	}

	if scanEndPointer > len(row)-1 {
		scanEndPointer = len(row) - 1
	}

	for i := scanStartPointer; i <= scanEndPointer; i++ {
		if row[i] == '.' {
			continue
		}

		if _, err := strconv.Atoi(string(row[i])); nil == err {
			continue
		}

		return true
	}

	return false
}

type Gear struct {
	x int
	y int
}

func solvePartTwo(input string) string {
	lines := strings.Split(input, "\n")

	schematic := make([][]rune, len(lines))

	for lineIndex, line := range lines {
		schematic[lineIndex] = []rune(line)
	}

	gearRatioProduct := 0
	gearMap := make([]Gear, 0)

	for rowIndex, schemaRow := range schematic {
		for characterIndex, character := range schemaRow {
			if character != '*' {
				continue
			}

			gearMap = append(gearMap, Gear{x: characterIndex, y: rowIndex})
		}
	}

	for _, gear := range gearMap {
		// calculateGearRatio should return 0 if not a valid gear
		gearRatioProduct += calculateGearRatio(schematic, gear)
	}

	return fmt.Sprint(gearRatioProduct)
}

func calculateGearRatio(schematic [][]rune, gear Gear) int {
	ratios := append(make([]int, 0), findNumbersForRow(schematic, gear.x, gear.y-1)...)
	ratios = append(ratios, findNumbersForRow(schematic, gear.x, gear.y)...)
	ratios = append(ratios, findNumbersForRow(schematic, gear.x, gear.y+1)...)

	if len(ratios) != 2 {
		return 0
	}

	return ratios[0] * ratios[1]
}

func findNumbersForRow(schematic [][]rune, x int, y int) []int {
	if y < 0 || y > len(schematic)-1 {
		return make([]int, 0)
	}

	gearRatios := make([]int, 0)

	realX := tryFindRealFirstDigitX(schematic, x, y)

	numberStartPointer := -1
	numberEndPointer := -1

	currentNumberString := ""
	for i := realX; i < len(schematic[y]); i++ {
		if numberStartPointer == -1 && numberEndPointer == -1 && i > x+1 {
			break
		}

		character := schematic[y][i]

		if _, err := strconv.Atoi(string(character)); nil == err {
			currentNumberString += string(character)

			if numberStartPointer == -1 {
				numberStartPointer = i
			}

			numberEndPointer = i

			// Numeric character at end of line must be handled as non-int
			// and loop must proceed
			if i != len(schematic[y])-1 {
				continue
			}
		}

		if numberStartPointer == -1 || numberEndPointer == -1 {
			continue
		}

		// With start and end pointer defined, lets scan for symbols!
		currentNumber, err := strconv.Atoi(currentNumberString)
		if nil != err {
			panic(fmt.Sprintf("\"%s\" is not a valid integer", currentNumberString))
		}

		gearRatios = append(gearRatios, currentNumber)

		// Reset, multiple parts per row are possible
		numberStartPointer = -1
		numberEndPointer = -1
		currentNumberString = ""
	}

	return gearRatios
}

func tryFindRealFirstDigitX(schematic [][]rune, x int, y int) int {
	if x-1 < 0 {
		return x
	}

	if _, err := strconv.Atoi(string(schematic[y][x-1])); nil != err {
		return x
	}

	for i := x - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(schematic[y][i])); nil != err {
			return i + 1
		}
	}

	// Lower boundary
	return 0
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
