package units

import (
	"github.com/govalues/decimal"
)

type Int int64
type Byte []byte
type String string

// 1000 or 1024
type Unit int

// decimal places - 1, 2, 3...
type DecimalCount int

const (
	Zero  DecimalCount = 0
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
	Space          bool         // Space between value and unit
}

type byteUnit struct {
	Letter string
	Name   string
}

const (
	// Binary - 1024 bytes in a kilobyte, i.e. 1KiB
	IEC Unit = 1024

	// Decimal - 1000 bytes in a kilobyte, i.e. 1KB
	SI Unit = 1000
	// spc string = " "
)

// Default options
var Options = &ByteOptions{
	Unit:           IEC,   // binary - 1024 based
	MaxDecimals:    One,   // 1 decimal place
	ShowByteLetter: true,  // show trailing 'b'
	Full:           false, // print Full name
	Space:          false, // space between value and unit
}

// Binary
var iec, _ = decimal.New(int64(IEC), 0)

// Decimal
var si, _ = decimal.New(int64(SI), 0)

// value container
var value decimal.Decimal

var p byteUnit

var unitSpace string

var divisor = map[Unit]decimal.Decimal{
	// Binary
	IEC: iec,

	// Decimal
	SI: si,
}

func (o *ByteOptions) byteLetter() string {
	if Options.ShowByteLetter {
		return "B"
	}
	return ""
}

var prefix = map[Unit][]byteUnit{
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

func (b *Unit) prefix() []byteUnit {
	return prefix[*b]
}

func (b *Unit) divisor() decimal.Decimal {
	return divisor[*b]
}

func (s String) String() string {
	return Int(len(s)).String()
}

func (s Byte) String() string {
	return Int(len(s)).String()
}

func (s Int) String() string {
	if Options.Space {
		unitSpace = " "
	}
	value, _ = decimal.New(int64(s), 0)

	for _, p = range Options.Unit.prefix() {
		if value.Less(Options.Unit.divisor()) {
			break
		}
		value, _ = value.Quo(Options.Unit.divisor())
	}

	value = value.Trunc(int(Options.MaxDecimals)).Trim(0)

	if Options.Full {
		if value.Equal(decimal.One) {
			return value.String() + unitSpace + p.Name
		}

		return value.String() + unitSpace + p.Name + "s"
	}

	return value.String() + unitSpace + p.Letter + Options.byteLetter()
}
