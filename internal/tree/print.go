package tree

import (
	"fmt"
	"io"
	"strings"
)

type Layout int

const (
	RightCenter Layout = iota
	RightTop
	RightDown
	TopDown
)

type PrintFunc func(*Node) string

func Print(in *Node, out io.Writer, l Layout) (err error) {
	fn, ok := allPrinters()[l]
	if !ok {
		fn = printTopDown
	}

	_, err = fmt.Fprint(out, fn(in))
	return
}

func allPrinters() map[Layout]PrintFunc {
	return map[Layout]PrintFunc{
		TopDown:     printTopDown,
		RightTop:    printRightTop,
		RightCenter: printRightCenter,
		RightDown:   printRightBottom,
	}
}

func printTopDown(root *Node) string {
	return printTreeBranchStyle(root,
		" ├─ ",
		" │  ",
		" ╰─ ",
		"    ")
}

func printRightTop(root *Node) string {
	display := buildTreeView(root, Top)
	var sb strings.Builder
	for _, line := range display.content {
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func printRightCenter(root *Node) string {
	display := buildTreeView(root, Center)
	var sb strings.Builder
	for _, line := range display.content {
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func printRightBottom(root *Node) string {
	display := buildTreeView(root, Bottom)
	var sb strings.Builder
	for _, line := range display.content {
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	return sb.String()
}

// printTreeBranchStyle recursively prints a tree in a classic branch style.
// branch : string prefix for non-last children (e.g., " ├─")
// branchCont : string prefix for continuing vertical lines (e.g., " │ ")
// lastBranch : string prefix for last child (e.g., " └─")
// lastBranchCont : string prefix for continuing after last child (e.g., " ")
// Returns a string with the tree formatted in a branch style.
func printTreeBranchStyle(root *Node, branch, branchCont, lastBranch, lastBranchCont string) string {
	var (
		sb    strings.Builder
		recur func(node *Node, prefix string)
	)

	// recur is a helper function that traverses nodes recursively.
	recur = func(node *Node, prefix string) {
		for i, child := range node.Children {
			last := i == len(node.Children)-1
			if !last {
				// Print non-last child with branch prefix
				sb.WriteString(prefix + branch + child.Content + "\n")
				// Recur with continuation prefix
				recur(child, prefix+branchCont)
			} else {
				// Print last child with last branch prefix
				sb.WriteString(prefix + lastBranch + child.Content + "\n")
				// Recur with last branch continuation
				recur(child, prefix+lastBranchCont)
			}
		}
	}

	// Print top-level children without any prefix
	for _, child := range root.Children {
		sb.WriteString(child.Content + "\n")
		recur(child, "")
	}

	return sb.String()
}
