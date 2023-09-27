package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	for idx, _ := range candies {
		bonus := giveExtra(candies, extraCandies, idx)
		isTrue := ithHasMost(idx, bonus)
		res = append(res, isTrue)
	}
	return res
}

func ithHasMost(idx int, slice []int) bool {
	var ithValue = slice[idx]
	result := true
	for _, v := range slice {
		if v > ithValue {
			result = false
		}
	}
	return result
}

func giveExtra(candiez []int, extraCandies, idx int) []int {
	var ia []int
	for i, v := range candiez {
		if i == idx {
			ia = append(ia, v+extraCandies)
		} else {
			ia = append(ia, v)
		}
	}
	return ia
}
