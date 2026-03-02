/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:
输入: numRows = 1
输出: [[1]]
*/

package main

func generate(numRows int) [][]int {
    lastRow := []int{1}
	res := make([][]int, 0)

	res = append(res, lastRow)
	for i := 0; i < numRows - 1; i++ {
		row := make([]int, 0)
		lastRow = append(lastRow, 0)
		lastRow = append([]int{0}, lastRow...)

		for left := 0; left < len(lastRow) - 1; left++ {
			right := left + 1
			row = append(row, lastRow[left] + lastRow[right])
		}
		lastRow = row

		res = append(res, row)
	}

	return res
}
