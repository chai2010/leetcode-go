// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestSolutionV0(t *testing.T) {
	for _, v := range tests {
		if v._break {
			break
		}
		result := SolutionV0(v.nums1, v.nums2)
		Assert(t, result == v.result, v, result)
	}
}

var tests = []struct {
	nums1, nums2 []int
	result       [2]int
	_break       bool
}{
	// empty
	{
		nums1:  []int{},
		nums2:  []int{},
		result: [2]int{},
	},

	// only one slice
	{
		nums1:  []int{1, 2, 3},
		nums2:  []int{},
		result: [2]int{2, 2},
	},
	{
		nums1:  []int{},
		nums2:  []int{1, 2, 3},
		result: [2]int{2, 2},
	},

	// only one slice
	{
		nums1:  []int{1, 2, 3, 4},
		nums2:  []int{},
		result: [2]int{2, 3},
	},
	{
		nums1:  []int{},
		nums2:  []int{1, 2, 3, 4},
		result: [2]int{2, 3},
	},

	// break
	{
		_break: true,
	},

	{
		nums1:  []int{1, 3},
		nums2:  []int{2},
		result: [2]int{2, 2},
	},
	{
		nums1:  []int{1, 2},
		nums2:  []int{3, 4},
		result: [2]int{2, 3},
	},

	// 123456789 => 5
	{
		nums1:  []int{1, 3, 5, 7, 8, 9},
		nums2:  []int{2, 4, 6},
		result: [2]int{5, 5},
	},
	{
		nums1:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		nums2:  []int{},
		result: [2]int{5, 5},
	},
	{
		nums1:  []int{1, 2, 3, 4},
		nums2:  []int{5, 6, 7, 8, 9},
		result: [2]int{5, 5},
	},

	// 12345678 => 45
	{
		nums1:  []int{1, 3, 5, 7, 8},
		nums2:  []int{2, 4, 6},
		result: [2]int{4, 5},
	},
	{
		nums1:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		nums2:  []int{},
		result: [2]int{4, 5},
	},
	{
		nums1:  []int{1, 2, 3, 4},
		nums2:  []int{5, 6, 7, 8},
		result: [2]int{4, 5},
	},
}
