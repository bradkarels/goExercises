package main

func maxVowels(s string, k int) int {
	vlaggen := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		vlaggen[i] = makeBool(s[i])
	}

	var mx int
	for i := 0; i < len(vlaggen)-k+1; i++ {
		ss := vlaggen[i : i+k]
		var sum int
		for h := 0; h < k; h++ {
			if ss[h] {
				sum++
			}
		}
		if sum > mx {
			mx = sum
		}
	}
	return mx
}

func makeBool(r uint8) bool {
	switch r {
	case 97:
		return true
	case 101:
		return true
	case 105:
		return true
	case 111:
		return true
	case 117:
		return true
	default:
		return false
	}
}
