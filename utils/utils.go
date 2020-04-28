package utils

import (
	"fmt"
	"runtime"
	"time"
)

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
