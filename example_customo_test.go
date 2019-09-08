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
		customo.AttrFgColor4BitBlack,
		customo.AttrBgColor4BitBrightCyan,
	)

	fmt.Println(str)
}
