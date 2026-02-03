/*
K个一组翻转链表

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
*/

package main

import "linked_list/tools"

func main() {
	k := 2
	arr := []int{1, 2, 3, 4, 5}
	list := tools.BuildList(arr)

	res := reverseKGroup(list, k)

	tools.PrintList(res)
}

func reverseKGroup(head *tools.ListNode, k int) *tools.ListNode {
	dummy := &tools.ListNode{Next: head}
	prev := dummy
	for {
		tail := prev.Next
		for i := 0; i < k; i++ {
			if tail == nil {
				return dummy.Next
			}

			tail = tail.Next
		}

		oldHead := prev.Next
		newHead := reverse1(oldHead, tail)
		oldHead.Next = tail

		prev.Next = newHead
		prev = oldHead
	}
}

func reverse1(head *tools.ListNode, tail *tools.ListNode) *tools.ListNode {
	dummy := &tools.ListNode{}
	cur := head

	for cur != tail {
		next := cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = next
	}
	tools.PrintList(dummy.Next)

	return dummy.Next
}
