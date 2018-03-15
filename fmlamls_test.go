// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance of frequent -(a*b) calculations.
// It shows the improvement in https://go-review.googlesource.com/c/go/+/100855

package go1

import "testing"

//go:noinline
func fmlamls(a, b, c, d float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	return c - a*b, c + a*b, b - a*d, b + a*d, c - b*d, c + b*d, a + c*d, a - c*d
}

func BenchmarkMLAMLS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			_, _, _, _, _, _, _, _ = fmlamls(12, 45, -23, -234)
		}
	}
}
