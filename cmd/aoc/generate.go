package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type cmdGenerate struct {
	day int
}

func (c *cmdGenerate) run() error {
	dayDir := fmt.Sprintf("solutions/day%d", c.day)
	err := os.MkdirAll(dayDir, 0755)
	if err != nil {
		return fmt.Errorf("os mkdirall: %w", err)
	}

	for part := 1; part <= 2; part++ {
		path := fmt.Sprintf("%s/part%d.go", dayDir, part)
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("os create %s: %w", path, err)
		}
		defer f.Close()

		tmpl, err := template.New("part").Parse(partTemplate)
		if err != nil {
			return fmt.Errorf("template parse: %w", err)
		}

		err = tmpl.Execute(f, struct {
			Day  int
			Part int
		}{
			Day:  c.day,
			Part: part,
		})
		if err != nil {
			return fmt.Errorf("template execute: %w", err)
		}
	}

	err = c.updateSolutionsFile()
	if err != nil {
		return fmt.Errorf("update solutions file: %w", err)
	}

	return nil
}

func (c *cmdGenerate) updateSolutionsFile() error {
	solutionsFile := "solutions/solutions.go"
	f, err := os.Create(solutionsFile)
	if err != nil {
		return fmt.Errorf("create solutions file: %w", err)
	}
	defer f.Close()

	tmpl, err := template.New("part").Parse(solutionsTemplate)
	if err != nil {
		return fmt.Errorf("template parse: %w", err)
	}

	dirs, err := os.ReadDir("solutions")
	if err != nil {
		return fmt.Errorf("read solutions directory: %w", err)
	}

	type sln struct {
		Pkg string
		Day string
	}
	var slns []sln

	for _, dir := range dirs {
		if !dir.IsDir() ||
			!strings.HasPrefix(dir.Name(), "day") {
			continue
		}
		slns = append(slns, sln{
			Pkg: dir.Name(),
			Day: strings.TrimPrefix(dir.Name(), "day"),
		})
	}

	err = tmpl.Execute(f, slns)
	if err != nil {
		return fmt.Errorf("template execute: %w", err)
	}

	return nil
}

func (c *cmdGenerate) name() string {
	return "generate"
}

func (c *cmdGenerate) description() string {
	return "Generate new day's files and register the functions in the solutions file"
}

func (c *cmdGenerate) flagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(c.name(), flag.ExitOnError)
	flagSet.IntVar(&c.day, "day", 0, "The day for which to generate files")
	return flagSet
}

const partTemplate = `package day{{.Day}}

func Part{{.Part}}(input string) (string, error) {
    // TODO: Implement solution
    return "", nil
}
`

const solutionsTemplate = `package solutions

import ({{ range $i, $sln := . }}
	"github.com/gdenis91/aoc-24/solutions/day{{ $sln.Day }}"{{ end }}
)

type solution func(input string) (string, error)

var (
	Solutions = map[int]map[int]solution{ {{ range $i, $sln := . }}
		{{ $sln.Day }}: {
			1: {{ $sln.Pkg }}.Part1,
			2: {{ $sln.Pkg }}.Part2,
		},{{ end }}
	}
)
`
