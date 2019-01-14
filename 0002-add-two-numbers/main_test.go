// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
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
