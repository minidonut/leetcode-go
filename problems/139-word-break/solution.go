package problem0139

import (
	"fmt"
)

func Solution(args ...interface{}) bool {
	s := args[0].(string)
	wordDict := args[1].([]string)

	fmt.Println(s)
	fmt.Println(wordDict)
	return false
}
