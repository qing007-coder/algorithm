/*
示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：
输入：s = "a"
输出：[["a"]]
*/

package main



func partition(s string) [][]string {
	res := make([][]string, 0)
    path := make([]string, 0)

	var dfs func(string)
	dfs = func(s string) {
		if len(s) == 0 {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return 
		}
		
		for i := 1; i <= len(s); i++ {
			substr := s[:i]
			if isPalindrome(substr) {
				path = append(path, s[:i])
				dfs(s[i:])
				path = path[:len(path) - 1]
			}
		}
	}

	dfs(s)

	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}