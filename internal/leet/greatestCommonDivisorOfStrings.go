package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1, str2 string) string {
	big, sml := getBigAndSmallString(str1, str2)
	//lgSlice := strings.Split(big, ``)
	smSlice := strings.Split(sml, ``)

	commonDivisor := ``
	for idx := range smSlice {
		div := strings.Join(smSlice[:idx+1], ``)
		if isDivisorOfBig := divideString(big, div); isDivisorOfBig {
			if isDivisorOfSml := divideString(sml, div); isDivisorOfSml {
				commonDivisor = div
				fmt.Println(commonDivisor)
			}
		}
	}
	return commonDivisor
}

func divideString(str, divisor string) bool {
	var isValidDivisor bool
	divLen := len(divisor)
	s := strings.Split(str, ``)
	if (len(s) % len(divisor)) == 0 {
		for {
			if len(s) >= len(divisor) {
				if strings.Join(s[:divLen], ``) == divisor {
					//fmt.Println(s)
					s = s[divLen:]
					//fmt.Println(s)
					//fmt.Println("----------")
				} else {
					break
				}
			} else {
				isValidDivisor = true
				break
			}
		}
	}
	return isValidDivisor
}

func getBigAndSmallString(s0, s1 string) (string, string) {
	if len(s0) > len(s1) {
		return s0, s1
	}
	return s1, s0
}
