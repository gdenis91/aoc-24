package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gdenis91/aoc-24/aoc"
)

func Part1(input string) (string, error) {
	lines := aoc.SplitLines(input)
	col0 := make([]int, len(lines))
	col1 := make([]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		p0, err := strconv.Atoi(fields[0])
		if err != nil {
			return "", fmt.Errorf("strconv atoi field0: %w", err)
		}
		col0[i] = p0

		p1, err := strconv.Atoi(fields[1])
		if err != nil {
			return "", fmt.Errorf("strconv atoi field1: %w", err)
		}
		col1[i] = p1
	}

	slices.Sort(col0)
	slices.Sort(col1)

	var total int
	for i := 0; i < len(col0); i++ {
		diff := diff(col0[i], col1[i])
		total += diff
	}

	return fmt.Sprint(total), nil
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
