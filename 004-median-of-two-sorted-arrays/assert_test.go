// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"
)

func Assert(tb testing.TB, condition bool, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprint(a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}

func Assertf(tb testing.TB, condition bool, format string, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprintf(format, a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}
