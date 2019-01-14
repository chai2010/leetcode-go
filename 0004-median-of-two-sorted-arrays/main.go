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
	if len(nums1) == 0 && len(nums2) == 0 {
		return [2]int{}
	}

	if len(nums1) == 0 {
		return [2]int{
			nums2[(len(nums2)-1)/2],
			nums2[(len(nums2)-0)/2],
		}
	}
	if len(nums2) == 0 {
		return [2]int{
			nums1[(len(nums1)-1)/2],
			nums1[(len(nums1)-0)/2],
		}
	}

	return [2]int{} // TODO
}
