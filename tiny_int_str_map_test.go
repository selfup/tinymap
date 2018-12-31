package tinymap

import (
	"testing"
)

func Test_IntStrMap_Get(t *testing.T) {
	intStrMap := new(IntStrMap)

	result, err := intStrMap.Get(42)

	if result != "" {
		t.Errorf("failed to return string default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Test_IntStrMap_Set(t *testing.T) {
	intStrMap := new(IntStrMap)

	intStrMap.Set(42, "9000")

	result, _ := intStrMap.Get(42)

	if result != "9000" {
		t.Errorf("failed to return correct Val from StrTuple")
	}
}

func Benchmark_IntStrMap_Get_Lower_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	intStrMap.Set(0, "_")

	for n := 0; n < b.N; n++ {
		intStrMap.Get(0)
	}
}

func Benchmark_IntStrMap_Get_Expected_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	intStrMap.Set(0, "8996")
	intStrMap.Set(1, "8997")
	intStrMap.Set(2, "8998")
	intStrMap.Set(3, "8999")
	intStrMap.Set(42, "9000")

	for n := 0; n < b.N; n++ {
		intStrMap.Get(42)
	}
}

func Benchmark_IntStrMap_Get_Upper_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		intStrMap.Set(i, "_")
	}

	intStrMap.Set(upperBound, "_")

	for n := 0; n < b.N; n++ {
		intStrMap.Get(upperBound)
	}
}

func Benchmark_IntStrMap_Set_Upper_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	for i := 0; i < b.N; i++ {
		val := string(i)

		if i < 100 {
			intStrMap.Set(i, val)
		} else {
			intStrMap.Set(1, val)
		}
	}
}

func Benchmark_IntStrMap_Delete_Upper_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	for i := 0; i < 100; i++ {
		val := string(i)

		intStrMap.Set(i, val)
	}

	for i := 0; i < b.N; i++ {
		if i < 100 {
			intStrMap.Delete(i)
		} else {
			intStrMap.Delete(i)
		}
	}
}
