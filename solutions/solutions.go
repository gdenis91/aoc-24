package solutions

import "github.com/gdenis91/aoc-24/solutions/day1"

type solution func(input string) (string, error)

var (
	Solutions = map[int]map[int]solution{
		1: {
			1: day1.Part1,
			2: day1.Part2,
		},
	}
)
