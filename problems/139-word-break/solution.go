package main

import (
	"strings"
)

func Solve(s string, wordDict []string) bool {
	n := len(s)
	cache := make([]bool, 1024) // assumption
	for i := 0; i < n; i++ {
		if !cache[i] && i != 0 {
			continue
		}

		substr := s[i:]
		for _, w := range wordDict {
			if cache[i+len(w)] {
				continue
			}
			if strings.HasPrefix(substr, w) {
				cache[i+len(w)] = true
			}
		}
	}

	return cache[len(s)]
}
