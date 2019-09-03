package customo

import (
	"strings"
)

// These attributes are added to the outputted string in the order below.
const (
	AttrBold = 1 << iota
	AttrItalic
	AttrBlink
	AttrCrossed
)

// C is the representation of attributes to be used in a string.
type C struct {
	attrs int
}

// New returns a new representation of attributes to be used in a string.
func New() *C {
	return &C{}
}

// AddAttrs adds attributes.
func (c *C) AddAttrs(attrs int) {
	c.attrs |= attrs
}

// FormatString formats str using the added attributes.
func (c *C) FormatString(str string) string {
	var b strings.Builder

	if c.attrs == 0 {
		return str
	}

	c.writeAttrs(&b)
	b.WriteString(str)
	b.WriteString("\x1b[0m")

	return b.String()
}

func (c *C) writeAttrs(b *strings.Builder) {
	atLeastOneFlag := false
	b.WriteString("\x1b[")

	// Bold
	if c.attrs&AttrBold != 0 {
		if atLeastOneFlag {
			b.WriteString(";1")
		} else {
			atLeastOneFlag = true
			b.WriteString("1")
		}
	}

	// Italic
	if c.attrs&AttrItalic != 0 {
		if atLeastOneFlag {
			b.WriteString(";3")
		} else {
			atLeastOneFlag = true
			b.WriteString("3")
		}
	}

	// Blink
	if c.attrs&AttrBlink != 0 {
		if atLeastOneFlag {
			b.WriteString(";5")
		} else {
			atLeastOneFlag = true
			b.WriteString("5")
		}
	}

	// Crossed
	if c.attrs&AttrCrossed != 0 {
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
