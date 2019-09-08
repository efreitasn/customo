package customo

import (
	"strconv"
	"strings"
)

// Attributes
const (
	AttrBold                = 1
	AttrItalic              = 3
	AttrBlink               = 5
	AttrCrossed             = 9
	AttrFg4BitBlack         = 30
	AttrFg4BitRed           = 31
	AttrFg4BitGreen         = 32
	AttrFg4BitYellow        = 33
	AttrFg4BitBlue          = 34
	AttrFg4BitMagenta       = 35
	AttrFg4BitCyan          = 36
	AttrFg4BitWhite         = 37
	AttrBg4BitBlack         = 40
	AttrBg4BitRed           = 41
	AttrBg4BitGreen         = 42
	AttrBg4BitYellow        = 43
	AttrBg4BitBlue          = 44
	AttrBg4BitMagenta       = 45
	AttrBg4BitCyan          = 46
	AttrBg4BitWhite         = 47
	AttrFg4BitBrightBlack   = 90
	AttrFg4BitBrightRed     = 91
	AttrFg4BitBrightGreen   = 92
	AttrFg4BitBrightYellow  = 93
	AttrFg4BitBrightBlue    = 94
	AttrFg4BitBrightMagenta = 95
	AttrFg4BitBrightCyan    = 96
	AttrFg4BitBrightWhite   = 97
	AttrBg4BitBrightBlack   = 100
	AttrBg4BitBrightRed     = 101
	AttrBg4BitBrightGreen   = 102
	AttrBg4BitBrightYellow  = 103
	AttrBg4BitBrightBlue    = 104
	AttrBg4BitBrightMagenta = 105
	AttrBg4BitBrightCyan    = 106
	AttrBg4BitBrightWhite   = 107
)

// Format formats str using the added attributes.
func Format(str string, attrs ...int) string {
	var b strings.Builder

	if len(attrs) == 0 {
		return str
	}

	writeAttrs(&b, attrs)
	b.WriteString(str)
	b.WriteString("\x1b[0m")

	return b.String()
}

func writeAttrs(b *strings.Builder, attrs []int) {
	b.WriteString("\x1b[")

	for i, attr := range attrs {
		b.WriteString(strconv.Itoa(attr))

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
