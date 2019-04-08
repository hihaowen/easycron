package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runJob(func() {
		fmt.Println("我是每隔3秒就运行的任务", time.Now().UTC())
	}, nextDurationForRecurrent(3))

	runJob(func() {
		fmt.Println("我是每天11:08:30才运行的任务", time.Now().UTC())
	}, nextDurationForDaily(11, 8, 30))

	runtime.Goexit()
}

// 适合每隔几秒就执行的任务
func nextDurationForRecurrent(interval int64) time.Duration {
	return time.Second * time.Duration(interval)
}

// 适合每天定时执行的任务
func nextDurationForDaily(hour, minute, second int) time.Duration {
	nowTime := time.Now()

	year, month, day := nowTime.Date()

	nextTime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)

	if ! nowTime.Before(nextTime) {
		nextTime = time.Date(year, month, day+1, hour, minute, second, 0, time.UTC)
	}

	return nextTime.Sub(nowTime)
}

// 执行任务
func runJob(fn func(), d time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(d):
				fn()
			}
		}
	}()
}
