/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
两两交换链表中的节点

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]
*/

package main

import "linked_list/tools"

func main() {
	arr1 := []int{1, 2, 3, 4}
	list := tools.BuildList(arr1)

	res := swapPairs(list)
	tools.PrintList(res)
}

func swapPairs(head *tools.ListNode) *tools.ListNode {
	dummy := &tools.ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		p, q := prev.Next, prev.Next.Next

		p.Next = q.Next
		q.Next = p
		prev.Next = q // prev后面就该是q 因为prev后面原来是p 交换完就是应该是q

		prev = p // 这里是p 是因为p的后面是下一次被反转的地方 q后面是p 所以不是q
	}

	return dummy.Next
}
