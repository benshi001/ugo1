// Copyright 2018 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of dense (a ^ ^b) and (a | ^b) calculations.
// It shows the improvement of https://go-review.googlesource.com/c/go/+/97036

package go1

import "testing"

//go:noinline
func eonorn(a, b, c, d int64) (int64, int64, int64, int64, int64, int64) {
	return a ^ ^b, b ^ ^c, c | ^d, d | ^a, b ^ ^d, a | ^c
}

func BenchmarkEONORN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _, _, _ = eonorn(0xaa, 0x55, 0x00, 0xff)
		}
	}
}
