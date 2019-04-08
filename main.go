package main

import (
	"easycron/cron"
	"fmt"
	"runtime"
	"time"
)

func main() {
	cron.Run(func() {
		fmt.Println("我是每隔3秒就运行的任务", time.Now().UTC())
	}, cron.RecurrentTimer{Interval: 3})

	cron.Run(func() {
		fmt.Println("我是每天 14:06:30 才运行的任务", time.Now().UTC())
	}, cron.DailyTimer{Hour: 14, Minute: 06, Second: 30})

	once := cron.OnceTimer{Year: 2019, Month: 4, Day: 8}
	once.Hour = 14
	once.Minute = 49
	once.Second = 25

	cron.Run(func() {
		fmt.Println("我是 2019-04-08 14:49:25 运行的任务", time.Now().UTC())
	}, once)

	runtime.Goexit()
}
