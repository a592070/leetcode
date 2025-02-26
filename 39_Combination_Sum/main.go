package main

import (
	"sort"
)

func main() {
}

type stack struct {
	arr []int
}

func NewStack(arr []int) *stack {
	return &stack{
		arr: arr,
	}
}
func (s *stack) Push(v int) {
	s.arr = append(s.arr, v)
}
func (s *stack) Pop() int {
	l := len(s.arr)
	if len(s.arr) == 0 {
		return -1
	}
	v := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return v
}

func (s *stack) Top() int {
	if len(s.arr) == 0 {
		return -1
	}
	return s.arr[len(s.arr)-1]
}
func (s *stack) ToArray() []int {
	return s.arr
}
func (s *stack) Empty() bool {
	return len(s.arr) == 0
}

func combinationSum(candidates []int, target int) [][]int {
	rs := [][]int{}

	sort.Ints(candidates)

	var greedy func(stack, int, stack)

	greedy = func(candidates stack, target int, selected stack) {
		if target == 0 {
			rs = append(rs, selected.ToArray())
			return
		}
		for !candidates.Empty() {
			val := candidates.Top()
			if val <= target {
				newTarget := target - val
				newSelected := NewStack(append([]int{}, selected.ToArray()...))
				newSelected.Push(val)
				greedy(candidates, newTarget, *newSelected)
			}

			candidates.Pop()
		}
		return
	}

	s := NewStack([]int{})
	for _, n := range candidates {
		s.Push(n)
	}
	greedy(*s, target, *NewStack([]int{}))

	return rs
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func([]int, int, []int)

	backtrack = func(candidates []int, target int, selected []int) {
		if target == 0 {
			result = append(result, append([]int{}, selected...))
			return
		}
		for i, val := range candidates {
			if val <= target {
				newTarget := target - val
				selected := append([]int{}, selected...)
				selected = append(selected, val)
				newCandidates := append([]int{}, candidates[i:]...)
				backtrack(newCandidates, newTarget, selected)
			}
		}
	}
	backtrack(candidates, target, []int{})
	return result
}
