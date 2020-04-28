package main

import "github.com/minidonut/leetcode-go/containers"

type Input struct {
	s        string
	wordDict []string
}

type Output = containers.Boolean

type Case struct {
	input  Input
	output Output
}

func GenerateCase() []Case {
	cases := []Case{
		{input: Input{s: "leetcode", wordDict: []string{"leet", "code"}}, output: Output{Value: true}},
		{input: Input{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, output: Output{Value: false}},
		{input: Input{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, output: Output{Value: false}},
		{input: Input{s: "applepie", wordDict: []string{"pie", "pear", "apple", "peach"}}, output: Output{Value: true}},
		{input: Input{s: "aaaaaaa", wordDict: []string{"aaaa", "aaa"}}, output: Output{Value: true}},
	}
	return cases
}
