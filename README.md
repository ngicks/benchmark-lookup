# Result

```
$ go test -benchmem -run=^$ -bench ^Benchmark github.com/ngicks/benchmark-lookup
goos: linux
goarch: amd64
pkg: github.com/ngicks/benchmark-lookup
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkLookUpSlice1-8         1000000000               0.0000047 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice5-8         1000000000               0.0000063 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice10-8        1000000000               0.0000104 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice15-8        1000000000               0.0000251 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice25-8        1000000000               0.0000210 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice50-8        1000000000               0.0000368 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice100-8       1000000000               0.0000651 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpSlice1000-8      1000000000               0.0005458 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap1-8           1000000000               0.0000052 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap5-8           1000000000               0.0000086 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap10-8          1000000000               0.0000193 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap15-8          1000000000               0.0000186 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap50-8          1000000000               0.0000131 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap100-8         1000000000               0.0000153 ns/op               0 B/op          0 allocs/op
BenchmarkLookUpMap1000-8        1000000000               0.0000132 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/ngicks/benchmark-lookup      0.147s
```

Tried best to be fair comparision.

It seems, if elemens is under 10, slice lookup is slightly faster.
