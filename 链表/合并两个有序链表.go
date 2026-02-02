/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/

package main

import "linked_list/tools"

func main() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	list1 := tools.BuildList(arr1)
	list2 := tools.BuildList(arr2)

	result := mergeTwoLists(list1, list2)
	tools.PrintList(result)
}

func mergeTwoLists(list1 *tools.ListNode, list2 *tools.ListNode) *tools.ListNode {
	p := list1
	q := list2
	head := &tools.ListNode{}
	temp := head

	for p != nil && q != nil {
		if p.Val < q.Val {
			temp.Next = p
			p = p.Next
		} else {
			temp.Next = q
			q = q.Next
		}
		temp = temp.Next
	}

	for p != nil {
		temp.Next = p
		p = p.Next
		temp = temp.Next
	}

	for q != nil {
		temp.Next = q
		q = q.Next
		temp = temp.Next
	}

	return head.Next
}
