package concurrencydemo

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// runtime常用函数
// Gosched():使当前go程放弃处理器，以让其它go程运行。
// 它不会挂起当前go程，因此当前go程未来会恢复执行

// Goexit(): 终止调用它的go程, 其它go程不会受影响

// NumCPU(): 返回本地机器的逻辑CPU个数
func ConcurrentCyTest() {
	go sing()
	go dance()

	cpuCount := runtime.NumCPU()
	fmt.Println("cpu-count", cpuCount)
	// 主线程不退出,如果退出所有的线程都会终止
	// for {

	// }
}

// 创建互斥锁
var lock = sync.Mutex{}

var sce []int = make([]int, 10)

func LockTest() {
	go producer()
	go consumer()
}

// 生产者
func producer() {
	// 加锁
	lock.Lock()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		sce[i] = num
		fmt.Println("producer:", num)
		time.Sleep(time.Millisecond * 500)
	}
	// 解锁
	lock.Unlock()
}

// 消费者
func consumer() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		num := sce[i]
		fmt.Println("consumer:", num)
	}
	lock.Unlock()
}

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在唱歌")
		// time.Sleep(time.Millisecond)
		// Gosched使当前go程放弃处理器，以让其它go程运行。
		// 它不会挂起当前go程，因此当前go程未来会恢复执行
		if i == 5 {
			runtime.Gosched()
		}
	}
}
func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在跳舞---")
		time.Sleep(time.Millisecond)
		if i == 3 {
			runtime.Goexit()
		}
		fmt.Println("dance finish")
	}
}
