package day4

import (
	_ "embed"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

//go:embed sample.txt
var SampleInput string

func Part1(input string) (string, error) {
	lines := aoc.SplitLines(input)
	grid := make([][]rune, len(lines))
	for i, l := range lines {
		for _, v := range l {
			grid[i] = append(grid[i], v)
		}
	}

	xmasCount := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'X' {
				xmasCount += checkXMAS(grid, x, y)
			}
		}
	}

	return strconv.Itoa(xmasCount), nil
}

func checkXMAS(grid [][]rune, x, y int) int {
	directions := []struct{ dx, dy int }{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	count := 0
	for _, d := range directions {
		found := true
		for i, c := range "XMAS" {
			nx, ny := x+d.dx*i, y+d.dy*i
			if nx < 0 || ny < 0 ||
				nx >= len(grid[0]) || ny >= len(grid) ||
				grid[ny][nx] != c {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}
	return count
}
