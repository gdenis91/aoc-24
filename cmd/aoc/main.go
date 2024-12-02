package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gdenis91/aoc-24/aoc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide a command")
		os.Exit(1)
	}

	if aoc.GetSessionCookie() == "" {
		aoc.PrintSessionHelp()
		os.Exit(1)
	}

	cmds := []cmd{
		&cmdPrintInput{},
		&cmdRun{},
	}

	for _, c := range cmds {
		if c.name() == os.Args[1] {
			flagSet := c.flagSet()
			flagSet.Parse(os.Args[2:])
			if err := c.run(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Println("Unknown command")
}

type cmd interface {
	run() error
	name() string
	description() string
	flagSet() *flag.FlagSet
}
