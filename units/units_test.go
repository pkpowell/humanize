package units

import (
	"testing"
)

var byteValues = []int64{
	23,
	1000,
	// 1024,
	1321,
	// 2000,
	2048,
	1000 * 1000 * 1000,
	1024 * 1024 * 1024,
	6000,
	6001,
	// 21325,
	243202,
	1839405,
	18394053,
	18394053123456,
	1839405312345612345,
}

func BenchmarkHumanizeIEC(b *testing.B) {
	Options.Unit = IEC
	Options.MaxDecimals = One
	Options.Full = false
	for range b.N {
		for _, v := range byteValues {
			b.Log(v, Int(v))
			// _ = Byter(v)
		}
	}
}

var container string

func BenchmarkHumanizeSI(b *testing.B) {
	Options.Unit = SI
	Options.MaxDecimals = Two
	Options.Full = false
	Options.Space = true
	// b.Log("Decimals: ", Options.MaxDecimals)
	for range b.N {
		for _, v := range byteValues {
			container = Int(v).String()
			// b.Log(v, Byter(v))
		}
	}
}

func TestHumanize(t *testing.T) {
	Options.Unit = IEC
	Options.Space = false
	Options.MaxDecimals = One
	Options.Full = false
	Options.ShowByteLetter = true
	for _, v := range byteValues {
		t.Log(v, Int(v))
	}
}

func TestHumanizeInt(t *testing.T) {
	Options.Unit = IEC
	Options.Space = false
	Options.MaxDecimals = One
	Options.Full = false
	Options.ShowByteLetter = true
	for _, v := range byteValues {
		t.Log(v, Int(v))
	}
}

func TestHumanizeStr(t *testing.T) {
	var b = []byte{}
	Options.Unit = IEC
	Options.Space = false
	Options.MaxDecimals = One
	Options.Full = false
	Options.ShowByteLetter = true
	for v := range 10 {
		b = make([]byte, v*2048*2048)
		t.Log(v, String(b))
	}
}

func TestHumanizeByte(t *testing.T) {
	var b = []byte{}
	Options.Unit = IEC
	Options.Space = false
	Options.MaxDecimals = One
	Options.Full = false
	Options.ShowByteLetter = true
	for v := range 10 {
		b = make([]byte, v*2048*2048)
		t.Log(v, Byte(b))
	}
}
