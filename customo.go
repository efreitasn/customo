package customo

import (
	"strconv"
	"strings"
)

// These attributes are added to the outputted string in the order below.
const (
	AttrBold = 1 << iota
	AttrItalic
	AttrBlink
	AttrCrossed
)

// Background colors.
const (
	Bg8BitBlack         = 40
	Bg8BitBrightBlack   = 100
	Bg8BitRed           = 41
	Bg8BitBrightRed     = 101
	Bg8BitGreen         = 42
	Bg8BitBrightGreen   = 102
	Bg8BitYellow        = 43
	Bg8BitBrightYellow  = 103
	Bg8BitBlue          = 44
	Bg8BitBrightBlue    = 104
	Bg8BitMagenta       = 45
	Bg8BitBrightMagenta = 105
	Bg8BitCyan          = 46
	Bg8BitBrightCyan    = 106
	Bg8BitWhite         = 47
	Bg8BitBrightWhite   = 107
)

// Foreground colors.
const (
	Fg8BitBlack         = 30
	Fg8BitBrightBlack   = 90
	Fg8BitRed           = 31
	Fg8BitBrightRed     = 91
	Fg8BitGreen         = 32
	Fg8BitBrightGreen   = 92
	Fg8BitYellow        = 33
	Fg8BitBrightYellow  = 93
	Fg8BitBlue          = 34
	Fg8BitBrightBlue    = 94
	Fg8BitMagenta       = 35
	Fg8BitBrightMagenta = 95
	Fg8BitCyan          = 36
	Fg8BitBrightCyan    = 96
	Fg8BitWhite         = 37
	Fg8BitBrightWhite   = 97
)

// C is the representation of attributes to be used in a string.
type C struct {
	attrs int
	bg    int
	fg    int
	bgSet bool
	fgSet bool
}

// New returns a new representation of attributes to be used in a string.
func New() *C {
	return &C{}
}

// AddAttrs adds attributes.
func (c *C) AddAttrs(attrs int) {
	c.attrs |= attrs
}

// SetBg8BitAttr sets the background attribute using an 8-bit color.
func (c *C) SetBg8BitAttr(color int) {
	c.bgSet = true
	c.bg = color
}

// SetFg8BitAttr sets the background attribute using an 8-bit color.
func (c *C) SetFg8BitAttr(color int) {
	c.fgSet = true
	c.fg = color
}

// FormatString formats str using the added attributes.
func (c *C) FormatString(str string) string {
	var b strings.Builder

	if c.attrs == 0 && !c.bgSet && !c.fgSet {
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

	// Bold.
	if c.attrs&AttrBold != 0 {
		if atLeastOneFlag {
			b.WriteString(";1")
		} else {
			atLeastOneFlag = true
			b.WriteString("1")
		}
	}

	// Italic.
	if c.attrs&AttrItalic != 0 {
		if atLeastOneFlag {
			b.WriteString(";3")
		} else {
			atLeastOneFlag = true
			b.WriteString("3")
		}
	}

	// Blink.
	if c.attrs&AttrBlink != 0 {
		if atLeastOneFlag {
			b.WriteString(";5")
		} else {
			atLeastOneFlag = true
			b.WriteString("5")
		}
	}

	// Crossed.
	if c.attrs&AttrCrossed != 0 {
		if atLeastOneFlag {
			b.WriteString(";9")
		} else {
			atLeastOneFlag = true
			b.WriteString("9")
		}
	}

	// Background color.
	if c.bgSet {
		if atLeastOneFlag {
			b.WriteString(";" + strconv.Itoa(c.bg))
		} else {
			atLeastOneFlag = true
			b.WriteString(strconv.Itoa(c.bg))
		}
	}

	// Foreground color.
	if c.fgSet {
		if atLeastOneFlag {
			b.WriteString(";" + strconv.Itoa(c.fg))
		} else {
			atLeastOneFlag = true
			b.WriteString(strconv.Itoa(c.fg))
		}
	}

	b.WriteString("m")
}

// Test purposes.
func escapeStr(str string) string {
	return strings.ReplaceAll(str, "\x1b", "\\x1b")
}
