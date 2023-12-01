package main

func isSubsequence(s, t string) bool {
	seq := []rune(s)
	lseq := len(seq)
	if lseq == 0 {
		return true
	}
	subseq := []rune(t)

	for _, r := range seq {
		if idx, ok := indexOfRuneSlice(r, subseq); ok {
			lseq--
			if lseq == 0 {
				return true
			}
			subseq = subseq[idx+1:]
			continue
		}
	}
	return false
}

func indexOfRuneSlice(r rune, runes []rune) (int, bool) {
	for idx, rr := range runes {
		if rr == r {
			return idx, true
		}
	}
	return -1, false
}
