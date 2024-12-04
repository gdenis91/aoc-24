package day4

import (
	_ "embed"
	"strconv"

	"github.com/gdenis91/aoc-24/aoc"
)

func Part2(input string) (string, error) {
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
			if grid[y][x] == 'M' || grid[y][x] == 'S' {
				if checkMAS(grid, x, y) {
					xmasCount++
				}
			}
		}
	}

	return strconv.Itoa(xmasCount), nil
}

func checkMAS(grid [][]rune, x, y int) bool {
	word := "MAS"
	if grid[y][x] == 'S' {
		word = "SAM"
	}
	for i, c := range word {
		nx, ny := x+i, y+i
		if nx >= len(grid[0]) ||
			ny >= len(grid) ||
			grid[ny][nx] != c {
			return false
		}
	}

	word = "MAS"
	x += 2
	if grid[y][x] == 'S' {
		word = "SAM"
	}
	for i, c := range word {
		nx, ny := x-i, y+i
		if nx < 0 ||
			ny >= len(grid) ||
			grid[ny][nx] != c {
			return false
		}
	}

	return true
}
