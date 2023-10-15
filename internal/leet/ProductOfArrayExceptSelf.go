package main

import "fmt"

type idxVal struct {
	idx   int
	value int
}

func productExceptZelf(nums []int) []int {
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
		fmt.Printf("i:%d j:%d\n", i, j)
		fmt.Println(left)
		fmt.Println(right)
	}
	for i := 0; i < l; i++ {
		ans[i] = left[i] * right[i]
	}
	fmt.Println(ans)
	return ans
}

func productExceptSelf(nums []int) []int {
	numslen := len(nums)
	products := make([]int, numslen)

	ivs := make(chan idxVal)

	for idx, _ := range nums {
		go productForIdx(idx, nums, ivs)
	}

	cnt := 0
	for {
		if cnt >= numslen {
			break
		}
		iv := <-ivs
		products[iv.idx] = iv.value
		cnt++
	}
	return products
}

func productForIdx(idx int, nums []int, c chan idxVal) {
	subbedOut := false
	product := 1
	for i, n := range nums {
		if !subbedOut && i == idx {
			subbedOut = true
			continue
		} else if n == 1 {
			continue
		} else if n == 0 {
			product = 0
			break
		}
		product *= n
	}
	c <- idxVal{idx: idx, value: product}
}
