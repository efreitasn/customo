package customo

import (
	"strings"
)

// These flags are added to the outputted string the order below.
const (
	FBold = 1 << iota
	FItalic
	FBlink
	FCrossed
)

type C struct {
	flags int
}

func New() *C {
	return &C{}
}

func (c *C) AddFlags(flags int) {
	c.flags |= flags
}

func (c *C) FormatString(str string) string {
	var b strings.Builder

	if c.flags == 0 {
		return str
	}

	c.writeFlags(&b)
	b.WriteString(str)
	b.WriteString("\x1b[0m")

	return b.String()
}

func (c *C) writeFlags(b *strings.Builder) {
	atLeastOneFlag := false
	b.WriteString("\x1b[")

	// Bold
	if c.flags&FBold != 0 {
		if atLeastOneFlag {
			b.WriteString(";1")
		} else {
			atLeastOneFlag = true
			b.WriteString("1")
		}
	}

	// Italic
	if c.flags&FItalic != 0 {
		if atLeastOneFlag {
			b.WriteString(";3")
		} else {
			atLeastOneFlag = true
			b.WriteString("3")
		}
	}

	// Blink
	if c.flags&FBlink != 0 {
		if atLeastOneFlag {
			b.WriteString(";5")
		} else {
			atLeastOneFlag = true
			b.WriteString("5")
		}
	}

	// Crossed
	if c.flags&FCrossed != 0 {
		if atLeastOneFlag {
			b.WriteString(";9")
		} else {
			atLeastOneFlag = true
			b.WriteString("9")
		}
	}

	b.WriteString("m")
}

// Test purposes
func escapeStr(str string) string {
	return strings.ReplaceAll(str, "\x1b", "\\x1b")
}
