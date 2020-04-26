package problem0139

type Input struct {
	S        string
	WordDict []string
}

type Case struct {
	Input  Input
	Output bool
}

var Cases []Case = []Case{
	{Input: Input{S: "leetcode", WordDict: []string{"leet", "code"}}},
}
