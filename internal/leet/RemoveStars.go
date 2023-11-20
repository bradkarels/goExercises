package main

var star = '*'

func removeStars(s string) string {
	r := string(doRemoveStars([]rune(s), 0))
	return r
}

func doRemoveStars(ra []rune, start int) []rune {
	if start < 0 {
		return ra
	} else {
		newRa, newIdx := seek(ra, start)
		return doRemoveStars(newRa, newIdx)
	}
}

func seek(ra []rune, idx int) ([]rune, int) {
	lra := len(ra)
	newIdx := -1
	var newRa []rune
	for i := idx; i < lra; i++ {
		if ra[i] == star {
			newRa = append(ra[:i-1], ra[i+1:]...)
			switch i {
			case 1:
				newIdx = 0
			default:
				newIdx = i - 2
			}
			break
		}
	}
	if newIdx == -1 {
		newRa = ra
	}
	return newRa, newIdx
}
