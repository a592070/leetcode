package main

import "fmt"

/**
ref.
https://github.com/mkirchner/linked-list-good-taste
*/

type Node struct {
	Val  int
	Next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) remove(index int) {
	var p **Node = &list.head
	for i := 0; i < index; i++ {
		p = &(*p).Next
	}
	*p = (*p).Next
}
func (list *LinkedList) removeTarget(target *Node) {
	var p **Node = &list.head
	for *p != target {
		p = &(*p).Next
	}
	*p = target.Next
}
func (list *LinkedList) removeTarget2(target *Node) {
	var p *Node = list.head
	for p != target {
		p = p.Next
	}
	// replace current node address with next
	*p = *target.Next
}
func (list *LinkedList) println() {
	p := list.head
	for p != nil {
		fmt.Println(p)
		p = p.Next
	}
}

func generate(size int) *LinkedList {
	head := &Node{
		Val:  0,
		Next: nil,
	}
	curr := head
	for i := 1; i < size; i++ {
		curr.Next = &Node{
			Val: i,
		}
		curr = curr.Next
	}
	return &LinkedList{head: head}
}

func main() {
	list := generate(10)

	//list.remove(3)
	list.removeTarget2(list.head.Next.Next.Next.Next)
	fmt.Println("==============")
	list.println()
}
