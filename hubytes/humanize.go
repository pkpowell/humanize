package hubytes

import (
	"fmt"

	"github.com/govalues/decimal"
)

type Byter int64

// 1000 or 1024
type Unit int

// decimal places - 1, 2, 3...
type DecimalCount int

const (
	One   DecimalCount = 1
	Two   DecimalCount = 2
	Three DecimalCount = 3
	Four  DecimalCount = 4
	Five  DecimalCount = 5
)

type ByteOptions struct {
	Unit           Unit         // IEC or SI
	MaxDecimals    DecimalCount // decimal places without trailing zeros
	ShowByteLetter bool         // Show trailing 'b' (bytes)
	Full           bool         // Show full name
}

type ByteUnit struct {
	Letter string
	Name   string
}

const (
	// Binary - 1024 bytes in a kilobyte, i.e. 1KiB
	IEC Unit = 1024

	// Decimal - 1000 bytes in a kilobyte, i.e. 1KB
	SI Unit = 1000
)

// Default options
var Options = &ByteOptions{
	Unit:           IEC,   // binary - 1024 based
	MaxDecimals:    One,   // 1 decimal place
	ShowByteLetter: true,  // show trailing 'b'
	Full:           false, // print Full name
}

func (o *ByteOptions) ByteLetter() string {
	if Options.ShowByteLetter {
		return "B"
	}
	return ""
}

var Prefix = map[Unit][]ByteUnit{
	// Binary - 1024 bytes in a kilobyte
	IEC: {
		{Letter: "", Name: "byte"},
		{Letter: "Ki", Name: "kibibyte"},
		{Letter: "Mi", Name: "mebibyte"},
		{Letter: "Gi", Name: "gibibyte"},
		{Letter: "Ti", Name: "tebibyte"},
		{Letter: "Pi", Name: "pebibyte"},
		{Letter: "Ei", Name: "exbibyte"},
		{Letter: "Zi", Name: "zebibyte"},
		{Letter: "Yi", Name: "yobibyte"},
	},

	// Decimal - 1000 bytes in a kilobyte
	SI: {
		{Letter: "", Name: "byte"},
		{Letter: "k", Name: "kilobyte"},
		{Letter: "M", Name: "megabyte"},
		{Letter: "G", Name: "gigabyte"},
		{Letter: "T", Name: "terabyte"},
		{Letter: "P", Name: "petabyte"},
		{Letter: "E", Name: "exabyte"},
		{Letter: "Z", Name: "zettabyte"},
		{Letter: "Y", Name: "yottabyte"},
		{Letter: "R", Name: "ronnabyte"},
		{Letter: "Q", Name: "quettabyte"},
	},
}

// Binary
var iec, _ = decimal.New(int64(IEC), 0)

// Decimal
var si, _ = decimal.New(int64(SI), 0)

// value container
var value decimal.Decimal
var p ByteUnit

var Divisor = map[Unit]decimal.Decimal{
	// Binary
	IEC: iec,
	// Decimal
	SI: si,
}

func (b *Unit) Prefix() []ByteUnit {
	return Prefix[*b]
}

func (b *Unit) Divisor() decimal.Decimal {
	return Divisor[*b]
}

func (s Byter) String() string {
	value, _ = decimal.New(int64(s), 0)

	for _, p = range Options.Unit.Prefix() {
		if value.Less(Options.Unit.Divisor()) {
			break
		}
		value, _ = value.Quo(Options.Unit.Divisor())
	}

	value = value.Trunc(int(Options.MaxDecimals)).Trim(0)

	if Options.Full {
		if value.Equal(decimal.One) {
			return fmt.Sprintf("%s %s", value, p.Name)
		}
		return fmt.Sprintf("%s %s", value, p.Name+"s")
	}

	return fmt.Sprintf("%s%s%s", value, p.Letter, Options.ByteLetter())
}
