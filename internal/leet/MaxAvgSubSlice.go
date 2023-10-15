package main

type maximus struct {
	isAssigned bool
	value      float64
}

func findMaxAverage(nums []int, k int) float64 {
	return getMax(nums, k)
}

func getMax(nums []int, k int) float64 {
	m := &maximus{
		isAssigned: false,
		value:      0,
	}
	l := len(nums)
	for i := 0; i < l-k+1; i++ {
		var subSlice []int
		for j := 0; j < k; j++ {
			subSlice = append(subSlice, nums[i+j])
		}
		avg := mkAvg(subSlice)
		if !m.isAssigned {
			m.isAssigned = true
			m.value = avg
		} else {
			if avg > m.value {
				m.value = avg
			}
		}
	}
	return m.value
}

func mkAvg(nums []int) float64 {
	l := float64(len(nums))
	sum := float64(mkSum(nums))
	return sum / l
}

func mkSum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
