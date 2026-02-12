/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]
*/

package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	res := subsets(list)
	fmt.Println("res:", res)
}

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	var backtrack func(int)
	backtrack = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])

			backtrack(i + 1)

			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
