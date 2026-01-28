/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。

示例 2：
输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。

示例 3：
输入：nums = [7,8,9,11,12]
输出：1
解释：最小的正数 1 没有出现。
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, -1, 1}
	result := firstMissingPositive(arr)
	fmt.Println("result:", result)
}

func firstMissingPositive(nums []int) int {
	for i, _ := range nums {
		for nums[i] > 0 && (nums[i]-1) != i && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, num := range nums {
		if i != (num - 1) {
			return i + 1
		}
	}

	return len(nums) + 1
}
