package main

import (
	"fmt"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "06"

func findSolutions(time int, distance int) int {
	possibleSolutions := 0

	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			possibleSolutions++
		}
	}

	return possibleSolutions
}

func solvePartOne(input string) string {
	lines := strings.Split(input, "\n")

	times := util.CleanUpAndConvertIntArray(strings.Split(lines[0], " ")[1:])
	distances := util.CleanUpAndConvertIntArray(strings.Split(lines[1], " ")[1:])

	sum := 0

	for i := 0; i < len(times); i++ {
		solutions := findSolutions(times[i], distances[i])

		if sum == 0 {
			sum = solutions
			continue
		}

		sum *= solutions
	}

	return fmt.Sprint(sum)
}

func assembleNum(nums []int) int {
	numberStr := ""

	for _, num := range nums {
		numberStr += fmt.Sprint(num)
	}

	return util.StringToInt(numberStr)
}

func solvePartTwo(input string) string {
	lines := strings.Split(input, "\n")

	times := util.CleanUpAndConvertIntArray(strings.Split(lines[0], " ")[1:])
	distances := util.CleanUpAndConvertIntArray(strings.Split(lines[1], " ")[1:])

	solution := findSolutions(assembleNum(times), assembleNum(distances))

	return fmt.Sprint(solution)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
