package main

import (
	"testing"
	"time"
)

func BenchmarkSleepWith(b *testing.B) {
	b.StopTimer()
	// 数据准备阶段
	time.Sleep(time.Second * 2)
	b.StartTimer()
	
	// 实测函数
	SleepWith()
}