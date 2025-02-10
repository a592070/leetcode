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

// It can be solved by using 2 arrays: leftProducts and rightProducts,
// which are present the product of value for each element from left and right side, respectively.
// For example, nums = [1,2,3]
// leftProducts = [x, 1, 2] => leftProducts[i] = leftProducts[i-1] * nums[i-1]
// rightProducts = [6, 3, x] => rightProducts[i] = rightProducts[i+1] * nums[i+1]
// So, for each element of answer array would be leftProducts[i] * rightProducts[i]
func productExceptSelf2(nums []int) []int {
	size := len(nums)
	leftProducts := make([]int, size)
	rightProducts := make([]int, size)
	rs := make([]int, size)

	leftProducts[0] = 1
	rightProducts[size-1] = 1
	for i := 1; i < size; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
		rightProducts[size-i-1] = rightProducts[size-i] * nums[size-i]
	}

	for i := 0; i < size; i++ {
		rs[i] = leftProducts[i] * rightProducts[i]
	}
	return rs
}
