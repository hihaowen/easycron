package easycron

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRecurrentTimerExecution(t *testing.T) {
	c := make(chan bool)

	fn := func() {
		c <- true
	}

	new(Task).Every(1).Run(fn)

	select {
	case <-c:
	case <-time.After(2 * time.Second):
		t.Error("没有间隔1秒运行")
	}
}

func TestDailyTimerNextDuration(t *testing.T) {
	nowTime := time.Now().UTC()

	actualStation, err := DailyTimer{nowTime.Hour(), nowTime.Minute(), nowTime.Second()}.NextDuration()

	actualTime := nowTime.Add(actualStation)

	assert.Nil(t, err)

	expectedTime := nowTime.AddDate(0, 0, 1)

	assert.Equal(t, expectedTime.Year(), actualTime.Year())
	assert.Equal(t, expectedTime.Month(), actualTime.Month())
	assert.Equal(t, expectedTime.Day(), actualTime.Day())
	assert.Equal(t, expectedTime.Hour(), actualTime.Hour())
	assert.Equal(t, expectedTime.Minute(), actualTime.Minute())
	// @todo
	// assert.Equal(t, expectedTime.Second(), actualTime.Second())
}
