package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitLines splits a string into lines, trimming whitespace from the start and end of the string.
func SplitLines(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

// Diff returns the absolute difference between a and b as an integer.
func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// IntFields converts a string of space-separated integers into a slice of integers.
// If any of the fields cannot be converted to an integer, an error is returned.
func IntFields(s string) ([]int, error) {
	fields := strings.Fields(s)
	values := make([]int, len(fields))
	for l := range fields {
		lvl, err := strconv.Atoi(fields[l])
		if err != nil {
			return nil, fmt.Errorf("strconv atoi: %w", err)
		}
		values[l] = lvl
	}
	return values, nil
}

// RemoveAt removes the element at index i from the slice values.
func RemoveAt[T any](values []T, i int) []T {
	newSlice := make([]T, len(values)-1)
	copy(newSlice, values[:i])
	copy(newSlice[i:], values[i+1:])
	return newSlice
}
