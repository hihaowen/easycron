# easycron
cron计划任务简单实现

```go
package main

import (
	"fmt"
	"github.com/hihaowen/easycron"
	"runtime"
	"time"
)

func main() {
	easycron.Run(func() {
		fmt.Println("我是每隔3秒就运行的任务", time.Now().UTC())
	}, easycron.RecurrentTimer{Interval: 3})

	easycron.Run(func() {
		fmt.Println("我是每天 14:06:30 才运行的任务", time.Now().UTC())
	}, easycron.DailyTimer{Hour: 14, Minute: 06, Second: 30})

	once := easycron.OnceTimer{Year: 2019, Month: 4, Day: 8}
	once.Hour = 14
	once.Minute = 49
	once.Second = 25

	easycron.Run(func() {
		fmt.Println("我是 2019-04-08 14:49:25 运行的任务", time.Now().UTC())
	}, once)

	runtime.Goexit()
}
```