/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]
*/

package main

import "fmt"

func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	result := maxSlidingWindow(arr, 3)
	fmt.Println("result:", result)
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	queue := make([]int, 0) // 这里存的是索引 不然不好判断这个数是不是出队列了

	for i, num := range nums {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < num { // 判断队尾是不是比当前数小
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		left := i - k + 1
		if left < 0 {
			continue
		}

		if left > queue[0] {
			queue = queue[1:]
		}

		result = append(result, nums[queue[0]])
	}

	return result
}
