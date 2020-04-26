package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/minidonut/leetcode-golang/utils"
)

func Validate(ans1 bool, ans2 bool) bool {
	return ans1 == ans2
}

func main() {

	cases := GenerateCase()

	utils.Print()
	for _, c := range cases {
		runtime.GC()

		start := time.Now()
		output := Solve(c.input.s, c.input.wordDict)
		elapsed := time.Since(start)

		Validate(output, c.output)
		fmt.Printf("Runtime: %s\n", elapsed)
		PrintMemUsage()
	}

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
