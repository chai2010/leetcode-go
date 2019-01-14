// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	// Output:
	// 2
	// 2.5
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	v := SolutionV0(nums1, nums2)
	return (float64(v[0]) + float64(v[1])) / 2
}

func SolutionV0(nums1 []int, nums2 []int) [2]int {
	return [2]int{} // TODO
}
