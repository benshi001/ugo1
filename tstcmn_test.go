// Copyright 2017 Ben Shi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go1

// This benchmark tests the performance of the comparason of x&y to 0 and x+y to 0.
// It shows improvement in https://go-review.googlesource.com/c/go/+/124935

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
func tstcmn(a, b, c, d int64) {
	if a&b > 0 {
		inc()
	}
	if a+b > 0 {
		inc()
	}
	if int32(a)+int32(b) > 0 {
		inc()
	}
	if int32(a)&int32(b) > 0 {
		inc()
	}
	if c&d > 0 {
		inc()
	}
	if c+d > 0 {
		inc()
	}
	if int32(c)+int32(d) > 0 {
		inc()
	}
	if int32(c)&int32(d) > 0 {
		inc()
	}
}

func BenchmarkTSTCMN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 16384; j++ {
			tstcmn(0, 0, 0, 0)
		}
	}
}
