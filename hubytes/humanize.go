package hubytes

import (
	"fmt"
	"math"
	"strconv"
)

type Byter int64

type ByteUnit int

const (
	one   = 10
	two   = 100
	three = 1000
	four  = 10000
	five  = 100000
)

type ByteOptions struct {
	Unit           ByteUnit
	ShowByteLetter bool
	MaxDecimals    float64
}

const (
	IEC ByteUnit = iota
	SI
)

var Options = &ByteOptions{
	Unit:           SI,
	ShowByteLetter: true,
	MaxDecimals:    one,
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

func (b *ByteUnit) Prefix() []string {
	return Prefix[*b]
}

func (b *ByteUnit) Divisor() float64 {
	return Divisor[*b]
}

func (s Byter) String() string {
	value := float64(s)
	var prefix string
	for _, prefix = range Options.Unit.Prefix() {
		if value < Options.Unit.Divisor() {
			break
		}
		value /= Options.Unit.Divisor()
	}
	flt := math.Round(value*Options.MaxDecimals) / Options.MaxDecimals
	return fmt.Sprintf("%s%s%s", strconv.FormatFloat(flt, 'f', -1, 32), prefix, Options.ByteLetter())
}
