package text

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

func MergeSideBySide(left, right string, gap int, fillChar string) string {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")

	// calcola larghezza massima del blocco sinistro in colonne Unicode
	maxLeft := 0
	for _, l := range leftLines {
		w := runewidth.StringWidth(l)
		if w > maxLeft {
			maxLeft = w
		}
	}

	// numero massimo di righe tra i due blocchi
	maxLines := len(leftLines)
	if len(rightLines) > maxLines {
		maxLines = len(rightLines)
	}

	if fillChar == "" {
		fillChar = " "
	}

	var out strings.Builder

	for i := 0; i < maxLines; i++ {
		var L, R string
		if i < len(leftLines) {
			L = leftLines[i]
		}
		if i < len(rightLines) {
			R = rightLines[i]
		}

		// padding a destra del blocco sinistro basato su larghezza visibile
		w := runewidth.StringWidth(L)
		if w < maxLeft {
			L += strings.Repeat(fillChar, maxLeft-w)
		}

		// gap orizzontale
		L += strings.Repeat(fillChar, gap)

		out.WriteString(L + R + "\n")
	}

	return out.String()
}
