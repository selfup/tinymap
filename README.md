[![GoDoc](https://godoc.org/github.com/selfup/tiny_map?status.svg)](https://godoc.org/github.com/selfup/tiny_map)
[![Build Status](https://travis-ci.org/selfup/tiny_map.svg?branch=master)](https://travis-ci.org/selfup/tiny_map)

# Tiny Map!

Package tinymap provides a simple to use interface that behaves like a HashMap.
Multiple "Maps" are exposed for use as well as structs that behave like "Tuples".
This package is intended to be used with tinygo, but has some nice little structs.
Feel free to use this as you wish. It is very small, but helps with embedded programming.

:tada:

### Basic usage

```go
// string/string map

strMap := new(StrMap)

strMap.Set("foo", "bar")
strMap.Get("foo")
strMap.Delete("foo")

// int/int map

intMap := new(IntMap)

intMap.Set(42, 9000)
intMap.Get(42)
intMap.Delete(42)

// int/[]byte map

byteMap := new(ByteMap)

byteMap.Set(42, []byte("9000"))
byteMap.Get(42)
byteMap.Delete(42)
```

### Benchmarks

```ocaml
tiny_map (master) $ ./scripts/bench.sh
+ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/selfup/tiny_map
Benchmark_ByteMap_Get_Single_Lower_Bound-8          1000000000               2.83 ns/op
Benchmark_ByteMap_Get_Max_Size_Upper_Bound-8        20000000                85.8 ns/op
Benchmark_IntMap_Get_Single_Lower_Bound-8           1000000000               2.04 ns/op
Benchmark_IntMap_Get_Max_Size_Upper_Bound-8         30000000                46.0 ns/op
Benchmark_StrMap_Get_Single_Lower_Bound-8           300000000                5.08 ns/op
Benchmark_StrMap_Get_Max_Size_Upper_Bound-8          5000000               374 ns/op
PASS
ok      github.com/selfup/tiny_map      12.938s

---

$ ./scripts/bench.sh
+ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/selfup/tinymap
Benchmark_ByteMap_Get_Single_Lower_Bound-4      500000000                2.95 ns/op
Benchmark_ByteMap_Get_Max_Size_Upper_Bound-4    10000000               210 ns/op
Benchmark_IntMap_Get_Single_Lower_Bound-4       1000000000               2.49 ns/op
Benchmark_IntMap_Get_Max_Size_Upper_Bound-4     10000000               199 ns/op
Benchmark_StrMap_Get_Single_Lower_Bound-4       300000000                4.82 ns/op
Benchmark_StrMap_Get_Max_Size_Upper_Bound-4      5000000               315 ns/op
PASS
ok      github.com/selfup/tinymap       13.517s
```

Using an int as an index is fastest for lookups/comparisons.

Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.

Under the hood TinyMap uses slices to store Tuples.

I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:
