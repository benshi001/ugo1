// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This benchmark tests the performance for dense byte store operations.
// It shows the improvement in https://go-review.googlesource.com/c/go/58450

package go1

import "testing"

var ga []int8

//go:noinline
func set(i int, v int8) {
	ga[i] = v
	ga[i+1] = v + 12
	ga[i+2] = v + 18
	ga[i+3] = v + 21
	ga[i+4] = v + 70
	ga[i+5] = v + 19
	ga[i+6] = v + 12
	ga[i+7] = v + 11
	ga[i+8] = v + 16
	ga[i+9] = v + 14
	ga[i+10] = v + 12
	ga[i+11] = v + 11
	ga[i+12] = v + 13
}

func init() {
	ga = make([]int8, 64)
}

func BenchmarkByteStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			set(41, 10)
			set(42, 20)
			set(43, 11)
			set(44, 17)
			set(14, 22)
			set(47, 33)
			set(49, 55)
		}
	}
}
