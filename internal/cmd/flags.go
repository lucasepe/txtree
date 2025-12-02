package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	Auto Format = iota + 1
	Text
	JSON
)

type FlagValues struct {
	SortKeys *bool
	Format   *Format
	Layout   *int
	ShowHelp *bool
}

// NewFlagSet creates a FlagSet with all supported CLI options.
func NewFlagSet() (*flag.FlagSet, *FlagValues) {
	fs := flag.NewFlagSet("txtree", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	input := Auto

	vals := &FlagValues{
		SortKeys: fs.Bool("s", false, "Sort keys alphabetically (JSON only)"),
		Layout:   fs.Int("l", 0, "Layout: (0=right-center, 1=right-top, 2=right-down, 3=top-down)"),
		Format:   &input,
		ShowHelp: fs.Bool("h", false, "Show help"),
	}

	fs.Var(vals.Format, "i", "Input format (auto, text, json)")

	return fs, vals
}

type Format int

func (f *Format) String() string {
	for name, val := range formatNames {
		if val == *f {
			return name
		}
	}
	return "unknown"
}

func (f *Format) Set(s string) error {
	s = strings.ToLower(s)
	val, ok := formatNames[s]
	if !ok {
		return fmt.Errorf("invalid input format: %s", s)
	}
	*f = val
	return nil
}

var (
	formatNames = map[string]Format{
		"auto": Auto,
		"text": Text,
		"json": JSON,
	}
)
