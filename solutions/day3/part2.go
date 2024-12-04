package day3

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

var conditionalMulPattern = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func Part2(input string) (string, error) {
	lines := aoc.SplitLines(input)
	var sum int
	enabled := true
	for _, line := range lines {
		matches := conditionalMulPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
				continue
			} else if match[0] == "don't()" {
				enabled = false
				continue
			} else if !enabled {
				continue
			}
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
