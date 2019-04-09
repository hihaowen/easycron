package easycron

import (
	"errors"
	"log"
	"time"
)

// 定义获取时间接口
type Timer interface {
	NextDuration() (time.Duration, error)
}

var (
	_ Timer = RecurrentTimer{}
	_ Timer = DailyTimer{}
	_ Timer = OnceTimer{}
)

// 适合每隔几秒就执行的任务
type RecurrentTimer struct {
	Interval int64
}

func (r RecurrentTimer) NextDuration() (time.Duration, error) {
	return time.Second * time.Duration(r.Interval), nil
}

// 适合每天定时执行的任务
type DailyTimer struct {
	Hour, Minute, Second int
}

func (d DailyTimer) NextDuration() (time.Duration, error) {
	nowTime := time.Now()

	year, month, day := nowTime.Date()

	nextTime := time.Date(year, month, day, d.Hour, d.Minute, d.Second, 0, time.UTC)

	if ! nowTime.Before(nextTime) {
		nextTime = time.Date(year, month, day+1, d.Hour, d.Minute, d.Second, 0, time.UTC)
	}

	return nextTime.Sub(nowTime), nil
}

// 适合就执行一次的任务
type OnceTimer struct {
	Year, Month, Day int
	DailyTimer
}

func (once OnceTimer) NextDuration() (time.Duration, error) {
	nowTime := time.Now()

	nextTime := time.Date(once.Year, time.Month(once.Month), once.Day, once.Hour, once.Minute, once.Second, 0, time.UTC)

	if ! nowTime.Before(nextTime) {
		return time.Duration(0), errors.New("当前任务已失效，不能被执行")
	}

	return nextTime.Sub(nowTime), nil
}

// 执行任务
func Run(fn func(), timer Timer) {
	duration, err := timer.NextDuration()
	if err != nil {
		log.Println("任务建立失败:", err)
		return
	}

	go func() {
		for {
			select {
			case <-time.After(duration):
				go fn()
			}
		}
	}()
}
