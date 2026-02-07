/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。

示例 1：
输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
*/

package main

import "binary_tree/model"

// 就是一个中序遍历
func kthSmallest(root *model.TreeNode, k int) int {
	res := 0
	var dfs func(node *model.TreeNode)
	dfs = func(node *model.TreeNode) {
		if node == nil || k == 0 {
			return
		}

		dfs(node.Left)

		if k == 0 {
			return
		}

		k--

		if k == 0 {
			res = node.Val
			return
		}

		dfs(node.Right)
	}
	dfs(root)
	return res
}
