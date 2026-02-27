/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：
输入： heights = [2,4]
输出： 4
*/

package main

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
    stack := make([]int, 0)
	res := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[stack[len(stack) - 1]] > heights[i] {
            height := heights[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
			width := i - stack[len(stack) - 1] - 1

			area := height * width
			res = max(res, area)
		}

		stack = append(stack, i)
	}

	return res
}