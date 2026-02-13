/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：
输入：digits = "2"
输出：["a","b","c"]
*/

package main

func letterCombinations(digits string) []string {
	var phoneMap = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	res := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)

			path = path[:len(path)-1]
		}
	}
	backtrack(0)

	return res
}
