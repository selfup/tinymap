package tinymap

import "testing"

func Benchmark_Get(b *testing.B) {
	tinyStrMap := new(TinyStrMap)

	tinyStrMap.Set("z", "_")
	tinyStrMap.Set("f", "_")
	tinyStrMap.Set("v", "_")
	tinyStrMap.Set("t", "_")
	tinyStrMap.Set("a", "_")
	tinyStrMap.Set("j", "_")
	tinyStrMap.Set("x", "_")

	for n := 0; n < b.N; n++ {
		tinyStrMap.Get("a")
	}
}
