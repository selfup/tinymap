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

On Windows:

```ocaml
$ ./scripts/bench.sh
goos: windows
goarch: amd64
pkg: github.com/selfup/tiny_map
Benchmark_Get_Single_Lower_Bound-4      300000000                4.81 ns/op
Benchmark_Get_Max_Size_Upper_Bound-4     5000000               333 ns/op
PASS
ok      github.com/selfup/tiny_map      4.109s
```

Upper Bound means the value being grabbed is the 100th element in a slice. So 333 ns/op is the constant lookup time for the last element you have included.

I have not yet added a catch block to prevent the slice to grow, but this should be used for small data sotrage :pray:

Under the hood TinyMap uses slices. Arrays could be used next if performance becomes an issue :smile:
