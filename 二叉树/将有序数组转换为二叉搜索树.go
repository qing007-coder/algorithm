/*
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。

示例 1：
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
*/

package main

import "binary_tree/model"

func sortedArrayToBST(nums []int) *model.TreeNode {
	var build func(left, right int) *model.TreeNode

	build = func(left, right int) *model.TreeNode {
		if left > right {
			return nil
		}

		mid := (left + right) / 2

		node := &model.TreeNode{Val: nums[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)

		return node
	}

	return build(0, len(nums)-1)
}
