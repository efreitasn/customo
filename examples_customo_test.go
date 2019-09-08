package customo_test

import (
	"fmt"

	"github.com/efreitasn/customo"
)

func ExampleFormat_colors4Bits() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFgColor4BitsBlack,
		customo.AttrBgColor4BitsBrightCyan,
	)

	fmt.Println(str)
}

func ExampleFormat_colors8Bits() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFgColor8Bits(232),
		customo.AttrBgColor8Bits(255),
	)

	fmt.Println(str)
}

func ExampleFormat_colors24Bits() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFgColor24Bits(0, 255, 0),
		customo.AttrBgColor24Bits(255, 0, 0),
	)

	fmt.Println(str)
}
