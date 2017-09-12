// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of frequent -(a*b) calculations.
// It shows the improvement in https://go-review.googlesource.com/c/go/+/61150

package go1

import "testing"

//go:noinline
func fpmul(a, b, c, d float64) (float64, float64, float64, float64, float64, float64) {
	return -a * b, -a * c, -a * d, -b * c, -b * d, -c * d
}

func BenchmarkFPMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _, _, _ = fpmul(12, 45, -23, -234)
		}
	}
}
