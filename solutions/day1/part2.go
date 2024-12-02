package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdenis91/aoc-24/aoc"
)

func Part2(input string) (string, error) {
	lines := aoc.SplitLines(input)
	num := make([]int, len(lines))
	times := make(map[int]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		p0, err := strconv.Atoi(fields[0])
		if err != nil {
			return "", fmt.Errorf("strconv atoi field0: %w", err)
		}
		num[i] = p0

		p1, err := strconv.Atoi(fields[1])
		if err != nil {
			return "", fmt.Errorf("strconv atoi field1: %w", err)
		}
		times[p1]++
	}

	var score int
	for _, n := range num {
		score += (n * times[n])
	}

	return fmt.Sprint(score), nil
}
