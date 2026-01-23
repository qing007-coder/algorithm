/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcabcbb"
	result := lengthOfLongestSubstring(str)
	fmt.Println("result:", result)
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	arr := strings.Split(s, "")
	dict := make(map[string]struct{})
	start := 0

	for end := 0; end < len(arr); end++ {
		// 把子串重复的 包括之前的 都删干净
		for _, ok := dict[arr[end]]; ok; {
			delete(dict, arr[start])
			start++
			_, ok = dict[arr[end]]
		}

		dict[arr[end]] = struct{}{}
		res = max(len(dict), res)
	}

	return res
}
