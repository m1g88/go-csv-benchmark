# CSV Read and ReadAll benchmark

benchmake between read and readall

`$ go test -bench=. -benchmem`
```
Benchmark_Read-8          153525             11343 ns/op            4808 B/op         21 allocs/op
Benchmark_ReadAll-8       130141             10050 ns/op            4976 B/op         24 allocs/op
```
