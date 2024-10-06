package hubytes

import (
	"testing"
)

var byteValues = []int64{
	23,
	1000,
	1001,
	1024,
	1321,
	2000,
	2001,
	2048,
	6000,
	6001,
	21325,
	243202,
	1839405,
	18394053,
	1000 * 1000 * 1000,
	1024 * 1024 * 1024,
}

// func TestHumanize(t *testing.T) {
// 	for _, v := range byteValues {
// 		fmt.Println(Byter(v))
// 		t.Log(v, Byter(v))
// 	}
// }

func BenchmarkHumanizeIEC(b *testing.B) {
	Options.Unit = IEC
	Options.MaxDecimals = One
	for range b.N {
		for _, v := range byteValues {
			b.Log(v, Byter(v))
			// _ = Byter(v)
		}
	}
}

func BenchmarkHumanizeSI(b *testing.B) {
	Options.Unit = SI
	Options.MaxDecimals = Two
	// b.Log("Decimals: ", Options.MaxDecimals)
	for range b.N {
		for _, v := range byteValues {
			// _ = Byter(v)
			b.Log(v, Byter(v))
		}
	}
}
