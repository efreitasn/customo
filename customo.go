// Package customo helps when creating strings using ANSI escape codes attributes, such as italic, bold and custom foreground and background colors.
package customo

import (
	"strings"
)

// Format formats str using attrs.
func Format(str string, attrs ...Attr) string {
	var b strings.Builder

	if len(attrs) == 0 {
		return str
	}

	writeAttrs(&b, attrs)
	b.WriteString(str)
	b.WriteString("\x1b[0m")

	return b.String()
}

func writeAttrs(b *strings.Builder, attrs []Attr) {
	b.WriteString("\x1b[")

	for i, attr := range attrs {
		b.WriteString(attr.ANSI())

		if i != len(attrs)-1 {
			b.WriteString(";")
		}
	}

	b.WriteString("m")
}

// Test purposes.
func escapeStr(str string) string {
	return strings.ReplaceAll(str, "\x1b", "\\x1b")
}
