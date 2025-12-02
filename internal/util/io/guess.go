package io

import "strings"

func LooksLikeJSON(data []byte) bool {
	// Strip UTF-8 BOM if present
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		data = data[3:]
	}

	// Trim spaces
	s := strings.TrimLeftFunc(string(data), func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\r' || r == '\t'
	})

	if len(s) == 0 {
		return false
	}

	// JSON usually starts with { or [
	return s[0] == '{' || s[0] == '['
}
