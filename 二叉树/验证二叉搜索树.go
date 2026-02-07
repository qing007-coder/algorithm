/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 严格小于 当前节点的数。
节点的右子树只包含 严格大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]
输出：true

示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/

package main

import (
	"binary_tree/model"
	"math"
)

func isValidBST(root *model.TreeNode) bool {
	var verify func(root *model.TreeNode, maxInt, minInt int) bool

	verify = func(root *model.TreeNode, maxInt, minInt int) bool {
		if root == nil {
			return true
		}

		if root.Val >= maxInt || root.Val <= minInt {
			return false
		}

		return verify(root.Left, root.Val, minInt) && verify(root.Right, maxInt, root.Val)
	}

	return verify(root, math.MaxInt, math.MinInt)
}
