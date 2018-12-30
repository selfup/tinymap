package tinymap

import (
	"testing"
)

func Test_StrMap_Get(t *testing.T) {
	strMap := new(StrMap)

	result, err := strMap.Get("a")

	if result != "" {
		t.Errorf("failed to return string default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
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
