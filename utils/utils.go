package utils

import (
	"fmt"
	"runtime"
	"time"
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

type Usage struct {
	start   time.Time
	elapsed time.Duration
	memory  uint64
	gc      uint32
}

func (u *Usage) Init() *Usage {
	runtime.GC()
	u.start = time.Now()
	return u
}

func (u *Usage) Measure() *Usage {
	u.elapsed = time.Since(u.start)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	u.memory = bToKB(m.Alloc)
	u.gc = m.NumGC

	return u
}

func (u Usage) String() string {
	return fmt.Sprintf("Runtime: %s Memory: %v KB NumGC = %v", u.elapsed, u.memory, u.gc)
}

func bToKB(b uint64) uint64 {
	return b / 1024
}

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}

type result struct {
	usage    Usage
	input    interface{}
	output   fmt.Stringer
	expected fmt.Stringer
	result   string
}

type CaseLogger struct {
	reuslts []result
}

func (c *CaseLogger) Init(length int) {
	c.reuslts = make([]result, length)
}

func (c *CaseLogger) Start(i int, input interface{}, expected fmt.Stringer) {
	r := &c.reuslts[i]
	r.input = input
	r.expected = expected
	r.usage = Usage{}
	r.result = string(colorRed) + "fail"
	r.usage.Init()
}

func (c *CaseLogger) Stop(i int, output fmt.Stringer) {
	r := &c.reuslts[i]
	r.usage.Measure()
	r.output = output
	if r.output.String() == r.expected.String() {
		r.result = string(colorGreen) + "success"
	}
}

func (c CaseLogger) Print() {
	for _, r := range c.reuslts {
		fmt.Println(string(colorReset)+"[input]    :", r.input)
		fmt.Println("[expected] :", r.expected)
		fmt.Println("[output]   :", r.output)
		fmt.Println("[result]   :", r.result+string(colorYellow))
		fmt.Println(r.usage)
	}
}
