/*
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。

示例 1：
输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。

示例 2：
输入：root = [1,2]
输出：1
*/

package main

import "binary_tree/model"

func diameterOfBinaryTree(root *model.TreeNode) int {
	maxDiameter := 0

	// 定义一个计算高度的辅助函数
	var depth func(*model.TreeNode) int
	depth = func(node *model.TreeNode) int {
		if node == nil {
			return 0 // 空节点高度为 0
		}

		// 递归计算左右子树的高度 这里是个后序遍历
		left := depth(node.Left)
		right := depth(node.Right)

		// 【关键点】更新全局最大直径
		// 经过当前节点的路径长度 = 左子树高度 + 右子树高度
		if left+right > maxDiameter {
			maxDiameter = left + right
		}

		// 返回当前节点的高度，给上一层节点使用
		// 当前高度 = max(左, 右) + 1
		return max(left, right) + 1
	}

	depth(root)
	return maxDiameter
}
