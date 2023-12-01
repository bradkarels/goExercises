package main

func addBinary(a string, b string) string {
	return doAddBinary([]rune(a), []rune(b), []rune{}, '0')
}

func doAddBinary(a, b, sum []rune, c rune) string {
	la := len(a)
	lb := len(b)
	if (la == 0 && lb == 0) && c == '0' {
		return string(sum)
	}
	var aa, bb rune
	var aLongHead, bLongHead []rune
	if la > 0 {
		aa = a[len(a)-1:][0]
		aLongHead = a[:len(a)-1]
	} else {
		aa = '0'
		aLongHead = []rune{}
	}
	if lb > 0 {
		bb = b[len(b)-1:][0]
		bLongHead = b[:len(b)-1]
	} else {
		bb = '0'
		bLongHead = []rune{}
	}
	switch {
	case aa == '0' && bb == '0' && c == '1':
		newSum := append([]rune{'1'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '0')
	case aa == '1' && bb == '0' && c == '0':
		newSum := append([]rune{'1'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '0')
	case aa == '1' && bb == '0' && c == '1':
		newSum := append([]rune{'0'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '1')
	case aa == '0' && bb == '1' && c == '0':
		newSum := append([]rune{'1'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '0')
	case aa == '0' && bb == '1' && c == '1':
		newSum := append([]rune{'0'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '1')
	case aa == '1' && bb == '1' && c == '1':
		newSum := append([]rune{'1'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '1')
	case aa == '1' && bb == '1' && c == '0':
		newSum := append([]rune{'0'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '1')
	default: // case aa == '0' && bb == '0' && c == '0':
		newSum := append([]rune{'0'}, sum...)
		return doAddBinary(aLongHead, bLongHead, newSum, '0')
	}
}
