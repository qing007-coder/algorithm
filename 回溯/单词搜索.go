/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
输出：true
示例 2：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"
输出：true
示例 3：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"
输出：false
*/

package main

import "fmt"

func main() {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}

	word := "ABCCED"
	res := exist(board, word)

	fmt.Println("res:", res)
}

func exist(board [][]byte, word string) bool {
    visited := make([][]bool, len(board))
	for i, _ := range board {
		visited[i] = make([]bool, len(board[0]))
	}

	var dfs func(x, y, k int) bool
	dfs = func(x, y, k int) bool {
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
			return false
		}

		if visited[x][y] || board[x][y] != word[k] {
			return false
		}

		if k == len(word) - 1 {
			return true
		}

		visited[x][y] = true
		if dfs(x - 1, y, k + 1) || dfs(x + 1, y, k + 1) || dfs(x, y - 1, k + 1) || dfs(x, y + 1, k + 1) {
			return true
		}

		visited[x][y] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}