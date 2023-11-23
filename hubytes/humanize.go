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

var Options = &ByteOptions{
	Unit:           SI,
	ShowByteLetter: true,
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
