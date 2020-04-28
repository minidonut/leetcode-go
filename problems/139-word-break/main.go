package main

import "github.com/minidonut/leetcode-go/utils"

func main() {
	cases := GenerateCase()

	logger := utils.CaseLogger{}
	logger.Init(len(cases))
	for i, c := range cases {
		logger.Start(i, c.input, c.output)
		output := Solve(c.input.s, c.input.wordDict)
		logger.Stop(i, Output{value: output})
	}

	logger.Print()
}
