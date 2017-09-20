// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// benchmark for testing performance of bit field extraction.

import (
	"testing"
)

var ar []int32

func init() {
	ar = make([]int32, 10)
}

func bfx() {
	ar[0] = ar[0] << 12 >> 16
	ar[1] = ar[1] << 12 >> 16
	ar[2] = ar[2] << 12 >> 16
	ar[3] = ar[3] << 12 >> 16
	ar[4] = ar[4] << 12 >> 16
	ar[5] = ar[5] << 12 >> 16
	ar[6] = ar[6] << 12 >> 16
	ar[7] = ar[7] << 12 >> 16
	ar[8] = ar[8] << 12 >> 16
	ar[9] = ar[9] << 12 >> 16
}

func BenchmarkBFX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 3096; j++ {
			bfx()
		}
	}
}
