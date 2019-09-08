# customo
[![GoDoc](https://godoc.org/github.com/efreitasn/customo?status.svg)](https://godoc.org/github.com/efreitasn/customo)  
customo helps when creating strings using ANSI escape codes attributes, such as italic, bold and custom foreground and background colors.

## Installation
```shell
go get github.com/efreitasn/customo
```

## Examples

### Bold, underlined string with white foreground and green background 4 bit colors.

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

### Bold string with gray foreground and white background 8 bit colors.

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
		customo.AttrFgColor8Bit(236),
		customo.AttrBgColor8Bit(255),
	)

	fmt.Println(str)
}
```

### Italic string with blue foreground and red background 24 bit colors.

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
		customo.AttrFgColor24Bit(0, 0, 255),
		customo.AttrBgColor24Bit(255, 0, 0),
	)

	fmt.Println(str)
}
```