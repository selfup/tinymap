package tinymap

import (
	"testing"
)

func TestTinyStrMap_Get(t *testing.T) {
	tinyStrMap := new(TinyStrMap)

	result, found := tinyStrMap.Get("a")

	if result != "" {
		t.Errorf("failed to return an empty string on the first result")
	}

	if found != false {
		t.Errorf("failed to return false on the second result")
	}
}

func Benchmark_Get_Single(b *testing.B) {
	tinyStrMap := new(TinyStrMap)

	tinyStrMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		tinyStrMap.Get("a")
	}
}

func Benchmark_Get_Multiple(b *testing.B) {
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
