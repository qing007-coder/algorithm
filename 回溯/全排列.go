/*
46. 全排列
中等
相关标签
premium lock icon
相关企业
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：
输入：nums = [1]
输出：[[1]]
*/

package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	res := permute(list)
	fmt.Println("result", res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, 0)

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}
