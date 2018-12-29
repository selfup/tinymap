package tinymap

import (
	"testing"
)

func Test_byteMap_Get(t *testing.T) {
	byteMap := new(ByteMap)

	_, err := byteMap.Get(1)

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_ByteMap_Get_Single_Lower_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	byteMap.Set(1, []byte("bar"))

	for n := 0; n < b.N; n++ {
		byteMap.Get(1)
	}
}

func Benchmark_ByteMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		byteMap.Set(i, []byte(string(i)))
	}

	for n := 0; n < b.N; n++ {
		byteMap.Get(upperBound)
	}
}
