package main

import (
	"fmt"

	"github.com/pkpowell/humanize/hubytes"
)

func main() {
	hubytes.Options.Unit = hubytes.IEC
	hubytes.Options.ShowByteLetter = false

	fmt.Printf("h %s\n", hubytes.Byter(123456))
	hubytes.Options.Unit = hubytes.SI

	fmt.Printf("h %s\n", hubytes.Byter(123456))
	// fmt.Printf("IEC %s\n", humanize.BytesIEC(123456))
	// fmt.Printf("SI %s\n", humanize.BytesSI(123456))
}
