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
			FBold | FCrossed,
			fmt.Sprintf("\x1b[1;9m%v\x1b[0m", normalStr),
		},
		{
			"Italic and blink",
			FItalic | FBlink,
			fmt.Sprintf("\x1b[3;5m%v\x1b[0m", normalStr),
		},
		{
			"Without flags",
			0,
			normalStr,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			c := New()

			c.AddFlags(test.flags)

			str := c.FormatString(normalStr)

			if str != test.str {
				t.Errorf("got %v, want %v", escapeStr(str), escapeStr(test.str))
			}
		})
	}
}
