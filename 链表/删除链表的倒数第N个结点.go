/*
删除链表的倒数第N个结点

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
*/

package main

import (
	"linked_list/tools"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := tools.BuildList(arr)
	n := 2
	res := removeNthFromEnd(list, n)
	tools.PrintList(res)
}

func removeNthFromEnd(head *tools.ListNode, n int) *tools.ListNode {
	// 用这个哨兵的原因是这个哨兵可以停在被删节点的前一个
	tempHead := &tools.ListNode{
		Next: head,
	}

	slow, fast := tempHead, tempHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return tempHead.Next
}
