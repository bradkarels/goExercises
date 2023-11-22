package main

import (
	"strconv"
	"unicode"
)

func decodeString(s string) string {
	if isEncoded(s) {
		result := []rune(s)
		for {
			rbi := getRightBraceIdx(result)
			if rbi == -1 {
				break
			}
			tail := result[rbi+1:]
			head := result[:rbi]
			ltrs := getLtrsFromRightToLeftBrace(head)
			head = head[:len(head)-1-len(ltrs)]
			multiplier, digits := getNbrsLeftOfRightBrace(head)
			head = head[:len(head)-digits]
			multStr := multiplyStr(ltrs, multiplier)
			result = append(head, append(multStr, tail...)...)
		}
		return string(result)
	}
	return s
}

func isEncoded(s string) bool {
	ra := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if unicode.IsDigit(ra[i]) || unicode.IsPunct(ra[i]) {
			return true
		}
	}
	return false
}

func multiplyStr(ra []rune, m int) []rune {
	res := []rune{}
	for i := 0; i < m; i++ {
		res = append(res, ra...)
	}
	return res
}

func getNbrsLeftOfRightBrace(ra []rune) (int, int) {
	//leInt := 0
	digits := 0
	nbrs := []rune{}
	for i := len(ra) - 1; i >= 0; i-- {
		if unicode.IsDigit(ra[i]) {
			digits++
			nbrs = append([]rune{ra[i]}, nbrs...)
		} else if unicode.IsPunct('[') {
			break
		}
	}
	nbr, _ := strconv.Atoi(string(nbrs))
	return nbr, digits
}

func getLtrsFromRightToLeftBrace(head []rune) []rune {
	ltrs := []rune{}
	for i := len(head) - 1; i >= 0; i-- {
		if unicode.IsLetter(head[i]) {
			ltrs = append([]rune{head[i]}, ltrs...)
		} else if unicode.IsPunct('[') {
			break
		}
	}
	return ltrs
}

func getRightBraceIdx(ra []rune) int {
	for idx, v := range ra {
		if v == ']' {
			return idx
		}
	}
	return -1
}
