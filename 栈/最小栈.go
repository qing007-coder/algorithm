/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
 

示例 1:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

package main

type MinStack struct {
	data []int
    minStack []int
}


func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		minStack: make([]int, 0),
	}
}


func (this *MinStack) Push(val int)  {
    this.data = append(this.data, val)
	
	if len(this.minStack) == 0 || this.minStack[len(this.minStack) - 1] >= val {
		this.minStack = append(this.minStack, val)
	}
}


func (this *MinStack) Pop()  {
    top := this.data[len(this.data) - 1]
	this.data = this.data[:len(this.data) - 1]

	if top == this.minStack[len(this.minStack) - 1] {
		this.minStack = this.minStack[:len(this.minStack) - 1]
	}
}


func (this *MinStack) Top() int {
    return this.data[len(this.data) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */