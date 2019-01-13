# 2. 两数相加

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

## 示例

1. 输入：`(2 -> 4 -> 3) + (5 -> 6 -> 4)`
1. 输出：`7 -> 0 -> 8`
1. 原因：`342 + 465 = 807`

## 准备工作

```go
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
```

为了便于测试，为链接实现`NewListNode`构造函数和`Slice`方法，用户和切片的双向转换。

## 方案V0（28ms）

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		p, q, carry = l1, l2, 0

		result        = new(ListNode)
		p_result_next = &result
	)

	for p != nil && q != nil {
		(*p_result_next).Val = (p.Val + q.Val + carry) % 10
		(*p_result_next).Next = new(ListNode)
		p_result_next = &((*p_result_next).Next)

		carry = (p.Val + q.Val + carry) / 10
		p, q = p.Next, q.Next
	}

	for p != nil {
		(*p_result_next).Val = (p.Val + carry) % 10
		(*p_result_next).Next = new(ListNode)
		p_result_next = &((*p_result_next).Next)

		carry = (p.Val + carry) / 10
		p = p.Next
	}
	for q != nil {
		(*p_result_next).Val = (q.Val + carry) % 10
		(*p_result_next).Next = new(ListNode)
		p_result_next = &((*p_result_next).Next)

		carry = (q.Val + carry) / 10
		q = q.Next
	}

	if carry > 0 {
		(*p_result_next).Val = 1
	} else {
		(*p_result_next) = nil
	}

	return result
}
```

解决的思路是合并两个单向链表，多余的部分追加到尾部。单向链表比较绕的部分是在尾部时需要将前一个节点的Next指针设置为nil，很多人会为前一个节点单独保存一个指针。这里有一个技巧是保存前一个节点Next指针的地址，这样就将单项链表的开头节点/普通节点/尾部节点都同一化处理了。

目前这个方案耗时28ms，最优的方案大约在20ms附近。鉴于Go语言的特性，28ms应该是Go语言版本中较优的方案了。
