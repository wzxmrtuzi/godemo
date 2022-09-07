package timedemo

import (
	"fmt"
	"time"
)

func TimingTest() {
	// newTimerTest()
	// afterTest()
	newTickerTest()
}

// 一次性定时器
// NewTimer函数
// func NewTimer(d Duration) *Timer
// NewTimer创建一个Timer，它会在到期后向Timer自身的C字段发送当时的时间
func newTimerTest() {
	var start = time.Now()
	fmt.Println("start", start)
	var timer = time.NewTimer(time.Second * 3)
	fmt.Println("before reading the code is executed")
	var end = <-timer.C
	fmt.Println("after reading the code is executed")
	fmt.Println("end", end)
}

// After函数
// func After(d Duration) <-chan Time
// 底层就是对NewTimer的封装, 只不过返回值不同而已
func afterTest() {
	var start = time.Now()
	fmt.Println("start", start)
	var timer = time.After(time.Second * 3)
	fmt.Println("before reading the code is executed")
	var end = <-timer
	fmt.Println("after reading the code is executed")
	fmt.Println("end", end)
}

// 周期性定时器
// NewTicker函数
// func NewTicker(d Duration) *Ticker
// 和NewTimer差不多, 只不过NewTimer只会往管道中写入一次数据, 而NewTicker每隔一段时间就会写一次
func newTickerTest() {
	var count = 0
	// 创建一个周期定时器
	var ticker = time.NewTicker(time.Second)
	//不断从重启定时器获取时间
	for {
		// 系统写入数据之前会被阻塞
		var t = <-ticker.C
		count++
		fmt.Println(t)
		if count > 6 {
			break
		}
	}
}
