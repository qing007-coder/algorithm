/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/

package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"

	result := findAnagrams(s, p)
	fmt.Println("result:", result)
}

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	var pFrequency, sFrequency [26]int
	for _, str := range p {
		pFrequency[str-'a']++
	}

	for end, str := range s {
		sFrequency[str-'a']++
		start := end - len(p) + 1
		if start < 0 {
			continue
		}

		if pFrequency == sFrequency {
			result = append(result, start)
		}

		sFrequency[s[start]-'a']--
	}

	return result
}
