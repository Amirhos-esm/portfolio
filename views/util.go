package views

import (
	"html"
	"strings"
)

func FormatDescription(s string) string {
	parts := strings.Split(s, "\n\n")
	var out strings.Builder
	for _, p := range parts {
		out.WriteString("<p>")
		out.WriteString(html.EscapeString(p))
		out.WriteString("</p>")
	}
	return out.String()
}
