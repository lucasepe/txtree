package cmd

import (
	"flag"
)

type Options struct {
	SortKeys bool
	Layout   int
	Format   Format
}

func Configure(fs *flag.FlagSet, vals *FlagValues, args []string) Options {
	if err := fs.Parse(args); err != nil {
		return Options{}
	}

	opts := Options{
		SortKeys: *vals.SortKeys,
		Format:   *vals.Format,
		Layout:   *vals.Layout,
	}

	if opts.Layout < 0 {
		opts.Layout = 0
	}

	if opts.Layout > 4 {
		opts.Layout = 0
	}

	return opts
}
