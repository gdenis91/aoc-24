package day2

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

//go:embed sample.txt
var SampleInput string

func Part1(input string) (string, error) {
	lines := aoc.SplitLines(input)
	var unsafe int
	for i := range lines {
		levels, err := aoc.IntFields(lines[i])
		if err != nil {
			return "", fmt.Errorf("int fields: %w", err)
		}

		shouldAsc := levels[0] < levels[len(levels)-1]
		for l := 1; l < len(levels); l++ {
			diff := aoc.Diff(levels[l], levels[l-1])
			if (shouldAsc && levels[l] < levels[l-1]) ||
				(!shouldAsc && levels[l] > levels[l-1]) ||
				(diff < 1 || diff > 3) {
				unsafe++
				break
			}
		}
	}
	return strconv.Itoa(len(lines) - unsafe), nil
}
