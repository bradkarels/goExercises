package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	}
	lenFb := len(flowerbed)
	nbr := n
	switch lenFb {
	case 1:
		if flowerbed[0] == 0 && n == 1 {
			return true
		}
	case 2:
		fb0 := &flowerbed[0]
		fb1 := &flowerbed[1]
		if *fb0 == 0 && *fb1 == 0 && n == 1 {
			return true
		}
	default: // lenFb > 2
		lastIdx := lenFb - 1
		for idx, curr := range flowerbed {
			if nbr <= 0 {
				return true
			}
			var prev, next int
			switch idx {
			case 0:
				prev = 0
				next = flowerbed[idx+1]
				if canPlant(&prev, &curr, &next) {
					flowerbed[idx] = 1
					nbr = nbr - 1
				}
				continue
			case lastIdx:
				prev = flowerbed[idx-1]
				next = 0
				if canPlant(&prev, &curr, &next) {
					//flowerbed[idx] = 1 // no need to update, last cycle
					nbr = nbr - 1
				} else {
					return false
				}
				continue
			default:
				prev = flowerbed[idx-1]
				next = flowerbed[idx+1]
				if canPlant(&prev, &curr, &next) {
					flowerbed[idx] = 1
					nbr = nbr - 1
				}
				continue
			}
		}
	}
	// Should never get here...
	switch {
	case nbr <= 0:
		return true
	default:
		return false
	}
}

func canPlant(prev, curr, next *int) bool {
	return *prev+*curr+*next == 0
}
