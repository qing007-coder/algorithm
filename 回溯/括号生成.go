/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]
*/

package main

import "fmt"

func main() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println("res:", res)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == n && right == n {
			tmp := make([]byte, len(path))
			copy(tmp, path)
			res = append(res, string(tmp))
			return
		}

		if left < n {
			path = append(path, '(')
			dfs(left + 1, right)
			path = path[:len(path) - 1]
		}

		if right < left {
			path = append(path, ')')
			dfs(left, right + 1)
			path = path[:len(path) - 1]
		}
	}

	dfs(0, 0)
	return res
}
