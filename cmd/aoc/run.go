package main

import "flag"

type cmdRun struct {
	day int
}

func (c *cmdRun) run() error {
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
	return flagSet
}
