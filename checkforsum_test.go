package checkforsum

import (
	"math/rand"
	"testing"
)

func rndIntSlc(len, min, max int) []int {
	res := make([]int, len)
	for i := range res {
		res[i] = rand.Intn(max-min+1) + min
	}
	return res
}

var s = rndIntSlc(1000, 1, 10000)

func Benchmark_findInUnsortedArray1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findInUnsortedArray1(s, 900)
	}
}

func Benchmark_findInUnsortedArray2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findInUnsortedArray2(s, 900)
	}
}
