/*
你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]
*/

package main

import "linked_list/tools"

func main() {
	arr := []int{-1, 5, 3, 4, 0}
	list := tools.BuildList(arr)

	res := sortList(list)
	tools.PrintList(res)
}

func sortList(head *tools.ListNode) *tools.ListNode {
	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}

	dummy := &tools.ListNode{Next: head}
	for step := 1; step < length; step *= 2 {
		newHead := dummy
		cur := newHead.Next
		for cur != nil {
			head1 := cur
			head2 := splitList(head1, step)
			cur = splitList(head2, step)

			tempHead, tail := mergeTwoLists1(head1, head2)
			newHead.Next = tempHead
			newHead = tail
		}
	}

	return dummy.Next
}

func splitList(head *tools.ListNode, size int) *tools.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	cur := head

	for i := 0; i < size-1 && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil || cur.Next == nil {
		return nil
	}

	newHead := cur.Next
	cur.Next = nil
	return newHead
}

func mergeTwoLists1(l1 *tools.ListNode, l2 *tools.ListNode) (*tools.ListNode, *tools.ListNode) {
	dummy := &tools.ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	return dummy.Next, cur
}
