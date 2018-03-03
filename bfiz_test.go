// Copyright 2018 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// This benchmark tests the performance of bit field operation.
// It shows improvement in https://go-review.googlesource.com/c/go/+/98455

import (
	"testing"
)

var ay []int32

func init() {
	ay = make([]int32, 10)
}

func bfiz() {
	ay[0] = ay[0] << 12 >> 16
	ay[1] = ay[1] << 12 >> 16
	ay[2] = ay[2] << 12 >> 16
	ay[3] = ay[3] << 12 >> 16
	ay[4] = ay[4] << 12 >> 16
	ay[5] = ay[5] << 12 >> 16
	ay[6] = ay[6] << 12 >> 16
	ay[7] = ay[7] << 12 >> 16
	ay[8] = ay[8] << 12 >> 16
	ay[9] = ay[9] << 12 >> 16
}

func BenchmarkBFIZ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 4096; j++ {
			bfiz()
		}
	}
}
