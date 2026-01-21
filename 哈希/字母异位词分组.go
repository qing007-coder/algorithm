/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

解释：

在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
示例 2:

输入: strs = [""]

输出: [[""]]

示例 3:

输入: strs = ["a"]

输出: [["a"]]



提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)

	fmt.Println("result:", result)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		// 把 s 排序，作为哈希表的 key
		tmp := []byte(s)
		slices.Sort(tmp)
		sortedS := string(tmp)
		// 排序后相同的字符串分到同一组
		m[sortedS] = append(m[sortedS], s)
	}

	// 哈希表的 value 保存分组后的结果
	return slices.Collect(maps.Values(m))
}
