package leet

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	head := func(s []string) (string, bool) {
		if len(s) > 0 {
			return s[:1][0], true
		}
		return "", false
	}

	tail := func(s []string) []string {
		if len(s) > 0 {
			return s[1:]
		}
		return []string{}
	}

	merging := []string{}

	word1chan := make(chan []string, 1)
	word2chan := make(chan []string, 1)
	go str2slice(word1, word1chan)
	go str2slice(word2, word2chan)

	s1 := <-word1chan
	s2 := <-word2chan

	doneWith1 := false
	doneWith2 := false

	for {
		if !doneWith1 {
			if head1, ok := head(s1); ok {
				merging = append(merging, head1)
				s1 = tail(s1)
			} else {
				flipDaSwitch(&doneWith1)
			}
		}
		if !doneWith2 {
			if head2, ok := head(s2); ok {
				merging = append(merging, head2)
				s2 = tail(s2)
			} else {
				flipDaSwitch(&doneWith2)
			}
		}
		if doneWith1 && doneWith2 {
			break
		}
	}
	return strings.Join(merging, ``)
}

func str2slice(word string, sliceChannel chan []string) {
	sliceChannel <- strings.Split(word, ``)
}

func flipDaSwitch(sw *bool) {
	*sw = !*sw
}
