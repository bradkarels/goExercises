package grok

// Simple binary search - assumes and ordered list of ints.
func search(ints []int, tgt int) bool {
	lenInts := len(ints)
	if lenInts == 1 {
		if ints[0] == tgt {
			return true
		}
		return false
	}
	var nieuwInts = make([]int, lenInts)
	copy(nieuwInts, ints)

	for len(nieuwInts) > 0 {
		midIdx := len(nieuwInts) / 2
		guess := nieuwInts[midIdx]
		if guess == tgt {
			return true
		} else if guess > tgt {
			nieuwInts = nieuwInts[:midIdx]
			continue
		} else if guess < tgt {
			nieuwInts = nieuwInts[midIdx+1:]
			continue
		}
	}
	return false
}
