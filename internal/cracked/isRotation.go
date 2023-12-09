package cracked

import (
	"strings"
)

func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	r1 := []rune(s1)
	r2 := []rune(s2)

	// find indecies of s1[0] in s2 (zero means not a rotation)
	indicies, ok := getIndicies(r2, r1[0])
	if !ok {
		return false
	}
	leRotation := false
	for i := 0; i < len(indicies); i++ {
		idx := indicies[i]
		var builder strings.Builder
		builder.WriteString(s2[idx:])
		builder.WriteString(s2[:idx])
		s3 := builder.String()
		if s3 == s1 {
			leRotation = true
			break
		}
	}
	return leRotation
}

func getIndicies(ra []rune, r rune) ([]int, bool) {
	res := []int{}
	for i := 0; i < len(ra); i++ {
		if ra[i] == r {
			res = append(res, i)
		}
	}
	if len(res) > 0 {
		return res, true
	} else {
		return nil, false
	}
}
