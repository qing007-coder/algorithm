/*
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""。

测试用例保证答案唯一。

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	result := minWindow(s, t)
	fmt.Println(result)
}

func minWindow(s string, t string) string {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	left, valid := 0, 0
	start := 0
	minLen := len(s) + 1

	for right := 0; right < len(s); right++ {
		c := s[right]

		if _, ok := tMap[c]; ok {
			sMap[c]++
			if sMap[c] == tMap[c] {
				valid++
			}
		}

		for valid == len(tMap) {
			if right-left+1 < minLen {
				start = left
				minLen = right - left + 1
			}

			d := s[left]
			left++

			if _, ok := tMap[d]; ok {
				if sMap[d] == tMap[d] {
					valid--
				}
				sMap[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
