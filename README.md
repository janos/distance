# Distance

[![GoDoc](https://godoc.org/resenje.org/distance?status.svg)](https://godoc.org/resenje.org/distance)
[![Build Status](https://travis-ci.org/janos/distance.svg?branch=master)](https://travis-ci.org/janos/distance)

Distance is a Go package with functions to calculate Hamming distance between two byte slices or unsigned integers, population (1 bite) count of byte slices or unsigned integers, and to XOR two byte slices.

## Benchmarks

Benchmarks on MacBook Pro 15'' (Mid 2014):

    BenchmarkPopCountUint64-8       2000000000           0.32 ns/op        0 B/op          0 allocs/op
    BenchmarkPopCount8Bytes-8       200000000            7.06 ns/op        0 B/op          0 allocs/op
    BenchmarkXOR8Bytes-8            100000000           11.1 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming1Byte-8         200000000            6.83 ns/op        0 B/op          0 allocs/op
    BenchmarkHamming8Bytes-8        100000000           10.0 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming16Bytes-8       100000000           13.5 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming32Bytes-8       100000000           19.9 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming33Bytes-8       100000000           19.0 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming64Bytes-8       50000000            33.8 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming128Bytes-8      20000000            59.7 ns/op         0 B/op          0 allocs/op
    BenchmarkHamming1KBytes-8        5000000           395 ns/op           0 B/op          0 allocs/op
    BenchmarkHamming10KBytes-8        500000          3724 ns/op           0 B/op          0 allocs/op
    BenchmarkHamming1MByte-8            3000        389581 ns/op           0 B/op          0 allocs/op

## License

Source files are distributed under the BSD-style license found in the LICENSE file.