package hubytes

import (
	"fmt"
	"math"
	"strconv"
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

func (o *ByteOptions) ByteLetter() string {
	if Options.ShowByteLetter {
		return "b"
	}
	return ""
}

var Prefix = map[ByteUnit][]string{
	IEC: {"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"},
	SI:  {"", "K", "M", "G", "T", "P", "E"},
}

var Divisor = map[ByteUnit]float64{
	IEC: 1000.0,
	SI:  1024.0,
}

func (b *ByteUnit) Prefix() (prefix []string) {
	return Prefix[*b]
}

func (b *ByteUnit) Divisor() (div float64) {
	return Divisor[*b]
}

func (s Byter) String() string {
	value := float64(s) * 10
	var p string
	for _, p = range Options.Unit.Prefix() {
		if math.Abs(value) < Options.Unit.Divisor() {
			break
		}
		value = math.Round(value / Options.Unit.Divisor())
	}

	return fmt.Sprintf("%s%s%s", strconv.FormatFloat(value/10, 'f', -1, 32), p, Options.ByteLetter())
}
