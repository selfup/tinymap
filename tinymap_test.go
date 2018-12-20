package tinymap

import "testing"

func Benchmark_Get(b *testing.B) {
	tinyStrMap := new(TinyStrMap)

	tinyStrMap.Set("a", "b")

	for n := 0; n < b.N; n++ {
		tinyStrMap.Get("a")
	}
}
