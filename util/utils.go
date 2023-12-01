package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Some helper functions to speed up challenge progress

// ReadInput reads the challenge file from filesystem
func ReadInput(part string) string {
	content, err := os.ReadFile(fmt.Sprintf("day%s/day%s.txt", part, part))

	if nil != err {
		panic(fmt.Sprintf("Missing day%s.txt input file! Error: %v", part, err))
	}

	return strings.TrimRight(string(content), "\n")
}

// PrintResult prints the result preformatted to the console.
func PrintResult[T any](part int, result T) {
	fmt.Println(fmt.Sprintf("Part %d result is: %v", part, result))
}

// StringToInt converts the passed string to an int.
// If the string cannot be converted, a panic will be thrown.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if nil != err {
		panic(err)
	}

	return i
}
