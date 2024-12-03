package day2

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

func Part2(input string) (string, error) {
	lines := aoc.SplitLines(input)
	var unsafe int
LINES:
	for l := range slices.Values(lines) {
		levels, err := aoc.IntFields(l)
		if err != nil {
			return "", fmt.Errorf("int fields: %w", err)
		}

		if isSafe(levels) {
			continue
		}
		for i := range len(levels) {
			if isSafe(aoc.RemoveAt(levels, i)) {
				continue LINES
			}
		}
		unsafe++
	}
	return strconv.Itoa(len(lines) - unsafe), nil
}

func isSafe(levels []int) bool {
	asc := levels[0] < levels[len(levels)-1]
	for l := 1; l < len(levels); l++ {
		diff := aoc.Diff(levels[l], levels[l-1])
		if (asc && levels[l] < levels[l-1]) ||
			(!asc && levels[l] > levels[l-1]) ||
			(diff < 1 || diff > 3) {
			return false
		}
	}
	return true
}
