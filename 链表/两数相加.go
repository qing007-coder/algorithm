/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/

package main

import (
	"linked_list/tools"
)

func main() {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	list1 := tools.BuildList(arr1)
	list2 := tools.BuildList(arr2)

	result := addTwoNumbers(list1, list2)
	tools.PrintList(result)
}

func addTwoNumbers(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {
	p := l1
	q := l2
	carry := 0
	head := &tools.ListNode{}
	temp := head

	for p != nil && q != nil {
		value := q.Val + p.Val + carry
		carry = value / 10
		value %= 10
		p.Val = value
		temp.Next = p
		temp = temp.Next
		p = p.Next
		q = q.Next
	}

	for p != nil {
		value := p.Val + carry
		carry = value / 10
		value %= 10
		p.Val = value
		temp.Next = p
		temp = temp.Next
		p = p.Next
	}

	for q != nil {
		value := q.Val + carry
		carry = value / 10
		value %= 10
		q.Val = value
		temp.Next = q
		temp = temp.Next
		q = q.Next
	}

	if carry != 0 {
		temp.Next = &tools.ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}
