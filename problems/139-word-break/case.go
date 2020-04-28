package main

type Input struct {
	s        string
	wordDict []string
}

type Output struct {
	value bool
}

func (o Output) String() string {
	if o.value {
		return "true"
	} else {
		return "false"
	}
}

type Case struct {
	input  Input
	output Output
}

func GenerateCase() []Case {
	cases := []Case{
		{input: Input{s: "leetcode", wordDict: []string{"leet", "code"}}, output: Output{value: true}},
		{input: Input{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, output: Output{value: false}},
		{input: Input{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, output: Output{value: false}},
		{input: Input{s: "applepie", wordDict: []string{"pie", "pear", "apple", "peach"}}, output: Output{value: true}},
		{input: Input{s: "aaaaaaa", wordDict: []string{"aaaa", "aaa"}}, output: Output{value: true}},
	}
	return cases
}
