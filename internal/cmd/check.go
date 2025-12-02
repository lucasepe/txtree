package cmd

import (
	"fmt"
	"os"
)

// CheckErr check for an error and eventually exit
func CheckErr(msg string, err error) {
	if err == nil {
		return
	}

	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	os.Exit(1)
}
