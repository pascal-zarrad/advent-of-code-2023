package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pascal-zarrad/advent-of-code-2023/util"
)

const day = "08"

////////// COMMON

const networkNodeConversionRegex = `([A-Z0-9]+)\s=\s\(([A-Z0-9]+), ([A-Z0-9]+)\)`

type NetworkNode struct {
	id    string
	left  string
	right string
}

func extractNetworkMap(rawNodes []string) map[string]NetworkNode {
	networkNodes := map[string]NetworkNode{}

	networkNodeConversionRegExpr := regexp.MustCompile(networkNodeConversionRegex)

	for _, rawNode := range rawNodes {
		if rawNode == "" {
			continue
		}

		parts := networkNodeConversionRegExpr.FindAllStringSubmatch(rawNode, -1)

		networkNode := NetworkNode{
			id:    parts[0][1],
			left:  parts[0][2],
			right: parts[0][3],
		}

		networkNodes[networkNode.id] = networkNode
	}

	return networkNodes
}

////////// PART 1

const (
	PART_1_STARTING_POINT  = "AAA"
	PART_1_FINISHING_POINT = "ZZZ"
)

func solvePartOne(input string) string {
	lines := strings.Split(input, "\n")

	instructions := []rune(lines[0])
	networkMap := extractNetworkMap(lines[2:])

	result := simulatePathPart1(instructions, networkMap)

	return fmt.Sprint(result)
}

func simulatePathPart1(instructions []rune, networkMap map[string]NetworkNode) int {
	currentNode := networkMap[PART_1_STARTING_POINT]
	endNode := networkMap[PART_1_FINISHING_POINT]

	stepsTaken := 0

	for currentNode.id != endNode.id {
		for i := 0; i < len(instructions); i++ {
			switch instructions[i] {
			case 'L':
				currentNode = networkMap[currentNode.left]
			case 'R':
				currentNode = networkMap[currentNode.right]
			}

			stepsTaken++

			if currentNode.id == endNode.id {
				break
			}
		}
	}

	return stepsTaken
}

////////// PART 2

const (
	PART_2_STARTING_INDICATOR  = "A"
	PART_2_FINISHING_INDICATOR = "Z"
)

func solvePartTwo(input string) string {
	lines := strings.Split(input, "\n")

	instructions := []rune(lines[0])
	networkMap := extractNetworkMap(lines[2:])

	stepsList := simulatePathPart2(instructions, networkMap)

	return fmt.Sprint(lcm(stepsList...))
}

func simulatePathPart2(instructions []rune, networkMap map[string]NetworkNode) []uint64 {
	currentNodes := make([]NetworkNode, 0)

	for _, networkNode := range networkMap {
		if !strings.HasSuffix(networkNode.id, PART_2_STARTING_INDICATOR) {
			continue
		}

		currentNodes = append(currentNodes, networkNode)
	}

	stepsTakenList := make([]uint64, 0)

	for _, currentNode := range currentNodes {
		stepsTaken := uint64(0)

		for !strings.HasSuffix(currentNode.id, PART_2_FINISHING_INDICATOR) {
			for i := 0; i < len(instructions); i++ {
				switch instructions[i] {
				case 'L':
					currentNode = networkMap[currentNode.left]
				case 'R':
					currentNode = networkMap[currentNode.right]
				}

				stepsTaken++

				if strings.HasSuffix(currentNode.id, PART_2_FINISHING_INDICATOR) {
					break
				}
			}
		}

		stepsTakenList = append(stepsTakenList, stepsTaken)
	}

	return stepsTakenList
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(numbers ...uint64) uint64 {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcmPair(result, numbers[i])
	}

	return result
}

func lcmPair(a, b uint64) uint64 {
	return a * b / gcd(a, b)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
