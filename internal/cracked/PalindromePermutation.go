package cracked

import (
	"fmt"
	"strings"
)

// Given a string write a function to determine if it is a permutation of a palindrome. (need not be real words...)
func isPalindromePermutation(s string) bool {
	s = strings.ToLower(s)
	ra := removeSpace([]rune(s))
	fmt.Printf("-->%s<--\n", string(ra))
	runeMap := make(map[rune]int)
	for _, v := range ra {
		cnt, exists := runeMap[v]
		if exists {
			runeMap[v] = cnt + 1
		} else {
			runeMap[v] = 1
		}
	}
	// At most one odd cnt...
	odds := 0
	for _, cnt := range runeMap {
		if cnt%2 != 0 {
			odds++
			if odds > 1 {
				return false
			}
		}
	}
	return true
}

func removeSpace(ra []rune) []rune {
	newRa := []rune{}
	for _, v := range ra {
		if v != ' ' {
			newRa = append(newRa, v)
		}
	}
	return newRa
}
