package main

import "fmt"

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	head := mergeTwoLists(l1, l2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var headNode = &ListNode{Val: -1}
	pre := headNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}
		pre = pre.Next
	}
	if l1 == nil {
		pre.Next = l2
	} else if l2 == nil {
		pre.Next = l1
	}
	return headNode.Next
}
