/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]
*/

package main

import "binary_tree/model"

func flatten(root *model.TreeNode) {
	if root == nil {
		return
	}

	// 1. 递归展开左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 2. 保存左右子树
	left := root.Left
	right := root.Right

	// 3. 左子树放到右边
	root.Left = nil
	root.Right = left

	// 4. 找到当前右链表的尾部
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	// 5. 把原右子树接上
	cur.Right = right
}
