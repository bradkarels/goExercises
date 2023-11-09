package main

import "fmt"

func longestOnes(nums []int, k int) int {
	//l := len(nums)
	//idx0 := 0
	//idx1 := 0
	ones := 0
	maxOnes := 0
	kk := k
	//for idx, v := range nums {
	// find largest "string" of numbers with k zeros
	//}
	for idx, v := range nums {
		if v == 1 {
			ones++
		} else if kk > 0 {
			ones++
			kk--
		} else {
			if ones > maxOnes {
				maxOnes = ones
			}
			ones = 0
		}
		fmt.Println(idx, v, kk, ones, maxOnes)
	}

	return 42
}
