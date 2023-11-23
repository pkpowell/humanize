package hubytes

import (
	"fmt"
	"math"
)

// type BytesIEC int64
// type BytesSI int64
type Byter int64

type ByteUnit int

type ByteOptions struct {
	Unit           ByteUnit
	ShowByteLetter bool
}

const (
	IEC ByteUnit = iota
	SI
)

var Options = &ByteOptions{}

func NewOptions() *ByteOptions {
	return &ByteOptions{
		Unit:           IEC,
		ShowByteLetter: false,
	}
}

func (s Byter) String() string {
	value := float64(s)
	var div float64
	var prefix []string
	b := ""

	if Options.Unit == SI {
		prefix = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}
		div = 1000.0
	} else {
		prefix = []string{"", "K", "M", "G", "T", "P", "E"}
		div = 1024.0
	}

	if Options.ShowByteLetter {
		b = "B"
	}

	p := ""
	for _, p = range prefix {
		if math.Abs(value) < div {
			break
		}
		value /= div
	}
	return fmt.Sprintf("%.1f%s%s", value, p, b)
}

// func (s BytesIEC) String() string {
// 	value := float64(s)
// 	var div float64
// 	if Options.Unit == SI {
// 		div = 1000.0
// 	} else {
// 		div = 1024.0
// 	}
// 	prefix := ""
// 	for _, prefix = range iecPrefixes {
// 		if math.Abs(value) < div {
// 			break
// 		}
// 		value /= div
// 	}
// 	return fmt.Sprintf("%.1f%s", value, prefix)
// }

// func (s BytesSI) String() string {
// 	value := float64(s)
// 	div := 1024.0
// 	prefix := ""
// 	for _, prefix = range siPrefixes {
// 		if math.Abs(value) < div {
// 			break
// 		}
// 		value /= div
// 	}
// 	return fmt.Sprintf("%.1f%s", value, prefix)
// }
