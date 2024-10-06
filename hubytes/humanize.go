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
	Unit           ByteUnit
	MaxDecimals    DecimalCount
	ShowByteLetter bool
}

const (
	IEC ByteUnit = iota
	SI
)

// Default options
var Options = &ByteOptions{
	Unit:           SI,
	ShowByteLetter: true,
	MaxDecimals:    One,
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

var iec, _ = decimal.NewFromInt64(1000, 0, 0)
var si, _ = decimal.NewFromInt64(1024, 0, 0)

var Divisor = map[ByteUnit]decimal.Decimal{
	IEC: iec,
	SI:  si,
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
