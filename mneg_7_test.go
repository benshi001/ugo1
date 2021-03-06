// Copyright 2018 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of dense -(a*b) calculations, b = 7*pow(2, n).
// It shows the improvement of https://go-review.googlesource.com/c/go/+/99495

package go1

import "testing"

//go:noinline
func mneg(a, b, c, d int64) (int64, int64, int64, int64) {
	return -a * 112, -c * 112, -d * 112, -b * 112
}

func BenchmarkMNEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _ = mneg(12, 45, -23, -234)
		}
	}
}
