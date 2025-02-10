package main

func main() {

}

// MinStack Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// Implement the MinStack class:
//
// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.
type MinStack struct {
	dataStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		dataStack: []int{},
		minStack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.dataStack = append(this.dataStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		min := this.minStack[len(this.minStack)-1]
		if min > val {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, min)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.dataStack) == 0 {
		return
	}
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.dataStack) == 0 {
		return -1
	}
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return -1
	}
	return this.minStack[len(this.minStack)-1]
}
