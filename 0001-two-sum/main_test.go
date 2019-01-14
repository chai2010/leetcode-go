// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"testing"
)

func TestTwoSumV0(t *testing.T) {
	for i, testCase := range tests {
		result := twoSumV0(testCase.nums, testCase.target)
		Assertf(t, reflect.DeepEqual(result, testCase.result),
			"%d: expect = %v, got = %v",
			i, testCase.result, result,
		)
	}
}

func TestTwoSumV1(t *testing.T) {
	for i, testCase := range tests {
		result := twoSumV1(testCase.nums, testCase.target)
		Assertf(t, reflect.DeepEqual(result, testCase.result),
			"%d: expect = %v, got = %v",
			i, testCase.result, result,
		)
	}
}

func TestTwoSumV2(t *testing.T) {
	for i, testCase := range tests {
		result := twoSumV2(testCase.nums, testCase.target)
		Assertf(t, reflect.DeepEqual(result, testCase.result),
			"%d: expect = %v, got = %v",
			i, testCase.result, result,
		)
	}
}

func TestTwoSumV3(t *testing.T) {
	for i, testCase := range tests {
		result := twoSumV3(testCase.nums, testCase.target)
		Assertf(t, reflect.DeepEqual(result, testCase.result),
			"%d: expect = %v, got = %v",
			i, testCase.result, result,
		)
	}
}

func TestTwoSumV4(t *testing.T) {
	for i, testCase := range tests {
		result := twoSumV4(testCase.nums, testCase.target)
		Assertf(t, reflect.DeepEqual(result, testCase.result),
			"%d: expect = %v, got = %v",
			i, testCase.result, result,
		)
	}
}

var tests = []struct {
	nums   []int
	target int
	result []int
}{
	{
		result: []int{0, 1},
		nums:   []int{2, 7, 11, 15},
		target: 9,
	},
	{
		result: []int{2, 4},
		nums:   []int{1, 2, 3, 4, 5},
		target: 8,
	},
	{
		result: []int{3, 4},
		nums:   []int{1, 2, 3, 4, 5},
		target: 9,
	},
	{
		result: []int{1, 2},
		nums:   []int{3, 2, 4},
		target: 6,
	},
}
