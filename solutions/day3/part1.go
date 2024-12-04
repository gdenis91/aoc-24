package day3

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

//go:embed sample.txt
var SampleInput string

var mulPattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func Part1(input string) (string, error) {
	lines := aoc.SplitLines(input)
	var sum int
	for _, line := range lines {
		matches := mulPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				return "", fmt.Errorf("strconv atoi: %w", err)
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				return "", fmt.Errorf("strconv atoi: %w", err)
			}
			sum += x * y
		}
	}
	return strconv.Itoa(sum), nil
}
