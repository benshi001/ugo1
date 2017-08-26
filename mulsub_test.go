// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of dense mul-sub operations.
// It shows the improvement in https://go-review.googlesource.com/c/go/58950

package go1

import "testing"

//go:noinline
func muls(a, b, c, d int32) (int32, int32, int32, int32, int32, int32, int32, int32, int32) {
	return a - b*c, b - a*c, c - a*b, b - c*d, c - b*d, d - b*c, a - c*d, c - a*d, d - a*c
}

func BenchmarkMulSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _, _, _, _, _, _ = muls(12, 45, -23, -234)
		}
	}
}
