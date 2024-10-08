package hubytes

import (
	"testing"
)

var byteValues = []int64{
	23,
	1000,
	1024,
	1321,
	2000,
	2048,
	1000 * 1000 * 1000,
	1024 * 1024 * 1024,
	6000,
	6001,
	21325,
	243202,
	1839405,
	18394053,
}

func BenchmarkHumanizeIEC(b *testing.B) {
	Options.Unit = IEC
	Options.MaxDecimals = One
	Options.Full = false
	for range b.N {
		for _, v := range byteValues {
			b.Log(v, Byter(v))
			// _ = Byter(v)
		}
	}
}

var container string

func BenchmarkHumanizeSI(b *testing.B) {
	Options.Unit = SI
	Options.MaxDecimals = Two
	Options.Full = false
	// b.Log("Decimals: ", Options.MaxDecimals)
	for range b.N {
		for _, v := range byteValues {
			container = Byter(v).String()
			// b.Log(v, Byter(v))
		}
	}
}
