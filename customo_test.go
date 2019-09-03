package customo

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title string
		flags int
		str   string
	}{
		{
			"Bold and crossed",
			AttrBold | AttrCrossed,
			fmt.Sprintf("\x1b[1;9m%v\x1b[0m", normalStr),
		},
		{
			"Italic and blink",
			AttrItalic | AttrBlink,
			fmt.Sprintf("\x1b[3;5m%v\x1b[0m", normalStr),
		},
		{
			"No flags",
			0,
			normalStr,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			c := New()

			c.AddAttrs(test.flags)

			str := c.FormatString(normalStr)

			if str != test.str {
				t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
			}
		})
	}
}

func TestBgColor(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title   string
		flags   int
		bgColor int
		str     string
	}{
		{
			"Bold and crossed with blue bg",
			AttrBold | AttrCrossed,
			Bg8BitBlue,
			fmt.Sprintf("\x1b[1;9;%vm%v\x1b[0m", Bg8BitBlue, normalStr),
		},
		{
			"Italic and blink with yellow bg",
			AttrItalic | AttrBlink,
			Bg8BitYellow,
			fmt.Sprintf("\x1b[3;5;%vm%v\x1b[0m", Bg8BitYellow, normalStr),
		},
		{
			"No flags with red bg",
			0,
			Bg8BitRed,
			fmt.Sprintf("\x1b[%vm%v\x1b[0m", Bg8BitRed, normalStr),
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			c := New()

			c.AddAttrs(test.flags)
			c.SetBg8BitAttr(test.bgColor)

			str := c.FormatString(normalStr)

			if str != test.str {
				t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
			}
		})
	}
}

func TestFgColor(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title   string
		flags   int
		fgColor int
		str     string
	}{
		{
			"Bold and crossed with blue fg",
			AttrBold | AttrCrossed,
			Fg8BitBlue,
			fmt.Sprintf("\x1b[1;9;%vm%v\x1b[0m", Fg8BitBlue, normalStr),
		},
		{
			"Italic and blink with yellow fg",
			AttrItalic | AttrBlink,
			Fg8BitYellow,
			fmt.Sprintf("\x1b[3;5;%vm%v\x1b[0m", Fg8BitYellow, normalStr),
		},
		{
			"No flags with red fg",
			0,
			Fg8BitRed,
			fmt.Sprintf("\x1b[%vm%v\x1b[0m", Fg8BitRed, normalStr),
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			c := New()

			c.AddAttrs(test.flags)
			c.SetFg8BitAttr(test.fgColor)

			str := c.FormatString(normalStr)

			if str != test.str {
				t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
			}
		})
	}
}
