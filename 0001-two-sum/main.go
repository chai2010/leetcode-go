// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
	)

	fmt.Println(twoSumV0(nums, target))
	// [0, 1]
}

func twoSumV0(nums []int, target int) []int {
	for i, vi := range nums {
		for j, vj := range nums[i+1:] {
			if vi+vj == target {
				return []int{i, i + 1 + j}
			}
		}
	}
	return nil
}
