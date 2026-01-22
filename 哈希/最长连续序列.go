/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
示例 3：

输入：nums = [1,0,1,2]
输出：3
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutive(arr)
	fmt.Println("result:", result)
}

func longestConsecutive(nums []int) (ans int) {
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true // 把 nums 转成哈希集合
	}

	for x := range has { // 遍历哈希集合
		if has[x-1] { // 如果 x 不是序列的起点，直接跳过
			continue
		}
		// x 是序列的起点
		y := x + 1
		for has[y] { // 不断查找下一个数是否在哈希集合中
			y++
		}
		// 循环结束后，y-1 是最后一个在哈希集合中的数
		ans = max(ans, y-x) // 从 x 到 y-1 一共 y-x 个数
	}
	return
}
