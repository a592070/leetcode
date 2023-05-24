package main

/*
https://leetcode.com/problems/implement-queue-using-stacks/
*/
func main() {

}

type MyQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructor() MyQueue {
	q := new(MyQueue)
	q.stack1 = NewStack()
	q.stack2 = NewStack()
	return *q
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stack2.IsEmpty() {
		for !this.stack1.IsEmpty() {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
	}

	return this.stack2.Pop().(int)
}

func (this *MyQueue) Peek() int {
	if this.stack2.IsEmpty() {
		for !this.stack1.IsEmpty() {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
	}
	return this.stack2.Peek().(int)
}

func (this *MyQueue) Empty() bool {
	return this.stack1.IsEmpty() && this.stack2.IsEmpty()
}

type Stack struct {
	list []interface{}
}

func NewStack() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0, 8)
	return s
}
func (s *Stack) Size() int {
	return len(s.list)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(element interface{}) {
	s.list = append(s.list, element)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[s.Size()-1]
	s.list = s.list[:s.Size()-1]
	return tmp
}
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[s.Size()-1]
	return tmp
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
