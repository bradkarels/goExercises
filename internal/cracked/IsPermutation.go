package cracked

import "strings"

// Given two strings write a function to determine if one is a permutation of
// the  other.
func isPermutation(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	// count of all rune must be cnt % 2 == 0 - crude map sol.
	runeMap := map[rune]int{}

	if len(a) != len(b) {
		return false
	}
	// for idx for loop - combine ops?
	for _, aVal := range a { // to fn
		cnt, exists := runeMap[aVal]
		if exists {
			runeMap[aVal] = cnt + 1
			continue
		} else {
			runeMap[aVal] = 1
		}
	}
	for _, bVal := range b { // to fn
		cnt, exists := runeMap[bVal]
		if exists {
			runeMap[bVal] = cnt + 1
			continue
		} else {
			// A new element signals not a permutation.
			return false
		}
	}
	for _, v := range runeMap {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
