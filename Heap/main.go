package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	minHeap := InitMinHeap([]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7})

	fmt.Println(minHeap.ToString())
	expected := &IntHeap{4, 6, 5, 8, 7, 10, 13, 12, 11, 9}
	heap.Init(expected)
	if !equals(*expected, minHeap.data) {
		fmt.Printf("Expected %v, got %v\n", *expected, minHeap.data)
	}

	fmt.Println(minHeap.Peek())
	minHeap.Pop()
	fmt.Println(minHeap.ToString())
	heap.Pop(expected)
	if !equals(*expected, minHeap.data) {
		fmt.Printf("Expected %v, got %v\n", expected, minHeap.data)
	}

	minHeap.Push(3)
	heap.Push(expected, 3)
	minHeap.Push(21)
	heap.Push(expected, 21)
	fmt.Println(minHeap.ToString())
	if !equals(*expected, minHeap.data) {
		fmt.Printf("Expected %v, got %v\n", expected, minHeap.data)
	}

}

func equals(arr []int, expect []int) bool {
	if len(arr) != len(expect) {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != expect[i] {
			return false
		}
	}
	return true
}

/*
MinHeap Tree
*/
type MinHeap struct {
	data []int
}

func InitMinHeap(arr []int) *MinHeap {
	heap := &MinHeap{
		data: make([]int, 0),
	}
	for _, v := range arr {
		heap.Push(v)
	}
	return heap
}

func (h MinHeap) ToString() string {
	builder := strings.Builder{}
	builder.WriteString("MinHeap{")
	for i, v := range h.data {
		builder.WriteString(fmt.Sprintf("[i=%d, v=%d]", i, v))
		if i != len(h.data)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

func (h MinHeap) Len() int {
	return len(h.data)
}

func (h MinHeap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) heapifyUp(i int) {
	for h.Less(i, h.parentIndex(i)) {
		h.Swap(i, h.parentIndex(i))
		i = h.parentIndex(i)
	}
}
func (h *MinHeap) heapifyDown(i int) {
	left, right := h.leftChildInde(i), h.rightChildInde(i)
	largest := i
	if left < h.Len() && h.Less(left, i) {
		largest = left
	}
	if right < h.Len() && h.Less(right, largest) {
		largest = right
	}
	if largest != i {
		h.Swap(i, largest)
		h.heapifyDown(largest)
	}
}

func (h MinHeap) leftChildInde(currentIndex int) int {
	return 2*currentIndex + 1
}
func (h MinHeap) rightChildInde(currentIndex int) int {
	return 2*currentIndex + 2
}

func (h MinHeap) parentIndex(currentIndex int) int {
	return (currentIndex - 1) / 2
}

func (h *MinHeap) Push(num interface{}) {
	h.data = append(h.data, num.(int))
	h.heapifyUp(h.Len() - 1)
}

func (h MinHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h.data[0]
}
func (h *MinHeap) Pop() interface{} {
	old := h.data
	size := h.Len()
	popEle := old[0]
	h.Swap(0, size-1)
	h.data = old[0 : size-1]
	h.heapifyDown(0)
	return popEle
}

/*
from container/heap/example_intheap_test.go
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
