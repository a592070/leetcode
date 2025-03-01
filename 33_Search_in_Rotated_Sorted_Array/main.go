package main

func main() {

}

func search(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
