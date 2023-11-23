package humanize

import (
	"fmt"
	"math"
)

type BytesIEC int64
type BytesSI int64

var iecPrefixes = []string{"", "K", "M", "G", "T", "P", "E"}
var siPrefixes = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}

func (s BytesIEC) String() string {
	value := float64(s)
	div := 1000.0
	prefix := ""
	for _, prefix = range iecPrefixes {
		if math.Abs(value) < div {
			break
		}
		value /= div
	}
	return fmt.Sprintf("%.1f%s", value, prefix)
}

func (s BytesSI) String() string {
	value := float64(s)
	div := 1024.0
	prefix := ""
	for _, prefix = range siPrefixes {
		if math.Abs(value) < div {
			break
		}
		value /= div
	}
	return fmt.Sprintf("%.1f%s", value, prefix)
}
