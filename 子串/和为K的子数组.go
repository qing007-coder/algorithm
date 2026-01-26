/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
*/

package main

import "fmt"

func main() {
	arr := []int{1, 1, 1}

	result := subarraySum(arr, 2)
	fmt.Println("result:", result)
}

func subarraySum(nums []int, k int) int {
	prefixMap := make(map[int]int)
	var result, total int

	prefixMap[0] = 1

	for _, num := range nums {
		total += num
		result += prefixMap[total-k]
		prefixMap[total]++
	}

	return result
}
