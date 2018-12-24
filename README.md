# Tiny Map!

:tada:

### Basic usage

```go
tinyStrMap := new(TinyStrMap)

tinyStrMap.Set("a", "_") // => bool
tinyStrMap.Get("a") // => string, bool
tinyStrMap.Delete("a") // => bool
```

### Benchmarks

```ocaml
tiny_map (master) $ ./scripts/bench.sh
+ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/selfup/tiny_map
Benchmark_TinyIntMap_Get_Single_Lower_Bound-8           1000000000               2.04 ns/op
Benchmark_TinyIntMap_Get_Max_Size_Upper_Bound-8         30000000                45.8 ns/op
Benchmark_TinyStrMap_Get_Single_Lower_Bound-8           300000000                5.07 ns/op
Benchmark_TinyStrMap_Get_Max_Size_Upper_Bound-8          5000000               377 ns/op
PASS
ok      github.com/selfup/tiny_map      8.016s
```

Upper Bound means the value being grabbed is the 100th element in a slice. So 333 ns/op is the constant lookup time for the last element you have included.

I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:

Under the hood TinyMap uses slices. Arrays could be used next if performance becomes an issue :smile:
