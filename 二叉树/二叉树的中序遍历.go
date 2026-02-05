/*
二叉树的中序遍历

给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
*/

package main

import "binary_tree/model"

func main() {

}

func inorderTraversal(root *model.TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*model.TreeNode)
	dfs = func(node *model.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
