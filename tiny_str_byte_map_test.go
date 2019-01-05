package tinymap

import (
	"bytes"
	"testing"
)

func Test_StrByteMap_Get(t *testing.T) {
	strByteMap := new(StrByteMap)

	var defaultByte []byte
	result, err := strByteMap.Get("foo")

	if !bytes.Equal(result, defaultByte) {
		t.Errorf("failed to return []byte default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Test_StrByteMap_Set(t *testing.T) {
	strByteMap := new(StrByteMap)

	key := []byte("bar")

	strByteMap.Set("foo", key)

	result, _ := strByteMap.Get("foo")

	if !bytes.Equal(result, key) {
		t.Errorf("failed to return correct Val from StrByteTuple")
	}
}

func Benchmark_StrByteMap_Get_Lower_Bound(b *testing.B) {
	strByteMap := new(StrByteMap)

	key := []byte("_")

	strByteMap.Set("a", key)

	for n := 0; n < b.N; n++ {
		strByteMap.Get("a")
	}
}

func Benchmark_StrByteMap_Get_Expected_Bound(b *testing.B) {
	strByteMap := new(StrByteMap)

	zero := []byte("8996")
	one := []byte("8997")
	two := []byte("8998")
	three := []byte("8999")
	fortyTwo := []byte("9000")

	strByteMap.Set("0", zero)
	strByteMap.Set("1", one)
	strByteMap.Set("2", two)
	strByteMap.Set("3", three)
	strByteMap.Set("42", fortyTwo)

	for n := 0; n < b.N; n++ {
		strByteMap.Get("42")
	}
}

func Benchmark_StrByteMap_Get_Upper_Bound(b *testing.B) {
	strByteMap := new(StrByteMap)

	upperBound := 100
	upperBoundStr := string(upperBound)

	for i := 0; i < upperBound; i++ {
		strByteMap.Set(string(i), []byte(string(i)))
	}

	strByteMap.Set(upperBoundStr, []byte(string(upperBound)))

	for n := 0; n < b.N; n++ {
		strByteMap.Get(upperBoundStr)
	}
}

func Benchmark_StrByteMap_Set_Upper_Bound(b *testing.B) {
	strByteMap := new(StrByteMap)

	for i := 0; i < b.N; i++ {
		val := []byte(string(i))

		if i < 100 {
			strByteMap.Set(string(i), val)
		} else {
			strByteMap.Set("1", val)
		}
	}
}

func Benchmark_StrByteMap_Delete_Upper_Bound(b *testing.B) {
	strByteMap := new(StrByteMap)

	for i := 0; i < 100; i++ {
		val := string(i)

		strByteMap.Set(val, []byte(val))
	}

	for i := 0; i < b.N; i++ {
		val := string(i)

		if i < 100 {
			strByteMap.Delete(val)
		} else {
			strByteMap.Delete("1")
		}
	}
}
