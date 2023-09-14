package main

func main() {

}

func twoSum(nums []int, target int) []int {
	size := len(nums)
	if size < 2 {
		return []int{0}
	}
	if size == 2 {
		return []int{0, 1}
	}
	tempStore := map[int]int{}
	for i := 0; i < len(nums); i++ {
		currVal := nums[i]
		expectIdx := tempStore[target-currVal]
		if expectIdx > 0 {
			return []int{i, expectIdx - 1}
		} else {
			tempStore[currVal] = i + 1
		}
	}
	return []int{}
}

/*
if nums is sorted
*/
func twoSumFromSorted(nums []int, target int) []int {
	size := len(nums)
	if size < 2 {
		return []int{0}
	}
	if size == 2 {
		return []int{0, 1}
	}
	tempStore := map[int]int{}
	for i := 0; i < len(nums); i++ {
		currVal := nums[i]
		expectIdx := tempStore[target-currVal]
		if expectIdx > 0 {
			return []int{i, expectIdx - 1}
		} else {
			tempStore[currVal] = i + 1
		}
	}
	return []int{}
}
