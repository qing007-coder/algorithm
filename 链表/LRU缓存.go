/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]
解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
*/

package main

type Node struct {
	Key, Value int
	Next, Prev *Node
}

type LRUCache struct {
	keyToNode map[int]*Node
	dummy     *Node
	capacity  int
}

func Constructor(capacity int) LRUCache {
	dummy := new(Node)
	dummy.Next = dummy
	dummy.Prev = dummy

	return LRUCache{
		keyToNode: make(map[int]*Node),
		dummy:     dummy,
		capacity:  capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	node := l.getNode(key)
	if node != nil {
		return node.Value
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	node := l.getNode(key)
	if node != nil {
		node.Value = value
		l.pushFront(node.Key)
		return
	}

	newNode := &Node{key, value, nil, nil}
	l.keyToNode[key] = newNode
	if len(l.keyToNode) > l.capacity {
		delete(l.keyToNode, l.dummy.Prev.Key)
		l.dummy.Prev = l.dummy.Prev.Prev
		l.dummy.Prev.Next = l.dummy
	}

	newNode.Next = l.dummy.Next
	newNode.Prev = l.dummy
	l.dummy.Next.Prev = newNode
	l.dummy.Next = newNode
}

func (l *LRUCache) remove(key int) {
	node := l.keyToNode[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRUCache) getNode(key int) *Node {
	node := l.keyToNode[key]
	if node == nil {
		return nil
	}

	l.pushFront(key)
	return node
}

func (l *LRUCache) pushFront(key int) {
	node := l.keyToNode[key]
	if node == nil {
		return
	}

	l.remove(key)
	node.Next = l.dummy.Next
	node.Prev = l.dummy

	l.dummy.Next.Prev = node
	l.dummy.Next = node
}
