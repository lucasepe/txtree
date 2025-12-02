package tree

import (
	"encoding/json"
	"fmt"
	"sort"
)

// FromJSON crea un *Node da un JSON arbitrario
// sortKeys: se true, le chiavi degli oggetti saranno ordinate alfabeticamente
func FromJSON(in []byte, sortKeys bool) (*Node, error) {
	var obj any

	if err := json.Unmarshal(in, &obj); err != nil {
		return nil, err
	}

	root := &Node{Content: ""}
	buildJSONNode(root, obj, sortKeys)
	return root, nil
}

func buildJSONNode(parent *Node, value any, sortKeys bool) {
	switch v := value.(type) {

	case map[string]any:
		// Oggetto JSON
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		if sortKeys {
			sort.Strings(keys)
		}

		for _, key := range keys {
			val := v[key]
			if leaf := jsonLeaf(val); leaf != "" {
				parent.Children = append(parent.Children,
					&Node{Content: " " + key + ": " + leaf + " "})
			} else {
				child := &Node{Content: " " + key + " "}
				parent.Children = append(parent.Children, child)
				buildJSONNode(child, val, sortKeys)
			}
		}

	case []any:
		// Array JSON
		for i, val := range v {
			label := fmt.Sprintf("[%d]", i)
			if leaf := jsonLeaf(val); leaf != "" {
				parent.Children = append(parent.Children,
					&Node{Content: " " + label + ": " + leaf + " "})
			} else {
				child := &Node{Content: label}
				parent.Children = append(parent.Children, child)
				buildJSONNode(child, val, sortKeys)
			}
		}
	}
}

func jsonLeaf(v any) string {
	switch x := v.(type) {
	case string:
		return fmt.Sprintf("%q", x)
	case float64:
		return fmt.Sprintf("%g", x)
	case bool:
		return fmt.Sprintf("%v", x)
	case nil:
		return "null"
	default:
		return ""
	}
}
