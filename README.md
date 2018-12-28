[![GoDoc](https://godoc.org/github.com/selfup/tiny_map?status.svg)](https://godoc.org/github.com/selfup/tiny_map)

# Tiny Map!

:tada:

### Basic usage

```go
// string/string map

tinyStrMap := new(TinyStrMap)

tinyStrMap.Set("foo", "bar")
tinyStrMap.Get("foo")
tinyStrMap.Delete("foo")

// int/int map

tinyIntMap := new(TinyIntMap)

tinyIntMap.Set(42, 9000)
tinyIntMap.Get(42)
tinyIntMap.Delete(42)

// int/[]byte map

tinyByteMap := new(TinyByteMap)

tinyByteMap.Set(42, []byte("9000"))
tinyByteMap.Get(42)
tinyByteMap.Delete(42)
```

### Benchmarks

```ocaml
tiny_map (master) $ ./scripts/bench.sh
+ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/selfup/tiny_map
Benchmark_TinyByteMap_Get_Single_Lower_Bound-8          1000000000               2.83 ns/op
Benchmark_TinyByteMap_Get_Max_Size_Upper_Bound-8        20000000                85.8 ns/op
Benchmark_TinyIntMap_Get_Single_Lower_Bound-8           1000000000               2.04 ns/op
Benchmark_TinyIntMap_Get_Max_Size_Upper_Bound-8         30000000                46.0 ns/op
Benchmark_TinyStrMap_Get_Single_Lower_Bound-8           300000000                5.08 ns/op
Benchmark_TinyStrMap_Get_Max_Size_Upper_Bound-8          5000000               374 ns/op
PASS
ok      github.com/selfup/tiny_map      12.938s
```

Clearly using an int as an index is fastest for lookups/comparisons.

Upper Bound means the value being grabbed is the 100th element in a slice of 100 elements.

Under the hood TinyMap uses slices.

I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:
