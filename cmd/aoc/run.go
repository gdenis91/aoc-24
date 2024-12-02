package main

import (
	"flag"
	"fmt"

	"github.com/gdenis91/aoc-24/aoc"
	"github.com/gdenis91/aoc-24/solutions"
)

type cmdRun struct {
	day  int
	part int
}

func (c *cmdRun) run() error {
	daySln := solutions.Solutions[c.day]
	if daySln == nil {
		return fmt.Errorf(
			"no solutions for day %d",
			c.day,
		)
	}

	sln := daySln[c.part]
	if sln == nil {
		return fmt.Errorf(
			"no solution for day %d of part %d",
			c.day,
			c.part,
		)
	}

	input, err := aoc.GetInput(c.day, c.part)
	if err != nil {
		return fmt.Errorf(
			"aoc get input for day %d part %d: %w",
			c.day,
			c.part,
			err,
		)
	}

	result, err := sln(input)
	if err != nil {
		return fmt.Errorf(
			"run solution day %d part %d: %w",
			c.day,
			c.part,
			err,
		)
	}
	fmt.Println(result)

	return nil
}

func (c *cmdRun) name() string {
	return "run"
}

func (c *cmdRun) description() string {
	return "Run the solution for the given day"
}

func (c *cmdRun) flagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(c.name(), flag.ExitOnError)
	flagSet.IntVar(&c.day, "day", 0, "The day for which to run the solution")
	flagSet.IntVar(
		&c.part,
		"part",
		0,
		"The part of the day for which to run the solution",
	)
	return flagSet
}
