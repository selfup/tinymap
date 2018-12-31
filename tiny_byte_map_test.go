package tinymap

import (
	"bytes"
	"testing"
)

func Test_byteMap_Get(t *testing.T) {
	byteMap := new(ByteMap)

	_, err := byteMap.Get([]byte("1"))

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Test_ByteMap_Set(t *testing.T) {
	byteMap := new(ByteMap)

	key := []byte("42")
	val := []byte("9000")

	byteMap.Set(key, val)

	result, _ := byteMap.Get(key)

	if !bytes.Equal(result, val) {
		t.Errorf("failed to return correct Val from StrTuple")
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

func Benchmark_ByteMap_Set_Upper_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	for i := 0; i < b.N; i++ {
		val := []byte(string(i))

		if i < 100 {
			byteMap.Set(val, val)
		} else {
			byteMap.Set([]byte("1"), val)
		}
	}
}

func Benchmark_ByteMap_Delete_Upper_Bound(b *testing.B) {
	byteMap := new(ByteMap)

	for i := 0; i < 100; i++ {
		val := []byte(string(i))

		byteMap.Set(val, val)
	}

	for i := 0; i < b.N; i++ {
		val := []byte(string(i))

		if i < 100 {
			byteMap.Delete(val)
		} else {
			byteMap.Delete(val)
		}
	}
}
