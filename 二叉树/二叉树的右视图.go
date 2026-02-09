/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]
解释：
示例 2：
输入：root = [1,2,3,4,null,null,null,5]
输出：[1,3,4,5]
解释：
示例 3：
输入：root = [1,null,3]
输出：[1,3]

示例 4：
输入：root = []
输出：[]
*/

package main

import "binary_tree/model"

func rightSideView(root *model.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*model.TreeNode{root}
	res := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i+1 == size {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
