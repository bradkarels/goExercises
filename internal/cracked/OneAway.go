package cracked

func oneAway(valid, attempt string) bool {
	if valid == attempt {
		return false // Not "one off!"
	}
	validRa := []rune(valid)
	attemptRa := []rune(attempt)

	lenValid := len(validRa)
	lenAttempt := len(attemptRa)
	max := lenValid + 1
	min := lenValid - 1
	if lenAttempt < min || lenAttempt > max {
		return false
	} else {
		switch {
		case lenAttempt == max: // extra  rune (len of attempt == max)
			return exactlyOneTooManyRune(validRa, attemptRa)
		case lenAttempt == min: // missing rune (len of attempt == min)
			return missingExactlyOneRune(validRa, attemptRa)
		default: // incorrect rune (len must be ==)
			return hasExactlyOneIncorrectRune(validRa, attemptRa)
		}
	}
}

func exactlyOneTooManyRune(v, a []rune) bool {
	vIdx := 0
	aIdx := 0
	errCnt := 0
	for errCnt < 2 {
		if vIdx+1 > len(v) || aIdx+1 > len(a) {
			break
		}
		if v[vIdx] == a[aIdx] {
			// good
			vIdx++
			aIdx++
			continue
		} else {
			errCnt++
			aIdx++
			continue
		}
	}
	return errCnt < 2
}

func missingExactlyOneRune(v, a []rune) bool {
	vIdx := 0
	aIdx := 0
	errCnt := 0
	for errCnt < 2 {
		//fmt.Printf("entering: %d %d\n", vIdx, aIdx)
		if vIdx+1 > len(v) || aIdx+1 > len(a) {
			//fmt.Printf("breaking: %d %d\n", vIdx, aIdx)
			break
		}
		if v[vIdx] == a[aIdx] {
			//fmt.Printf("muthafucka - we match!: %d %d\n", v[vIdx], a[aIdx])
			vIdx++
			aIdx++
			continue
		} else {
			errCnt++
			vIdx++
			continue
		}
	}
	//fmt.Printf("errCnt: %d %t\n", errCnt, errCnt < 2)
	return errCnt < 2
}

func hasExactlyOneIncorrectRune(v, a []rune) bool {
	// equal len "ensured"
	nbrIncorrect := 0
	for i := 0; i < len(v); i++ {
		if v[i] == a[i] {
			continue
		} else {
			nbrIncorrect++
			if nbrIncorrect > 1 {
				return false
			}
		}
	}
	return true
}
