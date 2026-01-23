/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9
*/

package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap2(arr)

	fmt.Println("result:", res)
}

// 前后缀分解
func trap1(height []int) int {
	// res 用来累计最终能接到的雨水总量
	res := 0
	length := len(height)

	// preMax[i] 表示：从左到右遍历时，
	// 下标 i 及其左边的最大柱子高度
	preMax := make([]int, length)

	// sufMax[i] 表示：从右到左遍历时，
	// 下标 i 及其右边的最大柱子高度
	sufMax := make([]int, length)

	// prefix：当前遍历到位置时，左侧的最大高度
	// suffix：当前遍历到位置时，右侧的最大高度
	prefix, suffix := 0, 0

	// 一次遍历同时构建前缀最大值和后缀最大值
	for i := 0; i < length; i++ {
		// 更新左侧最大高度
		if prefix < height[i] {
			prefix = height[i]
		}
		preMax[i] = prefix

		// 更新右侧最大高度（从数组末尾向前）
		if suffix < height[length-i-1] {
			suffix = height[length-i-1]
		}
		sufMax[length-1-i] = suffix
	}

	// 计算每个位置能够接到的雨水
	// 水位由左右最大高度中的较小值决定
	for i := 0; i < length; i++ {
		res += min(preMax[i], sufMax[i]) - height[i]
	}

	return res
}

// 相向双指针
func trap2(height []int) int {
	// length 表示柱子数量
	length := len(height)

	// res 用来累计最终接到的雨水总量
	res := 0

	// preMax：从左侧遍历到当前位置时，左边出现过的最大高度
	// sufMax：从右侧遍历到当前位置时，右边出现过的最大高度
	preMax, sufMax := 0, 0

	// left、right 为左右双指针
	left, right := 0, length-1

	// 当左右指针相遇时，所有位置都已处理完
	for left < right {
		// 更新当前左右两侧的最大高度
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		// 谁的最大高度更小，谁这一侧的雨水就已经可以确定
		if preMax > sufMax {
			// 右侧的最大高度更小
			// 当前 right 位置能接的水量只由 sufMax 决定
			res += sufMax - height[right]
			// 右指针向左收缩
			right--
		} else {
			// 左侧的最大高度更小（或相等）
			// 当前 left 位置能接的水量只由 preMax 决定
			res += preMax - height[left]
			// 左指针向右收缩
			left++
		}
	}

	return res
}
