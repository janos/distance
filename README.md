# Distance

[![GoDoc](https://godoc.org/resenje.org/distance?status.svg)](https://godoc.org/resenje.org/distance)

Distance is a Go package with functions to calculate Hamming distance between two byte slices or unsigned integers, population (1 bite) count of byte slices or unsigned integers, and to XOR two byte slices.

## Benchmarks

Benchmarks on MacBook Pro 15'' (Mid 2014):

    BenchmarkPopCountUint64-8       1000000000               0.299 ns/op           0 B/op          0 allocs/op
    BenchmarkPopCount8Bytes-8       170056086                7.28 ns/op            0 B/op          0 allocs/op
    BenchmarkXOR8Bytes-8            126639039                9.45 ns/op            0 B/op          0 allocs/op
    BenchmarkHamming1Byte-8         219764226                5.42 ns/op            0 B/op          0 allocs/op
    BenchmarkHamming8Bytes-8        141195171                8.59 ns/op            0 B/op          0 allocs/op
    BenchmarkHamming16Bytes-8       92126083                12.4 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming32Bytes-8       56141322                20.2 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming33Bytes-8       53014400                20.8 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming64Bytes-8       31677573                35.8 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming128Bytes-8      17375133                67.1 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming1KBytes-8        2557718               458 ns/op               0 B/op          0 allocs/op
    BenchmarkHamming10KBytes-8        257976              4370 ns/op               0 B/op          0 allocs/op
    BenchmarkHamming1MByte-8            2464            444376 ns/op               0 B/op          0 allocs/op

## License

Source files are distributed under the BSD-style license found in the LICENSE file.
