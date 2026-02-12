/*
	前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/
package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{children: [26]*Trie{}}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}

		node = node.children[idx]
	}

	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchWithPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchWithPrefix(prefix)
	return node != nil
}

func (this *Trie) searchWithPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		fmt.Println(node)
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}

		node = node.children[idx]
	}

	return node
}

func main() {
	trie := Constructor()
	trie.Insert("abc")
	fmt.Println(trie.Search("abc"))
}
