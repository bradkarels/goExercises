package main

func longestSubarray(nums []int) int {
	var condensed []int
	nbrOfOnes := 0
	for _, v := range nums {
		if v == 0 {
			if nbrOfOnes > 0 {
				condensed = append(condensed, nbrOfOnes, 0)
				nbrOfOnes = 0
			} else {
				condensed = append(condensed, 0)
			}
		} else {
			nbrOfOnes++
		}
	}
	if nbrOfOnes > 0 {
		condensed = append(condensed, nbrOfOnes)
	}
	l := len(condensed)
	switch l {
	case 0:
		return 0
	case 1:
		if condensed[0] == 0 {
			return 0
		}
		return condensed[0] - 1
	case 2:
		return condensed[0] + condensed[1]
	default:
		maximus := 0
		for i := 0; i < l-2; i++ {
			sum := condensed[i] + condensed[i+1] + condensed[i+2]
			if sum > maximus {
				maximus = sum
			}
		}
		return maximus
	}
}
