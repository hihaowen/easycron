# easycron
cron计划任务简单实现

```go
new(Task).Every(3).Run(func() {
	fmt.Println("我是每隔3秒就运行的任务", time.Now().UTC())
})
```