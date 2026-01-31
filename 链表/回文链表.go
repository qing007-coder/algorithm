/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false
*/

package main

import (
	"fmt"
	"linked_list/tools"
)

func main() {
	head := tools.BuildList([]int{1, 2, 2, 1})
	res := isPalindrome(head)
	fmt.Println("result:", res)
}

func isPalindrome(head *tools.ListNode) bool {
	middle := middleNode(head)
	middleStart := reverse(middle)
	for middleStart != nil && head != nil {
		if head.Val != middleStart.Val {
			return false
		}
		head = head.Next
		middleStart = middleStart.Next
	}

	return true
}

func middleNode(head *tools.ListNode) *tools.ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *tools.ListNode) *tools.ListNode {
	p := head
	temp := &tools.ListNode{}
	for p != nil {
		q := p.Next
		p.Next = temp.Next
		temp.Next = p
		p = q
	}

	return temp.Next
}
