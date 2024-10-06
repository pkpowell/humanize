package hubytes

import (
	"fmt"

	"github.com/govalues/decimal"
)

type Byter int64

type ByteUnit int
type DecimalCount int

const (
	One   DecimalCount = 1
	Two   DecimalCount = 2
	Three DecimalCount = 3
	Four  DecimalCount = 4
	Five  DecimalCount = 5
)

type ByteOptions struct {
	Unit           ByteUnit     // IEC or SI
	MaxDecimals    DecimalCount // decimal places without trailing zeros
	ShowByteLetter bool         // Show trailing 'b' (bytes)
}

const (
	// Binary - 1024 bytes in a kilobyte, i.e. 1KiB
	IEC ByteUnit = 1024

	// Decimal - 1000 bytes in a kilobyte, i.e. 1KB
	SI ByteUnit = 1000
)

// Default options
var Options = &ByteOptions{
	Unit:           IEC,  // Binary
	MaxDecimals:    One,  // 1 decimal place
	ShowByteLetter: true, // Show trailing 'b'
}

func (o *ByteOptions) ByteLetter() string {
	if Options.ShowByteLetter {
		return "B"
	}
	return ""
}

var Prefix = map[ByteUnit][]string{
	// Binary - 1024 bytes in a kilobyte
	IEC: {"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"},

	// Decimal - 1000 bytes in a kilobyte
	SI: {"", "K", "M", "G", "T", "P", "E"},
}

// Binary
var iec, _ = decimal.New(int64(IEC), 0)

// Decimal
var si, _ = decimal.New(int64(SI), 0)

var Divisor = map[ByteUnit]decimal.Decimal{
	// Binary
	IEC: iec,
	// Decimal
	SI: si,
}

func (b *ByteUnit) Prefix() []string {
	return Prefix[*b]
}

func (b *ByteUnit) Divisor() decimal.Decimal {
	return Divisor[*b]
}

func (s Byter) String() string {
	value, _ := decimal.New(int64(s), 0)

	var p string
	for _, p = range Options.Unit.Prefix() {
		if value.Less(Options.Unit.Divisor()) {
			break
		}
		value, _ = value.Quo(Options.Unit.Divisor())
	}
	return fmt.Sprintf("%s %s%s", value.Trunc(int(Options.MaxDecimals)).Trim(0), p, Options.ByteLetter())
}
