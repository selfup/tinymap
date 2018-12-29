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

func Benchmark_StrMap_Get_Single_Lower_Bound(b *testing.B) {
	strMap := new(StrMap)

	strMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		strMap.Get("a")
	}
}

func Benchmark_StrMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	strMap := new(StrMap)

	for i := 0; i < 100; i++ {
		strMap.Set(string(i), "_")
	}

	strMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		strMap.Get("a")
	}
}
