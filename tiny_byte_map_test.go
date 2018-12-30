package tinymap

import (
	"testing"
)

func Test_byteMap_Get(t *testing.T) {
	byteMap := new(ByteMap)

	_, err := byteMap.Get([]byte("1"))

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_ByteMap_Get_Lower_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	byteMap.Set([]byte("1"), []byte("bar"))

	for n := 0; n < b.N; n++ {
		byteMap.Get([]byte("1"))
	}
}

func Benchmark_ByteMap_Get_Expected_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	byteMap.Set([]byte("0"), []byte("8996"))
	byteMap.Set([]byte("0"), []byte("8997"))
	byteMap.Set([]byte("0"), []byte("8998"))
	byteMap.Set([]byte("0"), []byte("8999"))
	byteMap.Set([]byte("42"), []byte("9000"))

	for n := 0; n < b.N; n++ {
		byteMap.Get([]byte("42"))
	}
}

func Benchmark_ByteMap_Get_Upper_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	upperBound := 100
	upperBoundStr := string(upperBound)
	upperBoundByte := []byte(upperBoundStr)

	for i := 0; i < upperBound; i++ {
		byteMap.Set([]byte(string(i)), upperBoundByte)
	}

	byteMap.Set(upperBoundByte, upperBoundByte)

	for n := 0; n < b.N; n++ {
		byteMap.Get(upperBoundByte)
	}
}
