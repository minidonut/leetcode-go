package main

import "github.com/minidonut/leetcode-go/pkg/utils"

func main() {
	cases := GenerateCase()

	logger := utils.CaseLogger{}
	logger.Init(len(cases))
	for i, c := range cases {
		logger.Start(i, c.input, c.output)
		output := Solve(c.input.Nums, c.input.Target)
		logger.Stop(i, Output{Value: output})
	}

	logger.Print()
}
