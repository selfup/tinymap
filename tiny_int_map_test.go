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

func Test_IntMap_Set(t *testing.T) {
	intMap := new(IntMap)

	intMap.Set(42, 9000)

	result, _ := intMap.Get(42)

	if result != 9000 {
		t.Errorf("failed to return correct Val from StrTuple")
	}
}

func Benchmark_IntMap_Get_Lower_Bound(b *testing.B) {
	intMap := new(IntMap)

	intMap.Set(42, 9000)

	for n := 0; n < b.N; n++ {
		intMap.Get(42)
	}
}

func Benchmark_IntMap_Get_Expected_Bound(b *testing.B) {
	intMap := new(IntMap)

	intMap.Set(0, 8996)
	intMap.Set(1, 8997)
	intMap.Set(2, 8998)
	intMap.Set(3, 8999)
	intMap.Set(42, 9000)

	for n := 0; n < b.N; n++ {
		intMap.Get(42)
	}
}

func Benchmark_IntMap_Get_Upper_Bound(b *testing.B) {
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

func Benchmark_IntMap_Set_Upper_Bound(b *testing.B) {
	intMap := new(IntMap)

	for i := 0; i < b.N; i++ {
		if i < 100 {
			intMap.Set(i, i)
		} else {
			intMap.Set(1, i)
		}
	}
}

func Benchmark_IntMap_Delete_Upper_Bound(b *testing.B) {
	intMap := new(IntMap)

	for i := 0; i < 100; i++ {
		intMap.Set(i, i)
	}

	for i := 0; i < b.N; i++ {
		if i < 100 {
			intMap.Delete(i)
		} else {
			intMap.Delete(i)
		}
	}
}
