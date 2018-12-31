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
Benchmark_ByteMap_Get_Lower_Bound-4             50000000                27.7 ns/op
Benchmark_ByteMap_Get_Expected_Bound-4          50000000                34.5 ns/op
Benchmark_ByteMap_Get_Upper_Bound-4              3000000               544 ns/op
Benchmark_ByteMap_Set_Upper_Bound-4              5000000               287 ns/op
Benchmark_ByteMap_Delete_Upper_Bound-4          100000000               11.3 ns/op
Benchmark_IntMap_Get_Lower_Bound-4              1000000000               2.55 ns/op
Benchmark_IntMap_Get_Expected_Bound-4           300000000                4.28 ns/op
Benchmark_IntMap_Get_Upper_Bound-4              30000000                45.1 ns/op
Benchmark_IntMap_Set_Upper_Bound-4              1000000000               2.83 ns/op
Benchmark_IntMap_Delete_Upper_Bound-4           2000000000               1.85 ns/op
Benchmark_IntStrMap_Get_Lower_Bound-4           1000000000               2.70 ns/op
Benchmark_IntStrMap_Get_Expected_Bound-4        300000000                5.05 ns/op
Benchmark_IntStrMap_Get_Upper_Bound-4           30000000                56.6 ns/op
Benchmark_IntStrMap_Set_Upper_Bound-4           100000000               20.4 ns/op
Benchmark_IntStrMap_Delete_Upper_Bound-4        2000000000               1.85 ns/op
Benchmark_StrMap_Get_Lower_Bound-4              300000000                4.89 ns/op
Benchmark_StrMap_Get_Expected_Bound-4           300000000                5.51 ns/op
Benchmark_StrMap_Get_Upper_Bound-4               5000000               391 ns/op
Benchmark_StrMap_Set_Upper_Bound-4              10000000               201 ns/op
Benchmark_StrMap_Delete_Upper_Bound-4           200000000                6.37 ns/op
PASS
ok      github.com/selfup/tinymap       44.851s
```

### Details

1. Using an int as an index is fastest for lookups/comparisons.
1. Using []byte as key is the slowest (makes sense)
1. Lower Bound means the value being grabbed is the 1th element in a slice of 1 element.
1. Expected Bound means the value is being grabbed at the 5th element in a slice of 5 elements.
1. Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.
1. Under the hood TinyMap uses slices to store Tuples.
1. I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:
