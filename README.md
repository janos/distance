# Distance

[![GoDoc](https://godoc.org/resenje.org/distance?status.svg)](https://godoc.org/resenje.org/distance)
[![Build Status](https://travis-ci.org/janos/distance.svg?branch=master)](https://travis-ci.org/janos/distance)

Distance is a Go package with functions to calculate Hamming distance between two byte slices or unsigned integers, population (1 bite) count of byte slices or unsigned integers, and to XOR two byte slices.

## Benchmarks

Benchmarks on MacBook Pro 15'' (Mid 2014):

    BenchmarkPopCountUint64-8       2000000000               0.30 ns/op            0 B/op          0 allocs/op
    BenchmarkPopCount8Bytes-8       200000000                7.10 ns/op            0 B/op          0 allocs/op
    BenchmarkXOR8Bytes-8            100000000               12.0 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming1Byte-8         200000000                7.23 ns/op            0 B/op          0 allocs/op
    BenchmarkHamming8Bytes-8        100000000               10.0 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming16Bytes-8       100000000               14.1 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming32Bytes-8       100000000               22.9 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming33Bytes-8       50000000                23.4 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming64Bytes-8       30000000                39.8 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming128Bytes-8      20000000                76.5 ns/op             0 B/op          0 allocs/op
    BenchmarkHamming1KBytes-8        3000000               465 ns/op               0 B/op          0 allocs/op
    BenchmarkHamming10KBytes-8        300000              4396 ns/op               0 B/op          0 allocs/op
    BenchmarkHamming1MByte-8            3000            465924 ns/op               0 B/op          0 allocs/op

## License

Source files are distributed under the BSD-style license found in the LICENSE file.