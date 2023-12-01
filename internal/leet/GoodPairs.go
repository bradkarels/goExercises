package main

// Given an array of integers nums, return the number of good pairs.
// A pair (i, j) is called good if nums[i] == nums[j] and i < j.
func numIdenticalPairs(nums []int) int {

	ln := len(nums)
	cnt := 0
	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if nums[i] == nums[j] {
				cnt++
			}
		}
	}
	return cnt
}
