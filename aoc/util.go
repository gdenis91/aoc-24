package aoc

import "strings"

func SplitLines(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}
