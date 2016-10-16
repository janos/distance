// Copyright (c) 2016 Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distance // import "resenje.org/distance"

// Hamming calculates the Hamming distance between two byte slices.
// Length of both slices can be arbitrary, but the distance
// will not be calculated on the remaining part of the longer slice.
func Hamming(a, b []byte) (c uint64) {
	l := len(a)
	lb := len(b)
	if lb < l {
		a, b = b, a
		l = lb
	}
	for l >= 8 {
		c += PopCountUint64(littleEndianUint64(a[:8]) ^ littleEndianUint64(b[:8]))
		a = a[8:]
		b = b[8:]
		l = len(a)
	}
	switch l {
	case 7:
		c += PopCountUint64(littleEndianUint56(a) ^ littleEndianUint56(b))
	case 6:
		c += PopCountUint64(littleEndianUint48(a) ^ littleEndianUint48(b))
	case 5:
		c += PopCountUint64(littleEndianUint40(a) ^ littleEndianUint40(b))
	case 4:
		c += uint64(PopCountUint32(littleEndianUint32(a) ^ littleEndianUint32(b)))
	case 3:
		c += uint64(PopCountUint32(littleEndianUint24(a) ^ littleEndianUint24(b)))
	case 2:
		c += uint64(PopCountUint16(littleEndianUint16(a) ^ littleEndianUint16(b)))
	case 1:
		c += uint64(PopCountUint8(littleEndianUint8(a) ^ littleEndianUint8(b)))
	}
	return
}

// PopCount counts the number of 1 bite in the byte slice.
func PopCount(buf []byte) (c uint64) {
	l := len(buf)
	for l >= 8 {
		c += PopCountUint64(littleEndianUint64(buf[:8]))
		buf = buf[8:]
		l = len(buf)
	}
	switch l {
	case 7:
		c += PopCountUint64(littleEndianUint56(buf))
	case 6:
		c += PopCountUint64(littleEndianUint48(buf))
	case 5:
		c += PopCountUint64(littleEndianUint40(buf))
	case 4:
		c += uint64(PopCountUint32(littleEndianUint32(buf)))
	case 3:
		c += uint64(PopCountUint32(littleEndianUint24(buf)))
	case 2:
		c += uint64(PopCountUint16(littleEndianUint16(buf)))
	case 1:
		c += uint64(PopCountUint8(littleEndianUint8(buf)))
	}
	return
}

// PopCountUint64 counts the number of 1 bite in uint64 integer.
func PopCountUint64(v uint64) uint64 {
	v = v - ((v >> 1) & 0x5555555555555555)
	v = (v & 0x3333333333333333) + ((v >> 2) & 0x3333333333333333)
	return (((v + (v >> 4)) & 0x0F0F0F0F0F0F0F0F) * 0x0101010101010101) >> 56
}

// PopCountUint32 counts the number of 1 bite in uint32 integer.
func PopCountUint32(v uint32) uint32 {
	v = v - ((v >> 1) & 0x55555555)
	v = (v & 0x33333333) + ((v >> 2) & 0x33333333)
	return (((v + (v >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24
}

// PopCountUint16 counts the number of 1 bite in uint16 integer.
func PopCountUint16(v uint16) uint16 {
	v = v - ((v >> 1) & 0x5555)
	v = (v & 0x3333) + ((v >> 2) & 0x3333)
	return (((v + (v >> 4)) & 0x0F0F) * 0x0101) >> 8
}

// PopCountUint8 counts the number of 1 bite in uint8 integer.
func PopCountUint8(v uint8) uint8 {
	v = v - ((v >> 1) & 0x55)
	v = (v & 0x33) + ((v >> 2) & 0x33)
	return (((v + (v >> 4)) & 0x0F) * 0x01)
}

// XOR calculates the difference byte slice between two byte slices.
// Length of both slices can be arbitrary, but the difference
// will not be calculated on the remaining part of the longer slice.
func XOR(a, b []byte) []byte {
	l := len(a)
	lb := len(b)
	if lb < l {
		a, b = b, a
		l = lb
	}
	for i := 0; i < l; i++ {
		a[i] = a[i] ^ b[i]
	}
	return a
}

func littleEndianUint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func littleEndianUint56(b []byte) uint64 {
	_ = b[6]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
}

func littleEndianUint48(b []byte) uint64 {
	_ = b[5]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40
}

func littleEndianUint40(b []byte) uint64 {
	_ = b[4]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32
}

func littleEndianUint32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func littleEndianUint24(b []byte) uint32 {
	_ = b[2]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

func littleEndianUint16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[0]) | uint16(b[1])<<8
}

func littleEndianUint8(b []byte) uint8 {
	_ = b[0]
	return uint8(b[0])
}
