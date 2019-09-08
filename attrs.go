package customo

import "strconv"

// Attr represents an ANSI escape codes attribute.
type Attr interface {
	ANSI() string
}

// SimpleAttr is a simple attribute.
type SimpleAttr int

// ANSI returns a representation to be used as an ANSI escape code.
func (s SimpleAttr) ANSI() string {
	return strconv.Itoa(int(s))
}

// Attributes.
const (
	AttrBold                     = SimpleAttr(1)
	AttrItalic                   = SimpleAttr(3)
	AttrUnderline                = SimpleAttr(4)
	AttrBlink                    = SimpleAttr(5)
	AttrStrikethrough            = SimpleAttr(9)
	AttrFgColor4BitBlack         = SimpleAttr(30)
	AttrFgColor4BitRed           = SimpleAttr(31)
	AttrFgColor4BitGreen         = SimpleAttr(32)
	AttrFgColor4BitYellow        = SimpleAttr(33)
	AttrFgColor4BitBlue          = SimpleAttr(34)
	AttrFgColor4BitMagenta       = SimpleAttr(35)
	AttrFgColor4BitCyan          = SimpleAttr(36)
	AttrFgColor4BitWhite         = SimpleAttr(37)
	AttrBgColor4BitBlack         = SimpleAttr(40)
	AttrBgColor4BitRed           = SimpleAttr(41)
	AttrBgColor4BitGreen         = SimpleAttr(42)
	AttrBgColor4BitYellow        = SimpleAttr(43)
	AttrBgColor4BitBlue          = SimpleAttr(44)
	AttrBgColor4BitMagenta       = SimpleAttr(45)
	AttrBgColor4BitCyan          = SimpleAttr(46)
	AttrBgColor4BitWhite         = SimpleAttr(47)
	AttrFgColor4BitBrightBlack   = SimpleAttr(90)
	AttrFgColor4BitBrightRed     = SimpleAttr(91)
	AttrFgColor4BitBrightGreen   = SimpleAttr(92)
	AttrFgColor4BitBrightYellow  = SimpleAttr(93)
	AttrFgColor4BitBrightBlue    = SimpleAttr(94)
	AttrFgColor4BitBrightMagenta = SimpleAttr(95)
	AttrFgColor4BitBrightCyan    = SimpleAttr(96)
	AttrFgColor4BitBrightWhite   = SimpleAttr(97)
	AttrBgColor4BitBrightBlack   = SimpleAttr(100)
	AttrBgColor4BitBrightRed     = SimpleAttr(101)
	AttrBgColor4BitBrightGreen   = SimpleAttr(102)
	AttrBgColor4BitBrightYellow  = SimpleAttr(103)
	AttrBgColor4BitBrightBlue    = SimpleAttr(104)
	AttrBgColor4BitBrightMagenta = SimpleAttr(105)
	AttrBgColor4BitBrightCyan    = SimpleAttr(106)
	AttrBgColor4BitBrightWhite   = SimpleAttr(107)
)

// FgColor8Bit is an 8 bit foreground color attribute.
type FgColor8Bit int

// AttrFgColor8Bit returns a new 8 bit foreground color attribute.
func AttrFgColor8Bit(color int) FgColor8Bit {
	return FgColor8Bit(color)
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c FgColor8Bit) ANSI() string {
	return "38;5;" + strconv.Itoa(int(c))
}

// BgColor8Bit is an 8 bit background color attribute.
type BgColor8Bit int

// AttrBgColor8Bit returns a new 8 bit background color attribute.
func AttrBgColor8Bit(bg8BitColor int) BgColor8Bit {
	return BgColor8Bit(bg8BitColor)
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c BgColor8Bit) ANSI() string {
	return "48;5;" + strconv.Itoa(int(c))
}

// FgColor24Bit is a 24 bit foreground color attribute.
type FgColor24Bit struct {
	red   int
	green int
	blue  int
}

// AttrFgColor24Bit returns a new 24 bit foreground color attribute.
func AttrFgColor24Bit(red, green, blue int) *FgColor24Bit {
	return &FgColor24Bit{red, green, blue}
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c *FgColor24Bit) ANSI() string {
	colorStr := strconv.Itoa(c.red) + ";"
	colorStr += strconv.Itoa(c.green) + ";"
	colorStr += strconv.Itoa(c.blue)

	return "38;2;" + colorStr
}

// BgColor24Bit is a 24 bit background color attribute.
type BgColor24Bit struct {
	red   int
	green int
	blue  int
}

// AttrBgColor24Bit returns a new 24 bit background color attribute.
func AttrBgColor24Bit(red, green, blue int) *BgColor24Bit {
	return &BgColor24Bit{red, green, blue}
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c *BgColor24Bit) ANSI() string {
	colorStr := strconv.Itoa(c.red) + ";"
	colorStr += strconv.Itoa(c.green) + ";"
	colorStr += strconv.Itoa(c.blue)

	return "48;2;" + colorStr
}
