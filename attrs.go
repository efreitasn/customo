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
	AttrBold                      = SimpleAttr(1)
	AttrItalic                    = SimpleAttr(3)
	AttrUnderline                 = SimpleAttr(4)
	AttrBlink                     = SimpleAttr(5)
	AttrStrikethrough             = SimpleAttr(9)
	AttrFgColor4BitsBlack         = SimpleAttr(30)
	AttrFgColor4BitsRed           = SimpleAttr(31)
	AttrFgColor4BitsGreen         = SimpleAttr(32)
	AttrFgColor4BitsYellow        = SimpleAttr(33)
	AttrFgColor4BitsBlue          = SimpleAttr(34)
	AttrFgColor4BitsMagenta       = SimpleAttr(35)
	AttrFgColor4BitsCyan          = SimpleAttr(36)
	AttrFgColor4BitsWhite         = SimpleAttr(37)
	AttrBgColor4BitsBlack         = SimpleAttr(40)
	AttrBgColor4BitsRed           = SimpleAttr(41)
	AttrBgColor4BitsGreen         = SimpleAttr(42)
	AttrBgColor4BitsYellow        = SimpleAttr(43)
	AttrBgColor4BitsBlue          = SimpleAttr(44)
	AttrBgColor4BitsMagenta       = SimpleAttr(45)
	AttrBgColor4BitsCyan          = SimpleAttr(46)
	AttrBgColor4BitsWhite         = SimpleAttr(47)
	AttrOverline                  = SimpleAttr(53)
	AttrFgColor4BitsBrightBlack   = SimpleAttr(90)
	AttrFgColor4BitsBrightRed     = SimpleAttr(91)
	AttrFgColor4BitsBrightGreen   = SimpleAttr(92)
	AttrFgColor4BitsBrightYellow  = SimpleAttr(93)
	AttrFgColor4BitsBrightBlue    = SimpleAttr(94)
	AttrFgColor4BitsBrightMagenta = SimpleAttr(95)
	AttrFgColor4BitsBrightCyan    = SimpleAttr(96)
	AttrFgColor4BitsBrightWhite   = SimpleAttr(97)
	AttrBgColor4BitsBrightBlack   = SimpleAttr(100)
	AttrBgColor4BitsBrightRed     = SimpleAttr(101)
	AttrBgColor4BitsBrightGreen   = SimpleAttr(102)
	AttrBgColor4BitsBrightYellow  = SimpleAttr(103)
	AttrBgColor4BitsBrightBlue    = SimpleAttr(104)
	AttrBgColor4BitsBrightMagenta = SimpleAttr(105)
	AttrBgColor4BitsBrightCyan    = SimpleAttr(106)
	AttrBgColor4BitsBrightWhite   = SimpleAttr(107)
)

// FgColor8Bits is an 8 bit foreground color attribute.
type FgColor8Bits int

// AttrFgColor8Bits returns a new 8 bit foreground color attribute.
func AttrFgColor8Bits(color int) FgColor8Bits {
	return FgColor8Bits(color)
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c FgColor8Bits) ANSI() string {
	return "38;5;" + strconv.Itoa(int(c))
}

// BgColor8Bits is an 8 bit background color attribute.
type BgColor8Bits int

// AttrBgColor8Bits returns a new 8 bit background color attribute.
func AttrBgColor8Bits(bg8BitsColor int) BgColor8Bits {
	return BgColor8Bits(bg8BitsColor)
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c BgColor8Bits) ANSI() string {
	return "48;5;" + strconv.Itoa(int(c))
}

// FgColor24Bits is a 24 bit foreground color attribute.
type FgColor24Bits struct {
	red   int
	green int
	blue  int
}

// AttrFgColor24Bits returns a new 24 bit foreground color attribute.
func AttrFgColor24Bits(red, green, blue int) *FgColor24Bits {
	return &FgColor24Bits{red, green, blue}
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c *FgColor24Bits) ANSI() string {
	colorStr := strconv.Itoa(c.red) + ";"
	colorStr += strconv.Itoa(c.green) + ";"
	colorStr += strconv.Itoa(c.blue)

	return "38;2;" + colorStr
}

// BgColor24Bits is a 24 bit background color attribute.
type BgColor24Bits struct {
	red   int
	green int
	blue  int
}

// AttrBgColor24Bits returns a new 24 bit background color attribute.
func AttrBgColor24Bits(red, green, blue int) *BgColor24Bits {
	return &BgColor24Bits{red, green, blue}
}

// ANSI returns a representation to be used as an ANSI escape code.
func (c *BgColor24Bits) ANSI() string {
	colorStr := strconv.Itoa(c.red) + ";"
	colorStr += strconv.Itoa(c.green) + ";"
	colorStr += strconv.Itoa(c.blue)

	return "48;2;" + colorStr
}
