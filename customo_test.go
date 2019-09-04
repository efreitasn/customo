package customo

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title    string
		attrs    int
		fgColor  int
		bgColor  int
		str      string
		errorNil bool
	}{
		{
			"Bold, crossed, bright, black fg and cyan bg",
			AttrBold | AttrCrossed,
			Fg4BitBlack,
			Bg4BitBrightCyan,
			fmt.Sprintf("\x1b[1;9;30;106m%v\x1b[0m", normalStr),
			true,
		},
		{
			"Italic, blink, bright black fg and blue bg",
			AttrItalic | AttrBlink,
			Fg4BitBrightBlack,
			Bg4BitBlue,
			fmt.Sprintf("\x1b[3;5;90;44m%v\x1b[0m", normalStr),
			true,
		},
		{
			"Bold, blink, yellow fg and white bg",
			AttrBold | AttrBlink,
			Fg4BitYellow,
			Bg4BitWhite,
			fmt.Sprintf("\x1b[1;5;33;47m%v\x1b[0m", normalStr),
			true,
		},
		{
			"Foreground color out of range",
			0,
			93811,
			0,
			"",
			false,
		},
		{
			"Background color out of range",
			0,
			0,
			93811,
			"",
			false,
		},
		{
			"No attrs",
			0,
			0,
			0,
			normalStr,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			c, err := New(test.fgColor, test.bgColor, test.attrs)

			if err == nil != test.errorNil {
				if test.errorNil {
					t.Errorf("got %v, want nil", err)
				} else {
					t.Errorf("got %v, want non-nil value", err)
				}
			} else if err == nil {
				str := c.FormatString(normalStr)

				if str != test.str {
					t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
				}
			}
		})
	}
}
