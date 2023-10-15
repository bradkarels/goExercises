package main

func maxOperations(nums []int, k int) int {
	cnt := 0
	pruned := prune(nums, k)

	for {
		idxSm, idxLg, found := findPair(pruned, k)
		if found {
			cnt++
			pruned = rmIdxs(pruned, []int{idxLg, idxSm})
			continue
		}
		break
	}
	return cnt
}

// Must use descending list of indexes...
func rmIdxs(nums []int, idxs []int) []int {
	for _, idx := range idxs {
		nums = dropAtIdx(nums, idx)
	}
	return nums
}

func findPair(nums []int, k int) (int, int, bool) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == k {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func prune(nums []int, k int) []int {
	var idxStack []int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			idxStack = append([]int{i}, idxStack...) // 5,4,3,2,1 decending...
		}
	}
	for _, idx := range idxStack {
		nums = dropAtIdx(nums, idx)
	}
	return nums
}

func dropAtIdx(nums []int, idx int) []int {
	return append(nums[:idx], nums[idx+1:]...)
}
