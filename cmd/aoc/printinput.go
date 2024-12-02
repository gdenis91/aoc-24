package main

import (
	"flag"
	"fmt"

	"github.com/gdenis91/aoc-24/aoc"
)

type cmdPrintInput struct {
	day  int
	part int
}

func (c *cmdPrintInput) run() error {
	input, err := aoc.GetInput(c.day, c.part)
	if err != nil {
		return fmt.Errorf("aoc get input: %w", err)
	}
	fmt.Println(input)
	return nil
}

func (c *cmdPrintInput) name() string {
	return "input"
}

func (c *cmdPrintInput) description() string {
	return "Print the input for the given day"
}

func (c *cmdPrintInput) flagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(c.name(), flag.ExitOnError)
	flagSet.IntVar(&c.day, "day", 1, "The day for which to print the input")
	flagSet.IntVar(&c.part, "part", 1, "The part for which to print the input")
	return flagSet
}
