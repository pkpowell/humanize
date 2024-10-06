package hubytes

import (
	"fmt"
	"strconv"

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
	// Binary - 1024 bytes in a kilobyte
	IEC ByteUnit = iota

	// Decimal - 1000 bytes in a kilobyte
	SI
)

// Default options
var Options = &ByteOptions{
	Unit:           IEC,  // Binary
	MaxDecimals:    One,  // 1 decimal place
	ShowByteLetter: true, // Show trailing 'b'
}

func (o *ByteOptions) ByteLetter() string {
	if Options.ShowByteLetter {
		return "b"
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
var iec, _ = decimal.NewFromInt64(1024, 0, 0)

// Decimal
var si, _ = decimal.NewFromInt64(1000, 0, 0)

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
	value, err := decimal.NewFromInt64(int64(s), 0, 0)
	if err != nil {
		fmt.Printf("decimal.NewFromInt64 error %s", err)
		return ""
	}

	var p string
	for _, p = range Options.Unit.Prefix() {
		if value.Less(Options.Unit.Divisor()) {
			break
		}
		value, err = value.Quo(Options.Unit.Divisor())
		if err != nil {
			fmt.Printf("decimal.NewFromInt64 error %s", err)
			return ""
		}
	}
	f, _ := value.Ceil(int(Options.MaxDecimals)).Float64()
	return fmt.Sprintf("%s%s%s", strconv.FormatFloat(f, 'f', -1, 32), p, Options.ByteLetter())
	// return fmt.Sprintf("%.1f%s%s", value, p, Options.ByteLetter())
}
