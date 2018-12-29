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
strMap.Get("foo")
strMap.Delete("foo")
```

### Benchmarks

````ocaml
tinymap (master) $ ./scripts/bench.sh
+ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/selfup/tinymap
Benchmark_ByteMap_Get_Single_Lower_Bound-8              50000000                28.4 ns/op
Benchmark_ByteMap_Get_Max_Size_Upper_Bound-8             2000000               778 ns/op
Benchmark_IntMap_Get_Single_Lower_Bound-8               1000000000               2.66 ns/op
Benchmark_IntMap_Get_Max_Size_Upper_Bound-8             100000000               19.0 ns/op
Benchmark_IntStrMap_Get_Single_Lower_Bound-8            1000000000               2.89 ns/op
Benchmark_IntStrMap_Get_Max_Size_Upper_Bound-8          100000000               21.9 ns/op
Benchmark_StrMap_Get_Single_Lower_Bound-8               300000000                5.49 ns/op
Benchmark_StrMap_Get_Max_Size_Upper_Bound-8              5000000               378 ns/op
PASS
ok      github.com/selfup/tinymap       18.533s
```

### Details

1. Using an int as an index is fastest for lookups/comparisons.
1. Using []byte as key is the slowest (makes sense)
1. Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.
1. Under the hood TinyMap uses slices to store Tuples.
1. I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:
````
