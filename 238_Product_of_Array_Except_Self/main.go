package main

func main() {

}

func productExceptSelf(nums []int) []int {
	sum := 1
	countZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZero++
			continue
		}
		sum *= nums[i]
	}
	rs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if countZero > 1 {
			rs[i] = 0
		} else if nums[i] == 0 {
			rs[i] = sum
		} else if countZero == 1 {
			rs[i] = 0
		} else {
			rs[i] = sum / nums[i]
		}
	}
	return rs
}
