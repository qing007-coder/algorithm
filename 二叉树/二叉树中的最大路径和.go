/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

示例 2：
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
*/
package main

import (
	"binary_tree/model"
	"math"
)

func maxPathSum(root *model.TreeNode) int {
	// 初始化为最小整数
	maxSum := math.MinInt64

	// 定义递归函数
	var gain func(*model.TreeNode) int
	gain = func(node *model.TreeNode) int {
		if node == nil {
			return 0
		}

		// 1. 递归计算左右子树的最大贡献值
		// 如果贡献是负数，直接舍弃（取 0）
		leftGain := max(gain(node.Left), 0)
		rightGain := max(gain(node.Right), 0)

		// 2. 计算以当前节点为“顶点”的闭合路径和
		// 这一步是在更新全局最大值，它对应的是“左中右”这种拱形路径
		currentPathSum := node.Val + leftGain + rightGain
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		// 3. 返回该节点给父节点提供的“单边最大贡献”
		// 因为路径不能分叉，所以只能选左或者右其中一条较长的路径向上返回
		return node.Val + max(leftGain, rightGain)
	}

	gain(root)
	return maxSum
}
