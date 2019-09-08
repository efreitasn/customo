# customo
[![GoDoc](https://godoc.org/github.com/efreitasn/customo?status.svg)](https://godoc.org/github.com/efreitasn/customo)  
customo helps when creating strings using ANSI escape codes attributes, such as italic, bold and custom foreground and background colors.

## Installation
```shell
go get github.com/efreitasn/customo
```

## Examples
Bold, underlined string with green background color and white foreground color.

```go
package main

import (
	"fmt"

	"github.com/efreitasn/customo"
)

func main() {
	str := customo.Format(
		"Custom string",
		customo.AttrBold,
		customo.AttrUnderline,
		customo.AttrBgColor4BitGreen,
		customo.AttrFgColor4BitWhite,
	)

	fmt.Println(str)
}
```