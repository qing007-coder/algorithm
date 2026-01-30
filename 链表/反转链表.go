/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]
*/

package main

import (
	"fmt"
	"linked_list/tools"
)

func main() {
	head := tools.BuildList([]int{1, 2, 3, 4, 5})
	res := reverseList(head)

	fmt.Println("result:")
	tools.PrintList(res)
}

// 头插法
func reverseList(head *tools.ListNode) *tools.ListNode {
	tempHead := &tools.ListNode{}
	p := head

	for p != nil {
		q := p.Next
		p.Next = tempHead.Next
		tempHead.Next = p

		p = q
	}

	return tempHead.Next
}
