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

func Benchmark_ByteMap_Get_Single_Lower_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	byteMap.Set([]byte("1"), []byte("bar"))

	for n := 0; n < b.N; n++ {
		byteMap.Get([]byte("1"))
	}
}

func Benchmark_ByteMap_Get_Max_Size_Upper_Bound(b *testing.B) {
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
