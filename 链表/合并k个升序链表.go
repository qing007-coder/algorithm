/*
合并k个升序链表

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]
*/

package main

import "linked_list/tools"

func main() {
	arr1 := []int{1, 4, 5}
	arr2 := []int{1, 3, 4}
	arr3 := []int{2, 6}

	list1 := tools.BuildList(arr1)
	list2 := tools.BuildList(arr2)
	list3 := tools.BuildList(arr3)

	list := []*tools.ListNode{list1, list2, list3}
	res := mergeKLists(list)
	tools.PrintList(res)
}

func mergeKLists(lists []*tools.ListNode) *tools.ListNode {
	if len(lists) == 0 {
		return nil
	}

	for internal := 1; internal < len(lists); internal *= 2 {
		for i := 0; i+internal < len(lists); i += internal * 2 {
			lists[i] = mergeTwoLists2(lists[i], lists[i+internal])
		}
	}

	return lists[0]
}

func mergeTwoLists2(list1 *tools.ListNode, list2 *tools.ListNode) *tools.ListNode {
	dummy := &tools.ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
