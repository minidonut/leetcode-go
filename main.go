package main

import (
	"fmt"

	problem0139 "github.com/minidonut/leetcode-golang/problems/139-word-break"
)

func main() {
	fmt.Println("Hello world!")
	for _, c := range problem0139.Cases {
		output := problem0139.Solution(c.Input.S, c.Input.WordDict)
		result := output == c.Output
		fmt.Println(result)
	}
}
