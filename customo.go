package customo

import (
	"errors"
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

// Foreground colors.
const (
	Fg4BitBlack         = 30
	Fg4BitBrightBlack   = 90
	Fg4BitRed           = 31
	Fg4BitBrightRed     = 91
	Fg4BitGreen         = 32
	Fg4BitBrightGreen   = 92
	Fg4BitYellow        = 33
	Fg4BitBrightYellow  = 93
	Fg4BitBlue          = 34
	Fg4BitBrightBlue    = 94
	Fg4BitMagenta       = 35
	Fg4BitBrightMagenta = 95
	Fg4BitCyan          = 36
	Fg4BitBrightCyan    = 96
	Fg4BitWhite         = 37
	Fg4BitBrightWhite   = 97
)

// Background colors.
const (
	Bg4BitBlack         = 40
	Bg4BitBrightBlack   = 100
	Bg4BitRed           = 41
	Bg4BitBrightRed     = 101
	Bg4BitGreen         = 42
	Bg4BitBrightGreen   = 102
	Bg4BitYellow        = 43
	Bg4BitBrightYellow  = 103
	Bg4BitBlue          = 44
	Bg4BitBrightBlue    = 104
	Bg4BitMagenta       = 45
	Bg4BitBrightMagenta = 105
	Bg4BitCyan          = 46
	Bg4BitBrightCyan    = 106
	Bg4BitWhite         = 47
	Bg4BitBrightWhite   = 107
)

// C is the representation of attributes to be used in a string.
type C struct {
	attrs int
	bg    int
	fg    int
}

// New returns a new representation of attributes to be used in a string.
// In order to no foreground color to be applied, pass 0 to fg4Bit. The same goes to bg4Bit and attrs.
func New(fg4Bit, bg4Bit, attrs int) (*C, error) {
	// Foreground color checks.
	fg4BitWithinNormalRange := (fg4Bit == 0 || (fg4Bit >= Fg4BitBlack && fg4Bit <= Fg4BitWhite))
	fg4BitWithinBrightRange := (fg4Bit == 0 || (fg4Bit >= Fg4BitBrightBlack && fg4Bit <= Fg4BitBrightWhite))

	if !(fg4BitWithinNormalRange || fg4BitWithinBrightRange) {
		return nil, errors.New("fg4Bit not within the allowed range")
	}

	// Background color checks
	bg4BitWithinNormalRange := (bg4Bit == 0 || (bg4Bit >= Bg4BitBlack && bg4Bit <= Bg4BitWhite))
	bg4BitWithinBrightRange := (bg4Bit == 0 || (bg4Bit >= Bg4BitBrightBlack && bg4Bit <= Bg4BitBrightWhite))

	if !(bg4BitWithinNormalRange || bg4BitWithinBrightRange) {
		return nil, errors.New("bg4Bit not within the allowed range")
	}

	return &C{
		fg:    fg4Bit,
		bg:    bg4Bit,
		attrs: attrs,
	}, nil
}

// MustNew executes New and panics if it returns an error.
func MustNew(fg4Bit, bg4Bit, attrs int) *C {
	c, err := New(fg4Bit, bg4Bit, attrs)

	if err != nil {
		panic(err)
	}

	return c
}

// FormatString formats str using the added attributes.
func (c *C) FormatString(str string) string {
	var b strings.Builder

	if c.attrs == 0 && c.bg == 0 && c.fg == 0 {
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

	// Foreground color.
	if c.fg != 0 {
		if atLeastOneFlag {
			b.WriteString(";" + strconv.Itoa(c.fg))
		} else {
			atLeastOneFlag = true
			b.WriteString(strconv.Itoa(c.fg))
		}
	}

	// Background color.
	if c.bg != 0 {
		if atLeastOneFlag {
			b.WriteString(";" + strconv.Itoa(c.bg))
		} else {
			atLeastOneFlag = true
			b.WriteString(strconv.Itoa(c.bg))
		}
	}

	b.WriteString("m")
}

// Test purposes.
func escapeStr(str string) string {
	return strings.ReplaceAll(str, "\x1b", "\\x1b")
}
