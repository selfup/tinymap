package tinymap

import (
	"testing"
)

func Test_StrMap_Get(t *testing.T) {
	strMap := new(StrMap)

	result, err := strMap.Get("foo")

	if result != "" {
		t.Errorf("failed to return string default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Test_StrMap_Set(t *testing.T) {
	strMap := new(StrMap)

	strMap.Set("foo", "bar")

	result, _ := strMap.Get("foo")

	if result != "bar" {
		t.Errorf("failed to return correct Val from StrTuple")
	}
}

func Benchmark_StrMap_Get_Lower_Bound(b *testing.B) {
	strMap := new(StrMap)

	strMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		strMap.Get("a")
	}
}

func Benchmark_StrMap_Get_Expected_Bound(b *testing.B) {
	strMap := new(StrMap)

	strMap.Set("0", "8996")
	strMap.Set("0", "8997")
	strMap.Set("0", "8998")
	strMap.Set("0", "8999")
	strMap.Set("42", "9000")

	for n := 0; n < b.N; n++ {
		strMap.Get("42")
	}
}

func Benchmark_StrMap_Get_Upper_Bound(b *testing.B) {
	strMap := new(StrMap)

	upperBound := 100
	upperBoundStr := string(upperBound)

	for i := 0; i < upperBound; i++ {
		strMap.Set(string(i), string(i))
	}

	strMap.Set(upperBoundStr, upperBoundStr)

	for n := 0; n < b.N; n++ {
		strMap.Get(upperBoundStr)
	}
}

func Benchmark_StrMap_Set_Upper_Bound(b *testing.B) {
	strMap := new(StrMap)

	for i := 0; i < b.N; i++ {
		val := string(i)

		if i < 100 {
			strMap.Set(val, val)
		} else {
			strMap.Set("1", val)
		}
	}
}

func Benchmark_StrMap_Delete_Upper_Bound(b *testing.B) {
	strMap := new(StrMap)

	for i := 0; i < 100; i++ {
		val := string(i)

		strMap.Set(val, val)
	}

	for i := 0; i < b.N; i++ {
		val := string(i)

		if i < 100 {
			strMap.Delete(val)
		} else {
			strMap.Delete("1")
		}
	}
}
