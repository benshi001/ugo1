// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// This benchmark test the performance of AND operations on 32-bit ARM.

import (
	"testing"
)

var a []uint32

func init() {
	a = make([]uint32, 16)
}

func doAnd() {
	a[0] &= 0xff0000ff
	a[1] &= 0xffff00ff
	a[2] &= 0xff00ffff
	a[3] &= 0xf00000ff
	a[4] &= 0xff00000f
	a[5] &= 0xfe00007f
	a[6] &= 0xff00007f
	a[7] &= 0xfe0000ff
	a[8] &= 0xfc00007f
	a[9] &= 0xfe00003f
	a[10] &= 0xfc00003f
	a[11] &= 0xfc0000ff
	a[12] &= 0xfffff07f
	a[13] &= 0xff01ffff
	a[14] &= 0xf000003f
	a[15] &= 0xff00003f
}

func BenchmarkAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 3096; j++ {
			doAnd()
		}
	}
}
