/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)

	fmt.Println(nums)
}

// 把 nums 当作栈
func moveZeroes1(nums []int) {
	stackSize := 0 // 另一种写法见【Go 写法二】
	for _, x := range nums {
		if x != 0 {
			nums[stackSize] = x // 把 x 入栈
			stackSize++
		}
	}
	clear(nums[stackSize:]) // nums[stackSize:] 所有元素置为 0
}

// 方法二：双指针+交换元素
// 方法一在最坏情况下（nums 全为 0），需要遍历 nums 两次。能否做到一次遍历？
//
// 核心思路
// 把 0 视作空位。我们要把所有非零元素都移到数组左边的空位上，并保证非零元素的顺序不变。
//
// 例如 nums=[0,0,1,2]，把 1 放到最左边的空位上，数组变成 。注意 1 移动过去后，在原来 1 的位置又产生了一个新的空位。也就是说，我们交换了 nums[0]=0 和 nums[2]=1 这两个数。
//
// 为了保证非零元素的顺序不变，我们需要维护最左边的空位的位置（下标）。
func moveZeroes2(nums []int) {
	p := 0                     // p指针作为指向左边的不为0的数的栈顶
	for q, num := range nums { // q指针是为了继续往下遍历数组
		if num != 0 {
			nums[p], nums[q] = num, nums[p]
			p++
		}
	}
}
