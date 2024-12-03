package solutions

import (
	"github.com/gdenis91/aoc-24/solutions/day1"
	"github.com/gdenis91/aoc-24/solutions/day2"
)

type solution func(input string) (string, error)

var (
	Solutions = map[int]map[int]solution{
		1: {
			1: day1.Part1,
			2: day1.Part2,
		},
		2: {
			1: day2.Part1,
			2: day2.Part2,
		},
	}

	SampleInput = map[int]string{
		1: day1.SampleInput,
		2: day2.SampleInput,
	}
)
