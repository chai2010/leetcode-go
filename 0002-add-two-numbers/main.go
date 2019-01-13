// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := NewListNode(5, 6, 4)

	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)

	sum := addTwoNumbers(l1, l2)
	fmt.Println("l1+l2:", sum)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		p, q, carry = l1, l2, 0

		result           = new(ListNode)
		result_prev_next = &result
	)

	for p != nil && q != nil {
		(*result_prev_next).Val = (p.Val + q.Val + carry) % 10
		(*result_prev_next).Next = new(ListNode)
		result_prev_next = &((*result_prev_next).Next)

		carry = (p.Val + q.Val + carry) / 10
		p, q = p.Next, q.Next
	}

	for p != nil {
		(*result_prev_next).Val = (p.Val + carry) % 10
		(*result_prev_next).Next = new(ListNode)
		result_prev_next = &((*result_prev_next).Next)

		carry = (p.Val + carry) / 10
		p = p.Next
	}
	for q != nil {
		(*result_prev_next).Val = (q.Val + carry) % 10
		(*result_prev_next).Next = new(ListNode)
		result_prev_next = &((*result_prev_next).Next)

		carry = (q.Val + carry) / 10
		q = q.Next
	}

	if carry > 0 {
		(*result_prev_next).Val = 1
	} else {
		(*result_prev_next) = nil
	}

	return result
}
