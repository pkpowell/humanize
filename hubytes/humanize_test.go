package hubytes

import (
	"fmt"
	"testing"
)

var byteValues = []int64{
	1024,
	1011 * 1024,
	1022 * 1024 * 1024,
	1012 * 1024 * 1024 * 1024,
	1013 * 1024 * 1024 * 1024 * 1024,
	1017 * 1024 * 1024 * 1024 * 1024 * 1024,
	1007 * 1024 * 1024 * 1024 * 1022,
}

func TestHumanize(t *testing.T) {
	for _, v := range byteValues {
		fmt.Println(Byter(v))
	}
}

func TestHumanizeIEC(t *testing.T) {
	for _, v := range byteValues {
		fmt.Println(Byter(v).IEC())
	}
}
