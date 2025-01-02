package main

import "fmt"

func main() {
}

// Time Limit Exceeded
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	rs := 0
	for i := 0; i < len(customers); i++ {
		temp := 0
		for j := 0; j < len(customers); j++ {
			if (j >= i && j < i+minutes) || grumpy[j] == 0 {
				temp += customers[j]
			}
		}

		if temp > rs {
			rs = temp
		}
	}
	return rs
}

func maxSatisfied2(customers []int, grumpy []int, minutes int) int {
	if len(customers) == 0 || len(grumpy) == 0 || minutes == 0 {
		return 0
	}
	var includeAll bool
	if len(customers) == minutes {
		includeAll = true
	}

	rs := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 || includeAll {
			rs += customers[i]
		}
	}

	if includeAll {
		return rs
	}

	plus := 0
	for i := 0; i <= len(customers)-minutes; i++ {
		temp := 0
		for j := i; j < i+minutes; j++ {
			if grumpy[j] == 1 {
				temp += customers[j]
			}
		}

		if temp > plus {
			plus = temp
		}
	}

	return rs + plus
}

// queue
type queue struct {
	sum         int
	size        int
	collections []node
}
type node struct {
	idx      int
	val      int
	isGrumpy bool
}

func (q *queue) push(n node) {
	var head node
	if len(q.collections) >= q.size {
		head = q.collections[0]
		q.collections = append(q.collections[1:], n)
	} else {
		q.collections = append(q.collections, n)
	}
	if head.isGrumpy {
		q.sum -= head.val
	}
	if n.isGrumpy {
		q.sum += n.val
	}
}

func maxSatisfied3(customers []int, grumpy []int, minutes int) int {
	if len(customers) == 0 || len(grumpy) == 0 || minutes == 0 {
		return 0
	}

	q := &queue{size: minutes}
	for minute := 0; minute < minutes; minute++ {
		q.push(node{idx: minute, val: customers[minute], isGrumpy: grumpy[minute] == 1})
	}
	rs := q.sum

	for customerIdx := minutes; customerIdx < len(customers); customerIdx++ {
		q.push(node{idx: customerIdx, val: customers[customerIdx], isGrumpy: grumpy[customerIdx] == 1})
		if q.sum > rs {
			fmt.Printf("%d: %d -> %d\n", customerIdx, rs, q.sum)
			rs = q.sum
		}
	}

	for customerIdx := 0; customerIdx < len(customers); customerIdx++ {
		if grumpy[customerIdx] == 0 {
			rs += customers[customerIdx]
		}
	}

	return rs
}
