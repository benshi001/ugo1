// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// This benchmark tests the performance of the comparason of x&y to 0 and x^y to 0.
// It shows improvement in https://go-review.googlesource.com/c/go/+/67490

import (
	"testing"
)

var a0 int32

func init() {
	a0 = 0
}

//go:noinline
func inc() {
	a0++
}

//go:noinline
func tstteq(a int32) {
	if x&0x000000ff != 0 {
		inc()
	}
	if x&0x0000ff00 != 0 {
		inc()
	}
	if x&0x00ff0000 != 0 {
		inc()
	}
	if x&0xff000000 != 0 {
		inc()
	}
	if x^0x000000ff == 0 {
		inc()
	}
	if x^0x0000ff00 == 0 {
		inc()
	}
	if x^0x00ff0000 == 0 {
		inc()
	}
	if x^0xff000000 == 0 {
		inc()
	}
}

func BenchmarkTSTTEQ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 16384; j++ {
			tstteq(0)
		}
	}
}
