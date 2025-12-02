package tree

import (
	"bufio"
	"io"
	"strings"
)

// Line is used internally during parsing.
type Line struct {
	Spaces  int // indentation in “virtual spaces”
	Content string
	Tree    *Node
}

// FromIndentedText reads indented text from r and returns a tree with a dummy root.
// Tabs are expanded to tabWidth spaces (default 4 recommended).
func FromIndentedText(r io.Reader, tabWidth int) (*Node, error) {
	lines := make([]Line, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		raw := scanner.Text()

		virtualSpaces := 0
		for i := 0; i < len(raw); i++ {
			if raw[i] == ' ' {
				virtualSpaces++
			} else if raw[i] == '\t' {
				virtualSpaces += tabWidth - (virtualSpaces % tabWidth)
			} else {
				break // esce dal ciclo dei caratteri
			}
		}

		content := strings.TrimSpace(raw)
		if content == "" {
			continue
		}
		lines = append(lines, Line{
			Spaces:  virtualSpaces,
			Content: " " + content + " ",
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	root := &Node{}
	for i := range lines {
		child := &Node{Content: lines[i].Content}
		lines[i].Tree = child

		// Find parent: nearest previous line with fewer virtual spaces
		parentFound := false
		for j := i - 1; j >= 0; j-- {
			if lines[j].Spaces < lines[i].Spaces {
				child.Parent = lines[j].Tree
				child.Parent.Children = append(child.Parent.Children, child)
				parentFound = true
				break
			}
		}
		if !parentFound {
			child.Parent = root
			root.Children = append(root.Children, child)
		}
	}

	return root, nil
}
