package tinymap

import (
	"testing"
)

func Test_IntMap_Get(t *testing.T) {
	intMap := new(IntMap)

	result, err := intMap.Get(1)

	if result != 0 {
		t.Errorf("failed to return int default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_IntMap_Get_Single_Lower_Bound(b *testing.B) {
	intMap := new(IntMap)

	intMap.Set(42, 9000)

	for n := 0; n < b.N; n++ {
		intMap.Get(42)
	}
}

func Benchmark_IntMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	intMap := new(IntMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		intMap.Set(i, i)
	}

	intMap.Set(upperBound, upperBound)

	for n := 0; n < b.N; n++ {
		intMap.Get(upperBound)
	}
}
