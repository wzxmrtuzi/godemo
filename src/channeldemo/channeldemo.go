package channeldemo

import (
	"fmt"
	"math/rand"
	"time"
)

// Channel是线程安全的, 也就是自带锁定功能,channel的本质是一个队列
// Channel声明和初始化
// 	声明: var 变量名chan 数据类型
// 	初始化: mych := make(chan 数据类型, 容量)
// 	Channel和切片还有字典一样, 必须make之后才能使用
// 	Channel和切片还有字典一样, 是引用类型
func ChannelTest() {
	// createChannel()
	// foreachChannel()
	// producerAndConsumerTest()
	// channelBufTest()
	// oneWayChannel()
	channelAsParam()
}

// Channel阻塞现象
//	单独在主线程中操作管道, 写满了会报错, 没有数据去获取也会报错
//	只要在协程中操作管道过, 写满了就会阻塞, 没有就数据去获取也会阻塞

// 有缓冲channel
// var cpch = make(chan int, 5)
// 无缓冲channel
var cpch = make(chan int)
var existch = make(chan bool, 1)

// 无缓冲channel
var nobufch = make(chan int, 0)

// 无缓冲Channel和有缓冲Channel
// 有缓冲管道具备异步的能力(写几个读一个或读几个)
// 无缓冲管道具备同步的能力(写一个读一个)

// IO的延迟说明: 看到的输出结果和我们想象的不太一样,
// 是因为IO输出非常消耗性能, 输出之后还没来得及赋值可能就跑去执行别的协程了
func channelBufTest() {
	// 有缓冲channel只写入不读取不会报错
	for i := 0; i < 5; i++ {
		cpch <- i
	}
	go func() {
		fmt.Println("nobufch", <-nobufch)
	}()
	// 写入之后在同一个线程读取会报错,需要提前新开一个线程读取
	// 在主程中先写入, 在子程中后读取也会报错
	nobufch <- 2
}

func producerAndConsumerTest() {
	go producer()
	go consumer()
	fmt.Println("existch before")
	<-existch
	fmt.Println("existch after")
}

// channel实现生产者消费者关系
func producer() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		cpch <- num
		fmt.Println("producer", num, i)
	}
	// 写入结束关闭chennel
	close(cpch)
	fmt.Println("channel complete")
}
func consumer() {
	index := 0
	for {
		if value, ok := <-cpch; ok {
			fmt.Println("consuner", value, index)
			index++
		} else {
			break
		}
	}
	fmt.Println("consumer finish")
	existch <- true
}

// 注意点:
// 管道中只能存放声明的数据类型, 不能存放其它数据类型
// 管道中如果已经没有数据, 再取就会报错
// 如果管道中数据已满, 再写入就会报错
func createChannel() {
	// 声明
	var mych chan int
	// 初始化channel,make(类型,容量)
	mych = make(chan int, 3)
	// 查看channel的长度和容量
	printChannelInfo(mych)
	// 向channel中写入数据
	mych <- 666
	mych <- 123
	printChannelInfo(mych)
	// 取出channel的数据,先进先出
	data := <-mych
	fmt.Println("data", data)
	printChannelInfo(mych)
}

// 遍历channel
func foreachChannel() {
	mych := make(chan int, 5)
	// 写入5个数据
	for i := 1; i < 6; i++ {
		mych <- i * 2
	}
	printChannelInfo(mych)
	// 输出不完全
	// for i := 0; i < len(mych); i++ {
	// 	fmt.Println("foreachPrintChannel", <-mych)
	// }

	// 写完数据先关闭channel
	close(mych)
	// 使用range遍历需要先关闭管道
	// for value := range mych {
	// 	fmt.Println("channelValue", value)
	// }
	// printChannelInfo(mych)
	// close主要用途:
	// 在企业开发中我们可能不确定管道有还没有有数据, 所以我们可能一直获取
	// 但是我们可以通过ok-idiom模式判断管道是否关闭, 如果关闭会返回false给ok
	for {
		if value, ok := <-mych; ok {
			fmt.Println(value, ok)
		} else {
			break
		}
	}
	printChannelInfo(mych)
}

func printChannelInfo(mych chan int) {
	fmt.Println("length", len(mych), "capacity", cap(mych))
}

// 单向管道和双向管道
// 默认情况下所有管道都是双向了(可读可写)
// 但是在企业开发中, 我们经常需要用到将一个管道作为参数传递
// 在传递的过程中希望对方只能单向使用, 要么只能写,要么只能读

// 双向管道
// var myCh chan int = make(chan int, 0)
// 单向管道
// var myCh chan<- int = make(chan<- int, 0)
// var myCh <-chan int = make(<-chan int, 0)
// 注意点:
// 双向管道可以自动转换为任意一种单向管道
// 单向管道不能转换为双向管道

func oneWayChannel() {
	// channel之间赋值是地址传递,以下三个channel指向同一个channel
	var mych = make(chan int, 5)
	// 只写
	var mych2 chan<- int
	mych2 = mych
	fmt.Println(mych2)

	// 只读
	var mych3 <-chan int
	mych3 = mych
	fmt.Println(mych3)

	// 双向channel可读可写
	mych <- 3
	mych <- 4
	fmt.Println(<-mych)

	// 只写channel
	mych2 <- 5
	// 只读
	fmt.Println("asdf", <-mych3)
}

func channelAsParam() {
	var mych = make(chan int, 5)
	// 只写
	var producer = func(mych chan<- int) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			num := rand.Intn(100)
			mych <- num
			fmt.Println("producer", num, i)
		}
		close(mych)
	}

	// 只读
	var consumer = func(mych <-chan int) {
		for {
			if value, ok := <-mych; ok {
				fmt.Println("consumer", value)
			} else {
				break
			}
		}
	}
	go producer(mych)
	consumer(mych)
}
