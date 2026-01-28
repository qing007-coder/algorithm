/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

提示：
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1

进阶：
一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix)

	fmt.Println("result:", matrix)
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 记录第一行是否包含 0
	firstRowHasZero := slices.Contains(matrix[0], 0)

	// 记录第一列是否包含 0
	firstColHasZero := false
	for _, row := range matrix {
		if row[0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// 用第一列 matrix[i][0] 保存 rowHasZero[i]
	// 用第一行 matrix[0][j] 保存 colHasZero[j]
	for i := 1; i < m; i++ { // 无需遍历第一行，如果 matrix[0][j] 本身是 0，那么相当于 colHasZero[j] 已经是 true
		for j := 1; j < n; j++ { // 无需遍历第一列，如果 matrix[i][0] 本身是 0，那么相当于 rowHasZero[i] 已经是 true
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 相当于 rowHasZero[i] = true
				matrix[0][j] = 0 // 相当于 colHasZero[j] = true
			}
		}
	}

	for i := 1; i < m; i++ { // 跳过第一行，留到最后修改
		for j := 1; j < n; j++ { // 跳过第一列，留到最后修改
			if matrix[i][0] == 0 || matrix[0][j] == 0 { // i 行或 j 列有 0
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一列一开始就包含 0，那么把第一列全变成 0
	if firstColHasZero {
		for _, row := range matrix {
			row[0] = 0
		}
	}

	// 如果第一行一开始就包含 0，那么把第一行全变成 0
	if firstRowHasZero {
		clear(matrix[0])
	}
}
