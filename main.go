package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runJob(func() {
		fmt.Println("i am job1", time.Now().UTC())
	}, 5)

	runJob(func() {
		fmt.Println("i am job2", time.Now().UTC())
	}, 1)

	runtime.Goexit()
}

func runJob(fn func(), sec int64) {
	go func() {
		for {
			select {
			case <-time.After(time.Second * time.Duration(sec)):
				fn()
			}
		}
	}()
}
