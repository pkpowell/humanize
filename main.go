package humanize

import (
	"fmt"
	"math"
)

type BytesIEC int64
type BytesSI int64
type Byter int64

type ByteUnit int

type ByteOptions struct {
	Format ByteUnit
	ShowB  bool
}

const (
	SI ByteUnit = iota
	IEC
)

var Options = &ByteOptions{
	Format: IEC,
	ShowB:  false,
}

var iecPrefixes = []string{"", "K", "M", "G", "T", "P", "E"}
var siPrefixes = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"}

func (s Byter) String() string {
	value := float64(s)
	var div float64
	var prefix []string
	if Options.Format == SI {
		prefix = siPrefixes
		div = 1000.0
	} else {
		div = 1024.0
	}
	p := ""
	for _, p = range prefix {
		if math.Abs(value) < div {
			break
		}
		value /= div
	}
	return fmt.Sprintf("%.1f%s", value, p)
}

func (s BytesIEC) String() string {
	value := float64(s)
	var div float64
	if Options.Format == SI {
		div = 1000.0
	} else {
		div = 1024.0
	}
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
