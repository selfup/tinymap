package tinymap

import (
	"testing"
)

func Test_tinyByteMap_Get(t *testing.T) {
	tinyByteMap := new(TinyByteMap)

	_, found := tinyByteMap.Get(1)

	if found != false {
		t.Errorf("failed to return false on the second result")
	}
}

func Benchmark_TinyByteMap_Get_Single_Lower_Bound(b *testing.B) {
	tinyByteMap := new(TinyByteMap)

	tinyByteMap.Set(1, []byte("bar"))

	for n := 0; n < b.N; n++ {
		tinyByteMap.Get(1)
	}
}

func Benchmark_TinyByteMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	tinyByteMap := new(TinyByteMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		tinyByteMap.Set(i, []byte(string(i)))
	}

	for n := 0; n < b.N; n++ {
		tinyByteMap.Get(upperBound)
	}
}
