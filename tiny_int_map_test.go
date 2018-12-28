package tinymap

import (
	"testing"
)

func Test_TinyIntMap_Get(t *testing.T) {
	tinyIntMap := new(TinyIntMap)

	result, err := tinyIntMap.Get(1)

	if result != 0 {
		t.Errorf("failed to return int default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_TinyIntMap_Get_Single_Lower_Bound(b *testing.B) {
	tinyIntMap := new(TinyIntMap)

	tinyIntMap.Set(1, 00011000)

	for n := 0; n < b.N; n++ {
		tinyIntMap.Get(1)
	}
}

func Benchmark_TinyIntMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	tinyIntMap := new(TinyIntMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		tinyIntMap.Set(i, i)
	}

	for n := 0; n < b.N; n++ {
		tinyIntMap.Get(upperBound)
	}
}
