package customo_test

import (
	"fmt"

	"github.com/efreitasn/customo"
)

func ExampleFormat() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrStrikethrough,
		customo.AttrFg4BitBlack,
		customo.AttrBg4BitBrightCyan,
	)

	fmt.Println(str)
}
