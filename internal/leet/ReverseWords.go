package main

import (
	"strings"
)

func reverseWords(s string) string {
	var wordStack []string
	var currentWord string

	for {
		if len(s) == 0 {
			if currentWord != `` {
				wordStack = append([]string{currentWord}, wordStack...)
			}
			break
		}
		idx := strings.Index(s, ` `)
		if idx != 0 {
			currentWord += string(s[0])
			s = s[1:]
			continue
		} else {
			if currentWord != `` {
				// add to stack
				wordStack = append([]string{currentWord}, wordStack...)
				// reset
				currentWord = ``
			} else {
				s = s[1:]
			}
		}
	}
	return strings.Join(wordStack, ` `)
}
