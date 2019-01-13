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
	fmt.Println(twoSumV1(nums, target))
	fmt.Println(twoSumV2(nums, target))
	fmt.Println(twoSumV3(nums, target))
	fmt.Println(twoSumV4(nums, target))

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

func twoSumV1(nums []int, target int) []int {
	var idxMap = map[int][]int{}
	for i, v := range nums {
		idxMap[v] = append(idxMap[v], i)
	}

	for i, v := range nums {
		if v+v == target {
			if len(idxMap[v]) == 2 {
				return idxMap[v]
			}
			continue
		}

		if _, ok := idxMap[target-v]; ok {
			return []int{i, idxMap[target-v][0]}
		}
	}

	return nil
}

func twoSumV2(nums []int, target int) []int {
	var idxMap = make(map[int]int32)
	for i, v := range nums {
		idxMap[v] = int32(i)
	}

	for i, v := range nums {
		if j, ok := idxMap[target-v]; ok && int(j) > i {
			return []int{i, int(j)}
		}
	}

	return nil
}

func twoSumV3(nums []int, target int) []int {
	var idxMap = make(map[int]int32)
	for i, v := range nums {
		if j, ok := idxMap[target-v]; ok {
			return []int{int(j), i}
		}
		idxMap[v] = int32(i)
	}
	return nil
}

func twoSumV4(nums []int, target int) []int {
	mid := len(nums) / 2
	A, B := nums[:mid], nums[mid:]

	result := make(chan []int, 3)

	// AB
	go func() {
		for i, vi := range A {
			for j, vj := range B {
				if vi+vj == target {
					if i != mid+j {
						result <- []int{i, mid + j}
						return
					}
				}
			}
		}
	}()

	// AA
	go func() {
		for i, vi := range A {
			for j, vj := range A[i+1:] {
				if vi+vj == target {
					result <- []int{i, i + 1 + j}
					return
				}
			}
		}
	}()

	// BB
	go func() {
		for i, vi := range B {
			for j, vj := range B[i+1:] {
				if vi+vj == target {
					result <- []int{mid + i, mid + i + 1 + j}
					return
				}
			}
		}
	}()

	return <-result
}

// A0A1

//A0A0
//A0A1

//A1A0
//A1A1

//AA
//AB
//BB
