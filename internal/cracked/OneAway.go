package cracked

// Check to see if a string goof can be corrected with a single character operation:
// 1. Insert a character
// 2. remove a character
// 3. replace a character
func oneAway(valid, attempt string) bool {
	if valid == attempt {
		return false // Not "one off!"
	}
	validRa := []rune(valid)
	attemptRa := []rune(attempt)

	lenValid := len(validRa)
	lenAttempt := len(attemptRa)
	maximum := lenValid + 1
	minium := lenValid - 1
	if lenAttempt < minium || lenAttempt > maximum {
		return false
	} else {
		switch {
		case lenAttempt == maximum: // extra  rune (len of attempt == maximum)
			return exactlyOneTooManyRune(validRa, attemptRa)
		case lenAttempt == minium: // missing rune (len of attempt == minium)
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
		if vIdx+1 > len(v) || aIdx+1 > len(a) {
			break
		}
		if v[vIdx] == a[aIdx] {
			vIdx++
			aIdx++
			continue
		} else {
			errCnt++
			vIdx++
			continue
		}
	}
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
