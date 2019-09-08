package customo

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title string
		attrs []Attr
		str   string
	}{
		{
			"Bold, crossed, bright, black fg and cyan bg",
			[]Attr{
				AttrBold,
				AttrStrikethrough,
				AttrFgColor4BitBlack,
				AttrBgColor4BitBrightCyan,
			},
			fmt.Sprintf("\x1b[1;9;30;106m%v\x1b[0m", normalStr),
		},
		{
			"Italic, blink, bright black fg and blue bg",
			[]Attr{
				AttrItalic,
				AttrBlink,
				AttrFgColor4BitBrightBlack,
				AttrBgColor4BitBlue,
			},
			fmt.Sprintf("\x1b[3;5;90;44m%v\x1b[0m", normalStr),
		},
		{
			"Bold, blink, yellow fg and white bg",
			[]Attr{
				AttrBold,
				AttrBlink,
				AttrFgColor4BitYellow,
				AttrBgColor4BitWhite,
			},
			fmt.Sprintf("\x1b[1;5;33;47m%v\x1b[0m", normalStr),
		},
		{
			"Bold, 8 bit white fg and 8 bit green bg",
			[]Attr{
				AttrBold,
				AttrFgColor8Bit(255),
				AttrBgColor8Bit(35),
			},
			fmt.Sprintf("\x1b[1;38;5;255;48;5;35m%v\x1b[0m", normalStr),
		},
		{
			"Underline, 24 bit red fg and 24 bit blue bg",
			[]Attr{
				AttrUnderline,
				AttrFgColor24Bit(255, 0, 0),
				AttrBgColor24Bit(0, 0, 255),
			},
			fmt.Sprintf("\x1b[4;38;2;255;0;0;48;2;0;0;255m%v\x1b[0m", normalStr),
		},
		{
			"No attrs",
			[]Attr{},
			normalStr,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			str := Format(normalStr, test.attrs...)

			if str != test.str {
				t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
			}
		})
	}
}
