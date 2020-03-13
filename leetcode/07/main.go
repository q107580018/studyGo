/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
package main

import "fmt"

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	re := addTwoNumbers(l1, l2)
	for re != nil {
		fmt.Println(re.Val)
		re = re.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	HeadNode := new(ListNode)
	pre := HeadNode
	i := 0 // 用于进位
	for {
		if l1 != nil {
			pre.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			pre.Val += l2.Val
			l2 = l2.Next
		}
		pre.Val += i
		if pre.Val > 9 {
			pre.Val = pre.Val % 10
			i = 1
		} else {
			i = 0
		}
		if l1 == nil && l2 == nil {
			if i == 1 {
				pre.Next = &ListNode{Val: 1}
			}
			return HeadNode
		}
		pre.Next = new(ListNode)
		pre = pre.Next

	}
}
