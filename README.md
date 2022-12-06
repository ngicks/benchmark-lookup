# benchmark-lookup

look up comparison. Map v.s. Slice.

## Comparison

Slice lookup uses dedicated function, that looks like

```go
func LookUpSlice[T comparable](tab []T, target T) int {
	for idx, val := range tab {
		if val == target {
			return idx
		}
	}

	return -1
}
```

Map lookup uses simply index access to map, that looks like

```go
_, ok := someMap["someIndex"]
if ok {
    // use value
}
```

Those should be what you write once a week.

Each benchmark function does repeative n times of lookup. That significantly increases test time so that difference become clear and noticeable.

## Result

```
$ go version
go version go1.19 linux/amd64
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/ngicks/benchmark-lookup
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkLookUpSlice1-8         1000000000               0.0000037 ns/op
BenchmarkLookUpSlice5-8         1000000000               0.0000110 ns/op
BenchmarkLookUpSlice10-8        1000000000               0.0000090 ns/op
BenchmarkLookUpSlice15-8        1000000000               0.0000227 ns/op
BenchmarkLookUpSlice25-8        1000000000               0.0000223 ns/op
BenchmarkLookUpSlice50-8        1000000000               0.0000363 ns/op
BenchmarkLookUpSlice100-8       1000000000               0.0000775 ns/op
BenchmarkLookUpSlice1000-8      1000000000               0.0008601 ns/op
BenchmarkLookUpMap1-8           1000000000               0.0000053 ns/op
BenchmarkLookUpMap5-8           1000000000               0.0000074 ns/op
BenchmarkLookUpMap10-8          1000000000               0.0000114 ns/op
BenchmarkLookUpMap15-8          1000000000               0.0000118 ns/op
BenchmarkLookUpMap50-8          1000000000               0.0000131 ns/op
BenchmarkLookUpMap100-8         1000000000               0.0000161 ns/op
BenchmarkLookUpMap1000-8        1000000000               0.0000132 ns/op
PASS
ok      github.com/ngicks/benchmark-lookup      0.158s
```

Tried best to be fair comparision.

It seems, if elemens is under 10, slice lookup is slightly faster.
