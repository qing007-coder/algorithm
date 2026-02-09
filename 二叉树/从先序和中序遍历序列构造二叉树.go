/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]
*/

package main

import "binary_tree/model"

func buildTree(preorder []int, inorder []int) *model.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var build func(preL, preR, inL, inR int) *model.TreeNode
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}

	build = func(preL, preR, inL, inR int) *model.TreeNode {
		if preL > preR {
			return nil
		}

		value := preorder[preL]
		root := &model.TreeNode{Val: value}
		leftSize := inMap[value] - inL

		root.Left = build(preL+1, preL+leftSize, inL, inMap[value]-1)
		root.Right = build(preL+leftSize+1, preR, inMap[value]+1, inR)

		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
