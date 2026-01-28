/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(arr, k)
	fmt.Println("result:", arr)
}

func rotate(arr []int, k int) {
	if len(arr) == 0 {
		return
	}

	k = k % len(arr)

	reverse(arr, 0, len(arr)-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
}

func reverse(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func rotate1(nums []int, k int) {
	res := make([]int, len(nums))

	for i, num := range nums {
		res[(i+k)%len(nums)] = num
	}

	copy(nums, res)
}
