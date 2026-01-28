/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := spiralOrder(matrix)

	fmt.Println("result:", res)
}

func spiralOrder(matrix [][]int) []int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
	res := make([]int, 0)

	x, y, direction := 0, 0, 0
	for len(res) < len(matrix)*len(matrix[0]) { // 一共走 mn 步
		res = append(res, matrix[x][y])
		matrix[x][y] = math.MaxInt // 标记，表示已经访问过（已经加入答案）

		nextX, nextY := x+dirs[direction][0], y+dirs[direction][1] //
		// 如果 (x, y) 出界或者已经访问过
		if nextX < 0 || nextY < 0 || nextX >= len(matrix) || nextY >= len(matrix[0]) || matrix[nextX][nextY] == math.MaxInt {
			direction = (direction + 1) % 4 // 右转 90°
		}

		x += dirs[direction][0]
		y += dirs[direction][1] // 走一步
	}

	return res
}
