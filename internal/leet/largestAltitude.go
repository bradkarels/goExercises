package main

func largestAltitude(gain []int) int {
	alt := 0
	maxAlt := alt
	for _, g := range gain {
		alt += g
		if alt > maxAlt {
			maxAlt = alt
		}
	}
	return maxAlt
}
