/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
示例 3：

输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	result := merge(intervals)
	fmt.Println("result:", result)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}

		if res[len(res)-1][1] >= interval[0] {
			if res[len(res)-1][1] > interval[1] { // 有可能interval被上一个区间包进去 [[1,4],[2,3]]
				continue
			}
			res[len(res)-1][1] = interval[1] // 没有包进去 就合并
		} else {
			res = append(res, interval)
		}
	}

	return res
}
