// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// This benchmark tests the performance of type promotion.
// It shows the improvement in https://go-review.googlesource.com/c/go/+/71992

import (
	"testing"
)

//go:noinline
func typeProAdd(a, b []int8) {
	c[0] = int32(a[0] + b[0])
	c[1] = int32(a[1] + b[1])
	c[2] = int32(a[2] + b[2])
	c[3] = int32(a[3] + b[3])
	c[4] = int32(a[4] + b[4])
	c[5] = int32(a[5] + b[5])
	c[6] = int32(a[6] + b[6])
	c[7] = int32(a[7] + b[7])
	c[8] = int32(a[8] + b[8])
	c[9] = int32(a[9] + b[9])
}

var a0, b0 []int8
var c0 []int32

func init() {
	a0 = make([]int8, 10)
	b0 = make([]int8, 10)
	c0 = make([]int32, 10)
}

func BenchmarkTypePro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			typeProAdd(a0, b0, c0)
		}
	}
}
