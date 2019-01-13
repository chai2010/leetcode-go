// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	for i, testCase := range tests {
		result := addTwoNumbers(
			NewListNode(testCase.a...),
			NewListNode(testCase.b...),
		)
		Assertf(t, reflect.DeepEqual(result.Slice(), testCase.sum),
			"%d: expect = %v, got = %v",
			i, testCase.sum, result,
		)
	}
}

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

var tests = []struct {
	a, b, sum []int
}{
	{
		a:   []int{2, 4, 3},
		b:   []int{5, 6, 4},
		sum: []int{7, 0, 8},
	},
	{
		a:   []int{0},
		b:   []int{7, 3},
		sum: []int{7, 3},
	},
	{
		a:   []int{5},
		b:   []int{5},
		sum: []int{0, 1},
	},
}
