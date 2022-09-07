package selectdemo

import (
	"fmt"
	"runtime"
	"time"
)

// select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作
// 如果有多个case都可以运行，select会随机选出一个执行，其他不会执行。
// 如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
// 如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

// select {
// case IO操作1:
// 	IO操作1读取或写入成功就执行
// case IO操作2:
// 	IO操作2读取或写入成功就执行
// default:
// 	如果上面case都没有成功，则进入default处理流程
// }
func SelectTest() {
	// channelSelectTest()
	selectSituation()
}

func channelSelectTest() {
	var mych = make(chan int, 5)
	var exitCh = make(chan bool)
	var producer = func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		exitCh <- true
	}
	go producer(mych)
	for {
		fmt.Println("start poll")
		select {
		case num := <-mych:
			fmt.Println("select", num)
		case <-exitCh:
			// 使用break没有作用
			return
		}
	}
}

// select应用场景
// 实现多路监听
// 实现超时处理
func selectSituation() {
	var mych = make(chan int, 5)
	var exitCh = make(chan bool)

	// 写入数据
	var producer = func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}
	// 读取数据
	var consumer = func(ch <-chan int) {
		for {
			select {
			case num := <-mych:
				fmt.Println(num)
			case <-time.After(time.Second * 2):
				fmt.Println("time after")
				exitCh <- true
				runtime.Goexit()
			}
		}
	}
	go producer(mych)
	go consumer(mych)
	fmt.Println(<-exitCh, "finish")
}
