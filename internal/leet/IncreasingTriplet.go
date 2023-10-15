package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	lnums := len(nums)
	if lnums < 3 {
		return false
	}
	var one = math.MinInt32
	var two = math.MinInt32
	pone := &one
	ptwo := &two
	for i := 0; i < lnums-2; i++ {
		*pone = nums[i]
		for j := i + 1; j < lnums-1; j++ {
			*ptwo = nums[j]
			if *ptwo > *pone {
				for k := j + 1; k < lnums; k++ {
					fmt.Printf("%d %d %d\n", *pone, *ptwo, nums[k])
					if nums[k] > two {
						return true
					}
				}
			}
		}
	}
	return false
}
