package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	ans := make([]int, l)
	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		j := len(nums) - i - 1
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
	}
	for i := 0; i < l; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
