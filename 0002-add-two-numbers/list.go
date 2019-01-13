// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v ...int) *ListNode {
	if len(v) == 0 {
		return nil
	}
	return &ListNode{
		Val:  v[0],
		Next: NewListNode(v[1:]...),
	}
}

func (p *ListNode) Slice() []int {
	if p == nil {
		return nil
	}
	if p.Next == nil {
		return []int{p.Val}
	}
	return append([]int{p.Val},
		p.Next.Slice()...,
	)
}

func (p *ListNode) String() string {
	return fmt.Sprintf("%v", p.Slice())
}
