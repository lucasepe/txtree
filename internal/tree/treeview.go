package tree

import (
	"strings"
)

// Node rappresenta un nodo dell'albero.
type Node struct {
	Content  string
	Parent   *Node
	Children []*Node
}

type Align int

const (
	Top Align = iota
	Center
	Bottom
)

// treeView represents a "pixel" view of a tree for right-aligned printing.
// Each string in content is a line, and entrance is the vertical line where
// parent connects to its children.
type treeView struct {
	entrance int
	content  []string
}

// buildTreeView recursively constructs a displayTree from a Node.
// It traverses the Node tree and merges all children into a right-aligned representation.
func buildTreeView(tree *Node, align Align) treeView {
	children := make([]treeView, len(tree.Children))
	for i, ch := range tree.Children {
		children[i] = buildTreeView(ch, align)
	}
	return mergeTreeView(tree.Content, children, align)
}

// mergeTreeView merges a node's content with its children into a displayTree.
// It preserves the parent-child connectors and aligns the parent line according to
// the specified Align (Top, Center, Bottom).
func mergeTreeView(content string, children []treeView, align Align) treeView {
	contentRunes := []rune(content)
	contentLen := len(contentRunes)

	// Prepare prefix spaces for children lines
	spacePrefix := strings.Repeat(" ", contentLen+3)

	ret := treeView{}

	// Leaf node: no children, just return content
	if len(children) == 0 {
		ret.content = []string{content}
		ret.entrance = 0
		return ret
	}

	// Flatten children lines and prepend the space prefix
	for _, child := range children {
		for _, line := range child.content {
			ret.content = append(ret.content, spacePrefix+line)
		}
	}

	// Determine the entrance line for the parent based on alignment
	switch align {
	case Top:
		ret.entrance = 0
	case Center:
		ret.entrance = len(ret.content) / 2
	case Bottom:
		ret.entrance = len(ret.content) - 1
	}

	// Insert parent content into the entrance line
	parentLine := []rune(ret.content[ret.entrance])
	parentLine = ensureRunesLen(string(parentLine), contentLen)
	copy(parentLine, contentRunes)

	ret.content[ret.entrance] = string(parentLine)

	// Compute first and last entrances among children (vertical range)
	firstEntrance := children[0].entrance
	lastEntrance := 0
	y := 0
	for _, child := range children {
		lastEntrance = y + child.entrance
		y += len(child.content)
	}

	// Draw vertical and horizontal connectors for each child
	y = 0
	for _, child := range children {
		start := y
		for range child.content {
			if y >= firstEntrance && y <= lastEntrance {
				// Draw vertical connector │
				lineRunes := []rune(ret.content[y])
				lineRunes = ensureRunesLen(string(lineRunes), contentLen+2)
				lineRunes[contentLen+1] = '│'
				ret.content[y] = string(lineRunes)
			}
			y++
		}

		// Refine the connector at the child's entrance
		childEntrance := start + child.entrance
		lineRunes := []rune(ret.content[childEntrance])
		lineRunes = ensureRunesLen(string(lineRunes), contentLen+3)

		switch {
		case firstEntrance == lastEntrance:
			lineRunes[contentLen+1] = '─'
		case childEntrance == firstEntrance:
			lineRunes[contentLen+1] = '╭' // '┌'
		case childEntrance < lastEntrance:
			lineRunes[contentLen+1] = '├'
		default:
			lineRunes[contentLen+1] = '╰' //'└'
		}

		// Draw horizontal connector to the child
		lineRunes[contentLen+2] = '─'
		ret.content[childEntrance] = string(lineRunes)
	}

	// Draw horizontal connector from parent
	lineRunes := []rune(ret.content[ret.entrance])
	lineRunes = ensureRunesLen(string(lineRunes), contentLen+2)
	lineRunes[contentLen] = '─'

	// Refine parent connector based on existing character
	switch lineRunes[contentLen+1] {
	case '─':
		lineRunes[contentLen+1] = '─'
	case '╭': //'┌':
		lineRunes[contentLen+1] = '┬'
	case '├':
		lineRunes[contentLen+1] = '┼'
	case '╰': //'└':
		lineRunes[contentLen+1] = '┴'
	case '│':
		lineRunes[contentLen+1] = '┤'
	default:
		lineRunes[contentLen+1] = '┤'
	}
	ret.content[ret.entrance] = string(lineRunes)

	return ret
}

// ensureRunesLen returns a slice of runes of length at least n.
// It copies the runes from s and, if necessary, pads with spaces
// to reach the desired length.
func ensureRunesLen(s string, n int) []rune {
	rs := []rune(s)
	if len(rs) >= n {
		return rs
	}
	// create padding runes filled with spaces
	pad := make([]rune, n-len(rs))
	for i := range pad {
		pad[i] = ' '
	}
	// append padding to the original runes
	return append(rs, pad...)
}
