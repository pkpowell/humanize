package hubytes

import (
	"fmt"
	"testing"
)

var byteValues = []int64{
	23,
	1321,
	21325,
	243202,
	1839405,
	18394053,
	1000 * 1000 * 1000,
	1024 * 1024 * 1024,
}

func TestHumanize(t *testing.T) {
	for _, v := range byteValues {
		fmt.Println(Byter(v))
		t.Log(v, Byter(v))
	}
}

func BenchmarkHumanizeIEC(b *testing.B) {
	for _, v := range byteValues {
		fmt.Println(Byter(v))
		b.Log(v, Byter(v))
	}
}