package cracked

import (
	"strconv"
)

func compress(s string) string {
	ra := []rune(s)
	lenra := len(ra)
	if lenra <= 1 {
		return s
	}

	res := []rune{}
	current := ra[0]
	idx := 1
	cnt := 1
	for idx < lenra {
		next := ra[idx]
		if current == next {
			cnt++
		} else {
			res = append(res, current)
			res = append(res, []rune(strconv.Itoa(cnt))...)
			cnt = 1
		}
		current = next
		idx++
		if idx == lenra {
			res = append(res, current)
			res = append(res, []rune(strconv.Itoa(cnt))...)
		}
	}
	if len(res) < lenra {
		return string(res)
	} else {
		return s
	}
}
