package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/lucasepe/txtree/internal/cmd"
	"github.com/lucasepe/txtree/internal/tree"
	ioutil "github.com/lucasepe/txtree/internal/util/io"
	textutil "github.com/lucasepe/txtree/internal/util/text"
)

func main() {
	fs, fv := cmd.NewFlagSet()
	fs.Usage = cmd.Usage(fs)

	opts := cmd.Configure(fs, fv, os.Args[1:])

	if *fv.ShowHelp {
		fs.Usage()
		return
	}

	input, err := ioutil.ReadInput(fs.Args())
	if err != nil {
		if errors.Is(err, ioutil.ErrNoInput) {
			fs.Usage()
			return
		}
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	input = textutil.Clean(input, 3)

	var root *tree.Node
	switch opts.Format {
	case cmd.Auto:
		if ioutil.LooksLikeJSON(input) {
			root, err = tree.FromJSON(input, opts.SortKeys)
		} else {
			root, err = tree.FromIndentedText(bytes.NewReader(input), 3)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input: %v\n", err)
		os.Exit(1)
	}

	err = tree.Print(root, os.Stdout, tree.Layout(opts.Layout))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error printing tree: %v\n", err)
		os.Exit(1)
	}
}
