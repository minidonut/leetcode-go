package main

import (
	validate "github.com/minidonut/leetcode-go/validator"
)

func main() {
	cases := GenerateCase()

	for i, c := range cases {

		output := Solve(c.input.s, c.input.wordDict)

		validate.SingleValue(output, c.output)
	}
}
