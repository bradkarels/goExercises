package main

type Stack struct {
	runes []rune
}

func (s *Stack) Push(r rune) {
	s.runes = append([]rune{r}, s.runes...)
}

func (s *Stack) IsEmpty() bool {
	if len(s.runes) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		var zero rune
		return zero // not ideal...
	}
	retVal := s.runes[0]
	s.runes = s.runes[1:]
	return retVal
}

func reverseVowels(s string) string {
	runes := []rune(s)

	var vPos []int
	var vStack Stack

	for idx, r := range runes {
		if runeIsVowel(r) {
			vPos = append(vPos, idx)
			vStack.Push(r)
		}
	}

	for _, i := range vPos {
		runes[i] = vStack.Pop()
	}
	return string(runes)

}

func runeIsVowel(r rune) bool {
	switch r {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	case 'A':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'O':
		return true
	case 'U':
		return true
	default:
		return false
	}
}
