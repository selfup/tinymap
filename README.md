[![GoDoc](https://godoc.org/github.com/selfup/tinymap?status.svg)](https://godoc.org/github.com/selfup/tinymap)
[![Build Status](https://travis-ci.org/selfup/tinymap.svg?branch=master)](https://travis-ci.org/selfup/tinymap)

# tinymap

Provides a simple to use interface that behaves like a `HashMap`.

Multiple `Maps` are exposed for use as well as structs that behave like `Tuples`.

This package is intended to be used with `tinygo`.

However it has some nice little structs, so feel free to use this as you wish.

It is very small, but helps with embedded programming.

:tada:

### Basic usage

Just one of the many Maps :smile:

```go
strMap := new(StrMap)

strMap.Set("foo", "bar")

val, err := strMap.Get("foo")

if err != nil {
  log.Fatal(err)
}

fmt.Print(val)

strMap.Delete(42)
```

### Benchmarks

```ocaml
$ ./scripts/bench.sh
+ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/selfup/tinymap
Benchmark_ByteMap_Get_Lower_Bound-4             50000000                25.7 ns/op
Benchmark_ByteMap_Get_Expected_Bound-4          50000000                29.3 ns/op
Benchmark_ByteMap_Get_Upper_Bound-4              3000000               530 ns/op
Benchmark_ByteMap_Set_Upper_Bound-4              5000000               302 ns/op
Benchmark_ByteMap_Delete_Upper_Bound-4          100000000               11.7 ns/op
Benchmark_IntMap_Get_Lower_Bound-4              1000000000               2.51 ns/op
Benchmark_IntMap_Get_Expected_Bound-4           300000000                4.19 ns/op
Benchmark_IntMap_Get_Upper_Bound-4              30000000                45.3 ns/op
Benchmark_IntMap_Set_Upper_Bound-4              1000000000               2.84 ns/op
Benchmark_IntMap_Delete_Upper_Bound-4           1000000000               2.09 ns/op
Benchmark_IntStrMap_Get_Lower_Bound-4           1000000000               2.69 ns/op
Benchmark_IntStrMap_Get_Expected_Bound-4        300000000                4.84 ns/op
Benchmark_IntStrMap_Get_Upper_Bound-4           30000000                55.8 ns/op
Benchmark_IntStrMap_Set_Upper_Bound-4           100000000               20.1 ns/op
Benchmark_IntStrMap_Delete_Upper_Bound-4        1000000000               2.07 ns/op
Benchmark_StrByteMap_Get_Lower_Bound-4          200000000                6.91 ns/op
Benchmark_StrByteMap_Get_Expected_Bound-4       100000000               14.8 ns/op
Benchmark_StrByteMap_Get_Upper_Bound-4           3000000               444 ns/op
Benchmark_StrByteMap_Set_Upper_Bound-4           5000000               253 ns/op
Benchmark_StrByteMap_Delete_Upper_Bound-4       200000000                6.45 ns/op
Benchmark_StrMap_Get_Lower_Bound-4              300000000                4.86 ns/op
Benchmark_StrMap_Get_Expected_Bound-4           300000000                4.95 ns/op
Benchmark_StrMap_Get_Upper_Bound-4               5000000               325 ns/op
Benchmark_StrMap_Set_Upper_Bound-4              10000000               196 ns/op
Benchmark_StrMap_Delete_Upper_Bound-4           200000000                6.33 ns/op
PASS
ok      github.com/selfup/tinymap       49.282s
```

### Details

1. Using an int as an index is fastest for lookups/comparisons.
1. Using []byte as key is the slowest (makes sense)
1. Lower Bound means the value being grabbed is the 1th element in a slice of 1 element.
1. Expected Bound means the value is being grabbed at the 5th element in a slice of 5 elements.
1. Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.
1. Under the hood TinyMap uses slices to store Tuples.
1. I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:
