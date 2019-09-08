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
		customo.AttrFgColor4BitBlack,
		customo.AttrBgColor4BitBrightCyan,
	)

	fmt.Println(str)
}

func ExampleFormat_colors8Bits() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFgColor8Bit(232),
		customo.AttrBgColor8Bit(255),
	)

	fmt.Println(str)
}

func ExampleFormat_colors24Bits() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFgColor24Bit(0, 255, 0),
		customo.AttrBgColor24Bit(255, 0, 0),
	)

	fmt.Println(str)
}
