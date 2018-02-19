// Copyright 2018 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of dense -(a*b) calculations.
// It shows the improvement of https://go-review.googlesource.com/c/go/+/95075

package go1

import "testing"

//go:noinline
func mneg(a, b, c, d int64) (int64, int64, int64, int64, int64, int64) {
	return -a * b, -a * c, -a * d, -b * c, -b * d, -c * d
}

func BenchmarkMNEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _, _, _ = mneg(12, 45, -23, -234)
		}
	}
}
