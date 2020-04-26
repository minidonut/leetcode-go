package main

type Input struct {
	s        string
	wordDict []string
}

type Case struct {
	input  Input
	output bool
}

func GenerateCase() []Case {
	cases := []Case{
		{input: Input{s: "leetcode", wordDict: []string{"leet", "code"}}, output: true},
		{input: Input{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, output: false},
		{input: Input{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, output: false},
		{input: Input{s: "applepie", wordDict: []string{"pie", "pear", "apple", "peach"}}, output: false},
		{input: Input{s: "aaaaaaa", wordDict: []string{"aaaa", "aaa"}}, output: false},
	}
	return cases
}
