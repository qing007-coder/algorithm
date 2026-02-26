/*
寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

package main

import "math"

// 太难了 放弃了
func findMedianSortedArrays(a, b []int) float64 {
    if len(a) > len(b) {
        a, b = b, a // 保证下面的 i 可以从 0 开始枚举
    }

    m, n := len(a), len(b)
    a = append([]int{math.MinInt}, append(a, math.MaxInt)...)
    b = append([]int{math.MinInt}, append(b, math.MaxInt)...)

    // 循环不变量：a[left] <= b[j+1]
    // 循环不变量：a[right] > b[j+1]
    left, right := 0, m+1
    for left+1 < right { // 开区间 (left, right) 不为空
        i := left + (right-left)/2
        j := (m+n+1)/2 - i
        if a[i] <= b[j+1] {
            left = i // 缩小二分区间为 (i, right)
        } else {
            right = i // 缩小二分区间为 (left, i)
        }
    }

    // 此时 left 等于 right-1
    // a[left] <= b[j+1] 且 a[right] > b[(j-1)+1] = b[j]，所以答案是 i=left
    i := left
    j := (m+n+1)/2 - i
    max1 := max(a[i], b[j])
    min2 := min(a[i+1], b[j+1])
    if (m+n)%2 > 0 {
        return float64(max1)
    }
    return float64(max1+min2) / 2
}
