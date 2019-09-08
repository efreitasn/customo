package customo

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	normalStr := "a"
	tests := []struct {
		title string
		attrs []int
		str   string
	}{
		{
			"Bold, crossed, bright, black fg and cyan bg",
			[]int{
				AttrBold,
				AttrStrikethrough,
				AttrFg4BitBlack,
				AttrBg4BitBrightCyan,
			},
			fmt.Sprintf("\x1b[1;9;30;106m%v\x1b[0m", normalStr),
		},
		{
			"Italic, blink, bright black fg and blue bg",
			[]int{
				AttrItalic,
				AttrBlink,
				AttrFg4BitBrightBlack,
				AttrBg4BitBlue,
			},
			fmt.Sprintf("\x1b[3;5;90;44m%v\x1b[0m", normalStr),
		},
		{
			"Bold, blink, yellow fg and white bg",
			[]int{
				AttrBold,
				AttrBlink,
				AttrFg4BitYellow,
				AttrBg4BitWhite,
			},
			fmt.Sprintf("\x1b[1;5;33;47m%v\x1b[0m", normalStr),
		},
		{
			"No attrs",
			[]int{},
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
