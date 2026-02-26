/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：
输入：n = 1
输出：[["Q"]]
*/

package main

import "fmt"

func main() {
	n := 4

	res := solveNQueens(n)
	fmt.Println("res:", res)
}

func solveNQueens(n int) [][]string {
    res := make([][]string, 0)
	colLocation := make([]int, n)

	columnSignal := make(map[int]bool)
	diagonalSignal1 := make(map[int]bool)
	diagonalSignal2 := make(map[int]bool)

	var backtrack func(int)
	backtrack =func(row int) {
		if row == n {
			res = append(res, build(colLocation))
			return 
		}


		for col := 0; col < n; col++ {
			if columnSignal[col] || diagonalSignal1[row-col] || diagonalSignal2[row+col] {
				continue
			}

			colLocation[row] = col 
			columnSignal[col] = true
			diagonalSignal1[row - col] = true
			diagonalSignal2[row + col] = true

			backtrack(row + 1)

			delete(columnSignal, col)
			delete(diagonalSignal1, row - col)
			delete(diagonalSignal2, row + col)
		}
	}

	backtrack(0)

	return res
}

func build(colLocation []int) []string {
	res := make([]string, 0)

	for i := 0; i < len(colLocation); i++ {
		location := colLocation[i]
		col := make([]byte, 0)
		for j := 0; j < len(colLocation); j++ {
			if j == location {
				col = append(col, 'Q')
				continue
			}

			col = append(col, '.')
		}

		res = append(res, string(col))
	}

	return res
}