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
func typeProAdd(a int8, b int16) (int32, int32, int32) {
	return int32(int16(a) + b), int32(int16(a<<2) + b), int32(int16(a<<4) + b)
}

func BenchmarkTypePro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 65536; j++ {
			_, _, _ = typeProAdd(119, 5000)
		}
	}
}
