package main

/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{Val: head.Val}
	end := res
	for {
		if head.Next == nil {
			return res
		}
		if head.Val == head.Next.Val {
			head = head.Next
		} else {
			end.Next = &ListNode{Val: head.Next.Val}
			end = end.Next
			head = head.Next

		}

	}
}
